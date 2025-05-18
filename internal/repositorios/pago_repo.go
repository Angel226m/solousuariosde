package repositorios

import (
	"database/sql"
	"errors"
	"sistema-tours/internal/entidades"
	"time"
)

// PagoRepository maneja las operaciones de base de datos para pagos
type PagoRepository struct {
	db *sql.DB
}

// NewPagoRepository crea una nueva instancia del repositorio
func NewPagoRepository(db *sql.DB) *PagoRepository {
	return &PagoRepository{
		db: db,
	}
}

// GetByID obtiene un pago por su ID
func (r *PagoRepository) GetByID(id int) (*entidades.Pago, error) {
	pago := &entidades.Pago{}
	query := `SELECT p.id_pago, p.id_reserva, p.id_metodo_pago, p.id_canal, 
              p.monto, p.fecha_pago, p.comprobante, p.estado,
              r.id_reserva as numero_reserva, c.nombres, c.apellidos, c.numero_documento,
              mp.nombre as nombre_metodo_pago, cv.nombre as nombre_canal_venta
              FROM pago p
              INNER JOIN reserva r ON p.id_reserva = r.id_reserva
              INNER JOIN cliente c ON r.id_cliente = c.id_cliente
              INNER JOIN metodo_pago mp ON p.id_metodo_pago = mp.id_metodo_pago
              INNER JOIN canal_venta cv ON p.id_canal = cv.id_canal
              WHERE p.id_pago = $1`

	err := r.db.QueryRow(query, id).Scan(
		&pago.ID, &pago.IDReserva, &pago.IDMetodoPago, &pago.IDCanal,
		&pago.Monto, &pago.FechaPago, &pago.Comprobante, &pago.Estado,
		&pago.NumeroReserva, &pago.NombreCliente, &pago.ApellidosCliente, &pago.DocumentoCliente,
		&pago.NombreMetodoPago, &pago.NombreCanalVenta,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("pago no encontrado")
		}
		return nil, err
	}

	return pago, nil
}

// Create guarda un nuevo pago en la base de datos
func (r *PagoRepository) Create(pago *entidades.NuevoPagoRequest) (int, error) {
	// Verificar que la reserva exista y su estado sea válido
	var estadoReserva string
	queryReserva := `SELECT estado FROM reserva WHERE id_reserva = $1`
	err := r.db.QueryRow(queryReserva, pago.IDReserva).Scan(&estadoReserva)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("la reserva especificada no existe")
		}
		return 0, err
	}

	if estadoReserva != "RESERVADO" {
		return 0, errors.New("solo se pueden registrar pagos para reservas en estado RESERVADO")
	}

	// Crear transacción para asegurar integridad
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	// Crear pago
	var id int
	query := `INSERT INTO pago (id_reserva, id_metodo_pago, id_canal, monto, comprobante)
              VALUES ($1, $2, $3, $4, $5)
              RETURNING id_pago`

	err = tx.QueryRow(
		query,
		pago.IDReserva,
		pago.IDMetodoPago,
		pago.IDCanal,
		pago.Monto,
		pago.Comprobante,
	).Scan(&id)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// Confirmar transacción
	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Update actualiza la información de un pago
func (r *PagoRepository) Update(id int, pago *entidades.ActualizarPagoRequest) error {
	query := `UPDATE pago SET
              id_reserva = $1,
              id_metodo_pago = $2,
              id_canal = $3,
              monto = $4,
              comprobante = $5,
              estado = $6
              WHERE id_pago = $7`

	_, err := r.db.Exec(
		query,
		pago.IDReserva,
		pago.IDMetodoPago,
		pago.IDCanal,
		pago.Monto,
		pago.Comprobante,
		pago.Estado,
		id,
	)

	return err
}

// UpdateEstado actualiza solo el estado de un pago
func (r *PagoRepository) UpdateEstado(id int, estado string) error {
	query := `UPDATE pago SET estado = $1 WHERE id_pago = $2`
	_, err := r.db.Exec(query, estado, id)
	return err
}

// Delete elimina un pago
func (r *PagoRepository) Delete(id int) error {
	// Verificar si el pago tiene comprobantes asociados
	var count int
	queryCheck := `SELECT COUNT(*) FROM comprobante_pago WHERE id_pago = $1`
	err := r.db.QueryRow(queryCheck, id).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("no se puede eliminar este pago porque tiene comprobantes asociados")
	}

	// Eliminar pago
	query := `DELETE FROM pago WHERE id_pago = $1`
	_, err = r.db.Exec(query, id)
	return err
}

