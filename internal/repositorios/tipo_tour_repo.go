package repositorios

import (
	"database/sql"
	"errors"
	"sistema-tours/internal/entidades"
	"time"
)

// TipoTourRepository maneja las operaciones de base de datos para tipos de tour
type TipoTourRepository struct {
	db *sql.DB
}

// NewTipoTourRepository crea una nueva instancia del repositorio
func NewTipoTourRepository(db *sql.DB) *TipoTourRepository {
	return &TipoTourRepository{
		db: db,
	}
}

// GetByID obtiene un tipo de tour por su ID
func (r *TipoTourRepository) GetByID(id int) (*entidades.TipoTour, error) {
	tipoTour := &entidades.TipoTour{}
	query := `SELECT id_tipo_tour, nombre, descripcion, duracion_minutos, precio_base, 
              cantidad_pasajeros, url_imagen 
              FROM tipo_tour 
              WHERE id_tipo_tour = $1`

	err := r.db.QueryRow(query, id).Scan(
		&tipoTour.ID, &tipoTour.Nombre, &tipoTour.Descripcion, &tipoTour.DuracionMinutos,
		&tipoTour.PrecioBase, &tipoTour.CantidadPasajeros, &tipoTour.URLImagen,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("tipo de tour no encontrado")
		}
		return nil, err
	}

	// Obtener los horarios relacionados
	horarios, err := r.GetHorariosByTipoTourID(id)
	if err == nil {
		tipoTour.Horarios = horarios
	}

	// Contar tours programados
	countQuery := `SELECT COUNT(*) FROM tour_programado WHERE id_tipo_tour = $1`
	err = r.db.QueryRow(countQuery, id).Scan(&tipoTour.ToursProgramados)
	if err != nil {
		tipoTour.ToursProgramados = 0
	}

	return tipoTour, nil
}

// GetByNombre obtiene un tipo de tour por su nombre
func (r *TipoTourRepository) GetByNombre(nombre string) (*entidades.TipoTour, error) {
	tipoTour := &entidades.TipoTour{}
	query := `SELECT id_tipo_tour, nombre, descripcion, duracion_minutos, precio_base, 
              cantidad_pasajeros, url_imagen 
              FROM tipo_tour 
              WHERE nombre = $1`

	err := r.db.QueryRow(query, nombre).Scan(
		&tipoTour.ID, &tipoTour.Nombre, &tipoTour.Descripcion, &tipoTour.DuracionMinutos,
		&tipoTour.PrecioBase, &tipoTour.CantidadPasajeros, &tipoTour.URLImagen,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("tipo de tour no encontrado")
		}
		return nil, err
	}

	return tipoTour, nil
}

