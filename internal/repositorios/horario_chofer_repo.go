package repositorios

import (
	"database/sql"
	"errors"
	"sistema-tours/internal/entidades"
	"strconv"
	"time"
)

// HorarioChoferRepository maneja las operaciones de base de datos para horarios de chofer
type HorarioChoferRepository struct {
	db *sql.DB
}

// NewHorarioChoferRepository crea una nueva instancia del repositorio
func NewHorarioChoferRepository(db *sql.DB) *HorarioChoferRepository {
	return &HorarioChoferRepository{
		db: db,
	}
}

// GetByID obtiene un horario de chofer por su ID
func (r *HorarioChoferRepository) GetByID(id int) (*entidades.HorarioChofer, error) {
	horario := &entidades.HorarioChofer{}
	query := `SELECT hc.id_horario_chofer, hc.id_usuario, hc.hora_inicio, hc.hora_fin,
              hc.disponible_lunes, hc.disponible_martes, hc.disponible_miercoles, 
              hc.disponible_jueves, hc.disponible_viernes, hc.disponible_sabado, 
              hc.disponible_domingo, hc.fecha_inicio, hc.fecha_fin,
              u.nombres, u.apellidos, u.numero_documento, u.telefono
              FROM horario_chofer hc
              INNER JOIN usuario u ON hc.id_usuario = u.id_usuario
              WHERE hc.id_horario_chofer = $1`

	err := r.db.QueryRow(query, id).Scan(
		&horario.ID, &horario.IDUsuario, &horario.HoraInicio, &horario.HoraFin,
		&horario.DisponibleLunes, &horario.DisponibleMartes, &horario.DisponibleMiercoles,
		&horario.DisponibleJueves, &horario.DisponibleViernes, &horario.DisponibleSabado,
		&horario.DisponibleDomingo, &horario.FechaInicio, &horario.FechaFin,
		&horario.NombreChofer, &horario.ApellidosChofer, &horario.DocumentoChofer, &horario.TelefonoChofer,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("horario de chofer no encontrado")
		}
		return nil, err
	}

	return horario, nil
}