// List lista todos los pagos
func (r *PagoRepository) List() ([]*entidades.Pago, error) {
	query := `SELECT p.id_pago, p.id_reserva, p.id_metodo_pago, p.id_canal, 
              p.monto, p.fecha_pago, p.comprobante, p.estado,
              r.id_reserva as numero_reserva, c.nombres, c.apellidos, c.numero_documento,
              mp.nombre as nombre_metodo_pago, cv.nombre as nombre_canal_venta
              FROM pago p
              INNER JOIN reserva r ON p.id_reserva = r.id_reserva
              INNER JOIN cliente c ON r.id_cliente = c.id_cliente
              INNER JOIN metodo_pago mp ON p.id_metodo_pago = mp.id_metodo_pago
              INNER JOIN canal_venta cv ON p.id_canal = cv.id_canal
              ORDER BY p.fecha_pago DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pagos := []*entidades.Pago{}

	for rows.Next() {
		pago := &entidades.Pago{}
		err := rows.Scan(
			&pago.ID, &pago.IDReserva, &pago.IDMetodoPago, &pago.IDCanal,
			&pago.Monto, &pago.FechaPago, &pago.Comprobante, &pago.Estado,
			&pago.NumeroReserva, &pago.NombreCliente, &pago.ApellidosCliente, &pago.DocumentoCliente,
			&pago.NombreMetodoPago, &pago.NombreCanalVenta,
		)
		if err != nil {
			return nil, err
		}
		pagos = append(pagos, pago)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pagos, nil
}

// ListByReserva lista todos los pagos de una reserva específica
func (r *PagoRepository) ListByReserva(idReserva int) ([]*entidades.Pago, error) {
	query := `SELECT p.id_pago, p.id_reserva, p.id_metodo_pago, p.id_canal, 
              p.monto, p.fecha_pago, p.comprobante, p.estado,
              r.id_reserva as numero_reserva, c.nombres, c.apellidos, c.numero_documento,
              mp.nombre as nombre_metodo_pago, cv.nombre as nombre_canal_venta
              FROM pago p
              INNER JOIN reserva r ON p.id_reserva = r.id_reserva
              INNER JOIN cliente c ON r.id_cliente = c.id_cliente
              INNER JOIN metodo_pago mp ON p.id_metodo_pago = mp.id_metodo_pago
              INNER JOIN canal_venta cv ON p.id_canal = cv.id_canal
              WHERE p.id_reserva = $1
              ORDER BY p.fecha_pago DESC`

	rows, err := r.db.Query(query, idReserva)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pagos := []*entidades.Pago{}

	for rows.Next() {
		pago := &entidades.Pago{}
		err := rows.Scan(
			&pago.ID, &pago.IDReserva, &pago.IDMetodoPago, &pago.IDCanal,
			&pago.Monto, &pago.FechaPago, &pago.Comprobante, &pago.Estado,
			&pago.NumeroReserva, &pago.NombreCliente, &pago.ApellidosCliente, &pago.DocumentoCliente,
			&pago.NombreMetodoPago, &pago.NombreCanalVenta,
		)
		if err != nil {
			return nil, err
		}
		pagos = append(pagos, pago)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pagos, nil
}

// ListByEstado lista todos los pagos con un estado específico
func (r *PagoRepository) ListByEstado(estado string) ([]*entidades.Pago, error) {
	query := `SELECT p.id_pago, p.id_reserva, p.id_metodo_pago, p.id_canal, 
              p.monto, p.fecha_pago, p.comprobante, p.estado,
              r.id_reserva as numero_reserva, c.nombres, c.apellidos, c.numero_documento,
              mp.nombre as nombre_metodo_pago, cv.nombre as nombre_canal_venta
              FROM pago p
              INNER JOIN reserva r ON p.id_reserva = r.id_reserva
              INNER JOIN cliente c ON r.id_cliente = c.id_cliente
              INNER JOIN metodo_pago mp ON p.id_metodo_pago = mp.id_metodo_pago
              INNER JOIN canal_venta cv ON p.id_canal = cv.id_canal
              WHERE p.estado = $1
              ORDER BY p.fecha_pago DESC`

	rows, err := r.db.Query(query, estado)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pagos := []*entidades.Pago{}

	for rows.Next() {
		pago := &entidades.Pago{}
		err := rows.Scan(
			&pago.ID, &pago.IDReserva, &pago.IDMetodoPago, &pago.IDCanal,
			&pago.Monto, &pago.FechaPago, &pago.Comprobante, &pago.Estado,
			&pago.NumeroReserva, &pago.NombreCliente, &pago.ApellidosCliente, &pago.DocumentoCliente,
			&pago.NombreMetodoPago, &pago.NombreCanalVenta,
		)
		if err != nil {
			return nil, err
		}
		pagos = append(pagos, pago)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pagos, nil
}

// ListByFecha lista todos los pagos de una fecha específica
func (r *PagoRepository) ListByFecha(fecha time.Time) ([]*entidades.Pago, error) {
	query := `SELECT p.id_pago, p.id_reserva, p.id_metodo_pago, p.id_canal, 
              p.monto, p.fecha_pago, p.comprobante, p.estado,
              r.id_reserva as numero_reserva, c.nombres, c.apellidos, c.numero_documento,
              mp.nombre as nombre_metodo_pago, cv.nombre as nombre_canal_venta
              FROM pago p
              INNER JOIN reserva r ON p.id_reserva = r.id_reserva
              INNER JOIN cliente c ON r.id_cliente = c.id_cliente
              INNER JOIN metodo_pago mp ON p.id_metodo_pago = mp.id_metodo_pago
              INNER JOIN canal_venta cv ON p.id_canal = cv.id_canal
              WHERE DATE(p.fecha_pago) = $1
              ORDER BY p.fecha_pago DESC`

	rows, err := r.db.Query(query, fecha)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pagos := []*entidades.Pago{}

	for rows.Next() {
		pago := &entidades.Pago{}
		err := rows.Scan(
			&pago.ID, &pago.IDReserva, &pago.IDMetodoPago, &pago.IDCanal,
			&pago.Monto, &pago.FechaPago, &pago.Comprobante, &pago.Estado,
			&pago.NumeroReserva, &pago.NombreCliente, &pago.ApellidosCliente, &pago.DocumentoCliente,
			&pago.NombreMetodoPago, &pago.NombreCanalVenta,
		)
		if err != nil {
			return nil, err
		}
		pagos = append(pagos, pago)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pagos, nil
}

// GetTotalPagadoByReserva obtiene el total pagado de una reserva específica
func (r *PagoRepository) GetTotalPagadoByReserva(idReserva int) (float64, error) {
	var totalPagado float64
	query := `SELECT COALESCE(SUM(monto), 0) 
              FROM pago 
              WHERE id_reserva = $1 AND estado = 'PROCESADO'`

	err := r.db.QueryRow(query, idReserva).Scan(&totalPagado)
	if err != nil {
		return 0, err
	}

	return totalPagado, nil
}