// Create guarda un nuevo tipo de tour en la base de datos
func (r *TipoTourRepository) Create(tipoTour *entidades.NuevoTipoTourRequest) (int, error) {
	var id int
	query := `INSERT INTO tipo_tour (nombre, descripcion, duracion_minutos, precio_base, 
              cantidad_pasajeros, url_imagen) 
              VALUES ($1, $2, $3, $4, $5, $6) 
              RETURNING id_tipo_tour`

	err := r.db.QueryRow(
		query,
		tipoTour.Nombre,
		tipoTour.Descripcion,
		tipoTour.DuracionMinutos,
		tipoTour.PrecioBase,
		tipoTour.CantidadPasajeros,
		tipoTour.URLImagen,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// Update actualiza la información de un tipo de tour
func (r *TipoTourRepository) Update(id int, tipoTour *entidades.ActualizarTipoTourRequest) error {
	query := `UPDATE tipo_tour SET 
              nombre = $1, 
              descripcion = $2, 
              duracion_minutos = $3, 
              precio_base = $4, 
              cantidad_pasajeros = $5, 
              url_imagen = $6 
              WHERE id_tipo_tour = $7`

	_, err := r.db.Exec(
		query,
		tipoTour.Nombre,
		tipoTour.Descripcion,
		tipoTour.DuracionMinutos,
		tipoTour.PrecioBase,
		tipoTour.CantidadPasajeros,
		tipoTour.URLImagen,
		id,
	)

	return err
}

// Delete elimina un tipo de tour
func (r *TipoTourRepository) Delete(id int) error {
	// Primero verificar que no tenga tours programados
	var count int
	countQuery := `SELECT COUNT(*) FROM tour_programado WHERE id_tipo_tour = $1`
	err := r.db.QueryRow(countQuery, id).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("no se puede eliminar el tipo de tour porque tiene tours programados")
	}

	// Eliminar horarios relacionados
	deleteHorariosQuery := `DELETE FROM horario_tour WHERE id_tipo_tour = $1`
	_, err = r.db.Exec(deleteHorariosQuery, id)
	if err != nil {
		return err
	}

	// Eliminar el tipo de tour
	deleteTipoTourQuery := `DELETE FROM tipo_tour WHERE id_tipo_tour = $1`
	_, err = r.db.Exec(deleteTipoTourQuery, id)
	return err
}

// List lista todos los tipos de tour
func (r *TipoTourRepository) List() ([]*entidades.TipoTour, error) {
	query := `SELECT id_tipo_tour, nombre, descripcion, duracion_minutos, precio_base, 
              cantidad_pasajeros, url_imagen 
              FROM tipo_tour 
              ORDER BY nombre`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tiposTour := []*entidades.TipoTour{}

	for rows.Next() {
		tipoTour := &entidades.TipoTour{}
		err := rows.Scan(
			&tipoTour.ID, &tipoTour.Nombre, &tipoTour.Descripcion, &tipoTour.DuracionMinutos,
			&tipoTour.PrecioBase, &tipoTour.CantidadPasajeros, &tipoTour.URLImagen,
		)
		if err != nil {
			return nil, err
		}

		// Contar tours programados
		countQuery := `SELECT COUNT(*) FROM tour_programado WHERE id_tipo_tour = $1`
		err = r.db.QueryRow(countQuery, tipoTour.ID).Scan(&tipoTour.ToursProgramados)
		if err != nil {
			tipoTour.ToursProgramados = 0
		}

		tiposTour = append(tiposTour, tipoTour)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tiposTour, nil
}

// GetHorariosByTipoTourID obtiene los horarios de un tipo de tour
func (r *TipoTourRepository) GetHorariosByTipoTourID(idTipoTour int) ([]*entidades.HorarioTour, error) {
	query := `SELECT id_horario, id_tipo_tour, hora_inicio, hora_fin, 
              disponible_lunes, disponible_martes, disponible_miercoles, 
              disponible_jueves, disponible_viernes, disponible_sabado, 
              disponible_domingo 
              FROM horario_tour 
              WHERE id_tipo_tour = $1 
              ORDER BY hora_inicio`

	rows, err := r.db.Query(query, idTipoTour)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	horarios := []*entidades.HorarioTour{}

	for rows.Next() {
		horario := &entidades.HorarioTour{}
		err := rows.Scan(
			&horario.ID, &horario.IDTipoTour, &horario.HoraInicio, &horario.HoraFin,
			&horario.DisponibleLunes, &horario.DisponibleMartes, &horario.DisponibleMiercoles,
			&horario.DisponibleJueves, &horario.DisponibleViernes, &horario.DisponibleSabado,
			&horario.DisponibleDomingo,
		)
		if err != nil {
			return nil, err
		}

		// Obtener nombre del tipo de tour
		var nombreTipoTour string
		nombreQuery := `SELECT nombre FROM tipo_tour WHERE id_tipo_tour = $1`
		err = r.db.QueryRow(nombreQuery, horario.IDTipoTour).Scan(&nombreTipoTour)
		if err == nil {
			horario.NombreTipoTour = nombreTipoTour
		}

		horarios = append(horarios, horario)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return horarios, nil
}

// GetHorarioByID obtiene un horario por su ID
func (r *TipoTourRepository) GetHorarioByID(id int) (*entidades.HorarioTour, error) {
	horario := &entidades.HorarioTour{}
	query := `SELECT id_horario, id_tipo_tour, hora_inicio, hora_fin, 
              disponible_lunes, disponible_martes, disponible_miercoles, 
              disponible_jueves, disponible_viernes, disponible_sabado, 
              disponible_domingo 
              FROM horario_tour 
              WHERE id_horario = $1`

	err := r.db.QueryRow(query, id).Scan(
		&horario.ID, &horario.IDTipoTour, &horario.HoraInicio, &horario.HoraFin,
		&horario.DisponibleLunes, &horario.DisponibleMartes, &horario.DisponibleMiercoles,
		&horario.DisponibleJueves, &horario.DisponibleViernes, &horario.DisponibleSabado,
		&horario.DisponibleDomingo,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("horario no encontrado")
		}
		return nil, err
	}

	// Obtener nombre del tipo de tour
	var nombreTipoTour string
	nombreQuery := `SELECT nombre FROM tipo_tour WHERE id_tipo_tour = $1`
	err = r.db.QueryRow(nombreQuery, horario.IDTipoTour).Scan(&nombreTipoTour)
	if err == nil {
		horario.NombreTipoTour = nombreTipoTour
	}

	return horario, nil
}

// CreateHorario guarda un nuevo horario de tour
func (r *TipoTourRepository) CreateHorario(horario *entidades.NuevoHorarioTourRequest) (int, error) {
	var id int

	// Convertir strings de hora a time.Time
	horaInicio, err := time.Parse("15:04", horario.HoraInicio)
	if err != nil {
		return 0, errors.New("formato de hora de inicio inválido")
	}

	horaFin, err := time.Parse("15:04", horario.HoraFin)
	if err != nil {
		return 0, errors.New("formato de hora de fin inválido")
	}

	query := `INSERT INTO horario_tour (id_tipo_tour, hora_inicio, hora_fin, 
              disponible_lunes, disponible_martes, disponible_miercoles, 
              disponible_jueves, disponible_viernes, disponible_sabado, 
              disponible_domingo) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
              RETURNING id_horario`

	err = r.db.QueryRow(
		query,
		horario.IDTipoTour,
		horaInicio,
		horaFin,
		horario.DisponibleLunes,
		horario.DisponibleMartes,
		horario.DisponibleMiercoles,
		horario.DisponibleJueves,
		horario.DisponibleViernes,
		horario.DisponibleSabado,
		horario.DisponibleDomingo,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateHorario actualiza un horario de tour
func (r *TipoTourRepository) UpdateHorario(id int, horario *entidades.ActualizarHorarioTourRequest) error {
	// Convertir strings de hora a time.Time
	horaInicio, err := time.Parse("15:04", horario.HoraInicio)
	if err != nil {
		return errors.New("formato de hora de inicio inválido")
	}

	horaFin, err := time.Parse("15:04", horario.HoraFin)
	if err != nil {
		return errors.New("formato de hora de fin inválido")
	}

	query := `UPDATE horario_tour SET 
              id_tipo_tour = $1, 
              hora_inicio = $2, 
              hora_fin = $3, 
              disponible_lunes = $4, 
              disponible_martes = $5, 
              disponible_miercoles = $6, 
              disponible_jueves = $7, 
              disponible_viernes = $8, 
              disponible_sabado = $9, 
              disponible_domingo = $10 
              WHERE id_horario = $11`

	_, err = r.db.Exec(
		query,
		horario.IDTipoTour,
		horaInicio,
		horaFin,
		horario.DisponibleLunes,
		horario.DisponibleMartes,
		horario.DisponibleMiercoles,
		horario.DisponibleJueves,
		horario.DisponibleViernes,
		horario.DisponibleSabado,
		horario.DisponibleDomingo,
		id,
	)

	return err
}

// DeleteHorario elimina un horario de tour
func (r *TipoTourRepository) DeleteHorario(id int) error {
	// Verificar que no haya tours programados con este horario
	var count int
	countQuery := `SELECT COUNT(*) FROM tour_programado WHERE id_horario = $1`
	err := r.db.QueryRow(countQuery, id).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("no se puede eliminar el horario porque tiene tours programados")
	}

	// Eliminar el horario
	query := `DELETE FROM horario_tour WHERE id_horario = $1`
	_, err = r.db.Exec(query, id)
	return err
}

// ListHorarios lista todos los horarios de tour
func (r *TipoTourRepository) ListHorarios() ([]*entidades.HorarioTour, error) {
	query := `SELECT h.id_horario, h.id_tipo_tour, h.hora_inicio, h.hora_fin, 
              h.disponible_lunes, h.disponible_martes, h.disponible_miercoles, 
              h.disponible_jueves, h.disponible_viernes, h.disponible_sabado, 
              h.disponible_domingo, t.nombre
              FROM horario_tour h
              INNER JOIN tipo_tour t ON h.id_tipo_tour = t.id_tipo_tour
              ORDER BY t.nombre, h.hora_inicio`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	horarios := []*entidades.HorarioTour{}

	for rows.Next() {
		horario := &entidades.HorarioTour{}
		err := rows.Scan(
			&horario.ID, &horario.IDTipoTour, &horario.HoraInicio, &horario.HoraFin,
			&horario.DisponibleLunes, &horario.DisponibleMartes, &horario.DisponibleMiercoles,
			&horario.DisponibleJueves, &horario.DisponibleViernes, &horario.DisponibleSabado,
			&horario.DisponibleDomingo, &horario.NombreTipoTour,
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