// Create guarda un nuevo horario de chofer en la base de datos
func (r *HorarioChoferRepository) Create(horario *entidades.NuevoHorarioChoferRequest) (int, error) {
	// Convertir strings HH:MM a time.Time para la base de datos
	horaInicio, err := parseTime(horario.HoraInicio)
	if err != nil {
		return 0, errors.New("formato de hora de inicio inválido, debe ser HH:MM")
	}

	horaFin, err := parseTime(horario.HoraFin)
	if err != nil {
		return 0, errors.New("formato de hora de fin inválido, debe ser HH:MM")
	}

	// Verificar que al menos un día esté disponible
	if !horario.DisponibleLunes && !horario.DisponibleMartes && !horario.DisponibleMiercoles &&
		!horario.DisponibleJueves && !horario.DisponibleViernes && !horario.DisponibleSabado &&
		!horario.DisponibleDomingo {
		return 0, errors.New("debe seleccionar al menos un día disponible")
	}

	var id int
	query := `INSERT INTO horario_chofer (id_usuario, hora_inicio, hora_fin, 
              disponible_lunes, disponible_martes, disponible_miercoles, 
              disponible_jueves, disponible_viernes, disponible_sabado, 
              disponible_domingo, fecha_inicio, fecha_fin) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) 
              RETURNING id_horario_chofer`

	err = r.db.QueryRow(
		query,
		horario.IDUsuario,
		horaInicio,
		horaFin,
		horario.DisponibleLunes,
		horario.DisponibleMartes,
		horario.DisponibleMiercoles,
		horario.DisponibleJueves,
		horario.DisponibleViernes,
		horario.DisponibleSabado,
		horario.DisponibleDomingo,
		horario.FechaInicio,
		horario.FechaFin,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// Update actualiza la información de un horario de chofer
func (r *HorarioChoferRepository) Update(id int, horario *entidades.ActualizarHorarioChoferRequest) error {
	// Convertir strings HH:MM a time.Time para la base de datos
	horaInicio, err := parseTime(horario.HoraInicio)
	if err != nil {
		return errors.New("formato de hora de inicio inválido, debe ser HH:MM")
	}

	horaFin, err := parseTime(horario.HoraFin)
	if err != nil {
		return errors.New("formato de hora de fin inválido, debe ser HH:MM")
	}

	// Verificar que al menos un día esté disponible
	if !horario.DisponibleLunes && !horario.DisponibleMartes && !horario.DisponibleMiercoles &&
		!horario.DisponibleJueves && !horario.DisponibleViernes && !horario.DisponibleSabado &&
		!horario.DisponibleDomingo {
		return errors.New("debe seleccionar al menos un día disponible")
	}

	query := `UPDATE horario_chofer SET 
              id_usuario = $1, 
              hora_inicio = $2, 
              hora_fin = $3, 
              disponible_lunes = $4, 
              disponible_martes = $5, 
              disponible_miercoles = $6, 
              disponible_jueves = $7, 
              disponible_viernes = $8, 
              disponible_sabado = $9, 
              disponible_domingo = $10,
              fecha_inicio = $11,
              fecha_fin = $12
              WHERE id_horario_chofer = $13`

	_, err = r.db.Exec(
		query,
		horario.IDUsuario,
		horaInicio,
		horaFin,
		horario.DisponibleLunes,
		horario.DisponibleMartes,
		horario.DisponibleMiercoles,
		horario.DisponibleJueves,
		horario.DisponibleViernes,
		horario.DisponibleSabado,
		horario.DisponibleDomingo,
		horario.FechaInicio,
		horario.FechaFin,
		id,
	)

	return err
}

// Delete elimina un horario de chofer
func (r *HorarioChoferRepository) Delete(id int) error {
	query := `DELETE FROM horario_chofer WHERE id_horario_chofer = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// List lista todos los horarios de chofer
func (r *HorarioChoferRepository) List() ([]*entidades.HorarioChofer, error) {
	query := `SELECT hc.id_horario_chofer, hc.id_usuario, hc.hora_inicio, hc.hora_fin,
              hc.disponible_lunes, hc.disponible_martes, hc.disponible_miercoles, 
              hc.disponible_jueves, hc.disponible_viernes, hc.disponible_sabado, 
              hc.disponible_domingo, hc.fecha_inicio, hc.fecha_fin,
              u.nombres, u.apellidos, u.numero_documento, u.telefono
              FROM horario_chofer hc
              INNER JOIN usuario u ON hc.id_usuario = u.id_usuario
              ORDER BY u.apellidos, u.nombres, hc.fecha_inicio DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	horarios := []*entidades.HorarioChofer{}

	for rows.Next() {
		horario := &entidades.HorarioChofer{}
		err := rows.Scan(
			&horario.ID, &horario.IDUsuario, &horario.HoraInicio, &horario.HoraFin,
			&horario.DisponibleLunes, &horario.DisponibleMartes, &horario.DisponibleMiercoles,
			&horario.DisponibleJueves, &horario.DisponibleViernes, &horario.DisponibleSabado,
			&horario.DisponibleDomingo, &horario.FechaInicio, &horario.FechaFin,
			&horario.NombreChofer, &horario.ApellidosChofer, &horario.DocumentoChofer, &horario.TelefonoChofer,
		)
		if err != nil {
			return nil, err
		}
		horarios = append(horarios, horario)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return horarios, nil
}

// ListByChofer lista todos los horarios de un chofer específico
func (r *HorarioChoferRepository) ListByChofer(idChofer int) ([]*entidades.HorarioChofer, error) {
	query := `SELECT hc.id_horario_chofer, hc.id_usuario, hc.hora_inicio, hc.hora_fin,
              hc.disponible_lunes, hc.disponible_martes, hc.disponible_miercoles, 
              hc.disponible_jueves, hc.disponible_viernes, hc.disponible_sabado, 
              hc.disponible_domingo, hc.fecha_inicio, hc.fecha_fin,
              u.nombres, u.apellidos, u.numero_documento, u.telefono
              FROM horario_chofer hc
              INNER JOIN usuario u ON hc.id_usuario = u.id_usuario
              WHERE hc.id_usuario = $1
              ORDER BY hc.fecha_inicio DESC`

	rows, err := r.db.Query(query, idChofer)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	horarios := []*entidades.HorarioChofer{}

	for rows.Next() {
		horario := &entidades.HorarioChofer{}
		err := rows.Scan(
			&horario.ID, &horario.IDUsuario, &horario.HoraInicio, &horario.HoraFin,
			&horario.DisponibleLunes, &horario.DisponibleMartes, &horario.DisponibleMiercoles,
			&horario.DisponibleJueves, &horario.DisponibleViernes, &horario.DisponibleSabado,
			&horario.DisponibleDomingo, &horario.FechaInicio, &horario.FechaFin,
			&horario.NombreChofer, &horario.ApellidosChofer, &horario.DocumentoChofer, &horario.TelefonoChofer,
		)
		if err != nil {
			return nil, err
		}
		horarios = append(horarios, horario)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return horarios, nil
}

// ListActiveByChofer lista los horarios activos de un chofer (donde la fecha actual está dentro del rango fecha_inicio y fecha_fin)
func (r *HorarioChoferRepository) ListActiveByChofer(idChofer int) ([]*entidades.HorarioChofer, error) {
	query := `SELECT hc.id_horario_chofer, hc.id_usuario, hc.hora_inicio, hc.hora_fin,
              hc.disponible_lunes, hc.disponible_martes, hc.disponible_miercoles, 
              hc.disponible_jueves, hc.disponible_viernes, hc.disponible_sabado, 
              hc.disponible_domingo, hc.fecha_inicio, hc.fecha_fin,
              u.nombres, u.apellidos, u.numero_documento, u.telefono
              FROM horario_chofer hc
              INNER JOIN usuario u ON hc.id_usuario = u.id_usuario
              WHERE hc.id_usuario = $1
              AND hc.fecha_inicio <= CURRENT_DATE
              AND (hc.fecha_fin IS NULL OR hc.fecha_fin >= CURRENT_DATE)
              ORDER BY hc.fecha_inicio DESC`

	rows, err := r.db.Query(query, idChofer)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	horarios := []*entidades.HorarioChofer{}

	for rows.Next() {
		horario := &entidades.HorarioChofer{}
		err := rows.Scan(
			&horario.ID, &horario.IDUsuario, &horario.HoraInicio, &horario.HoraFin,
			&horario.DisponibleLunes, &horario.DisponibleMartes, &horario.DisponibleMiercoles,
			&horario.DisponibleJueves, &horario.DisponibleViernes, &horario.DisponibleSabado,
			&horario.DisponibleDomingo, &horario.FechaInicio, &horario.FechaFin,
			&horario.NombreChofer, &horario.ApellidosChofer, &horario.DocumentoChofer, &horario.TelefonoChofer,
		)
		if err != nil {
			return nil, err
		}
		horarios = append(horarios, horario)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return horarios, nil
}

// ListByDia lista todos los horarios de choferes disponibles para un día específico (1=Lunes, 7=Domingo)
func (r *HorarioChoferRepository) ListByDia(diaSemana int) ([]*entidades.HorarioChofer, error) {
	var condition string
	switch diaSemana {
	case 1:
		condition = "hc.disponible_lunes = true"
	case 2:
		condition = "hc.disponible_martes = true"
	case 3:
		condition = "hc.disponible_miercoles = true"
	case 4:
		condition = "hc.disponible_jueves = true"
	case 5:
		condition = "hc.disponible_viernes = true"
	case 6:
		condition = "hc.disponible_sabado = true"
	case 7:
		condition = "hc.disponible_domingo = true"
	default:
		return nil, errors.New("día de la semana inválido, debe ser un número entre 1 (Lunes) y 7 (Domingo)")
	}

	query := `SELECT hc.id_horario_chofer, hc.id_usuario, hc.hora_inicio, hc.hora_fin,
              hc.disponible_lunes, hc.disponible_martes, hc.disponible_miercoles, 
              hc.disponible_jueves, hc.disponible_viernes, hc.disponible_sabado, 
              hc.disponible_domingo, hc.fecha_inicio, hc.fecha_fin,
              u.nombres, u.apellidos, u.numero_documento, u.telefono
              FROM horario_chofer hc
              INNER JOIN usuario u ON hc.id_usuario = u.id_usuario
              WHERE ` + condition + `
              AND hc.fecha_inicio <= CURRENT_DATE
              AND (hc.fecha_fin IS NULL OR hc.fecha_fin >= CURRENT_DATE)
              ORDER BY u.apellidos, u.nombres, hc.hora_inicio`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	horarios := []*entidades.HorarioChofer{}

	for rows.Next() {
		horario := &entidades.HorarioChofer{}
		err := rows.Scan(
			&horario.ID, &horario.IDUsuario, &horario.HoraInicio, &horario.HoraFin,
			&horario.DisponibleLunes, &horario.DisponibleMartes, &horario.DisponibleMiercoles,
			&horario.DisponibleJueves, &horario.DisponibleViernes, &horario.DisponibleSabado,
			&horario.DisponibleDomingo, &horario.FechaInicio, &horario.FechaFin,
			&horario.NombreChofer, &horario.ApellidosChofer, &horario.DocumentoChofer, &horario.TelefonoChofer,
		)
		if err != nil {
			return nil, err
		}
		horarios = append(horarios, horario)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return horarios, nil
}

// VerifyHorarioOverlap verifica si hay solapamiento entre horarios para un mismo chofer
func (r *HorarioChoferRepository) VerifyHorarioOverlap(idChofer int, horaInicio, horaFin time.Time, fechaInicio, fechaFin *time.Time, excludeID int) (bool, error) {
	var query string
	var args []interface{}

	if fechaFin == nil {
		// No tiene fecha fin (indefinido)
		query = `SELECT COUNT(*) FROM horario_chofer 
                  WHERE id_usuario = $1 
                  AND ((hora_inicio <= $2 AND hora_fin > $2) OR (hora_inicio < $3 AND hora_fin >= $3) OR (hora_inicio >= $2 AND hora_fin <= $3))
                  AND fecha_inicio <= $4
                  AND (fecha_fin IS NULL OR fecha_fin >= $4)`
		args = []interface{}{idChofer, horaInicio, horaFin, *fechaInicio}
	} else {
		// Tiene fecha fin
		query = `SELECT COUNT(*) FROM horario_chofer 
                  WHERE id_usuario = $1 
                  AND ((hora_inicio <= $2 AND hora_fin > $2) OR (hora_inicio < $3 AND hora_fin >= $3) OR (hora_inicio >= $2 AND hora_fin <= $3))
                  AND (
                    (fecha_inicio <= $4 AND (fecha_fin IS NULL OR fecha_fin >= $4)) OR 
                    (fecha_inicio <= $5 AND (fecha_fin IS NULL OR fecha_fin >= $5)) OR
                    (fecha_inicio >= $4 AND (fecha_fin IS NULL OR fecha_fin <= $5))
                  )`
		args = []interface{}{idChofer, horaInicio, horaFin, *fechaInicio, *fechaFin}
	}

	if excludeID > 0 {
		query += " AND id_horario_chofer != $" + strconv.Itoa(len(args)+1)
		args = append(args, excludeID)
	}

	var count int
	err := r.db.QueryRow(query, args...).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
