package entidades

import "time"

// TourProgramado representa la estructura de un tour programado en el sistema
type TourProgramado struct {
	ID             int       `json:"id_tour_programado" db:"id_tour_programado"`
	IDTipoTour     int       `json:"id_tipo_tour" db:"id_tipo_tour"`
	IDEmbarcacion  int       `json:"id_embarcacion" db:"id_embarcacion"`
	IDHorario      int       `json:"id_horario" db:"id_horario"`
	Fecha          time.Time `json:"fecha" db:"fecha"`
	CupoMaximo     int       `json:"cupo_maximo" db:"cupo_maximo"`
	CupoDisponible int       `json:"cupo_disponible" db:"cupo_disponible"`
	Estado         string    `json:"estado" db:"estado"` // PROGRAMADO, COMPLETADO, CANCELADO

	// Campos adicionales para mostrar información relacionada
	NombreTipoTour       string  `json:"nombre_tipo_tour,omitempty" db:"-"`
	PrecioBase           float64 `json:"precio_base,omitempty" db:"-"`
	DuracionMinutos      int     `json:"duracion_minutos,omitempty" db:"-"`
	NombreEmbarcacion    string  `json:"nombre_embarcacion,omitempty" db:"-"`
	CapacidadEmbarcacion int     `json:"capacidad_embarcacion,omitempty" db:"-"`
	NombreChofer         string  `json:"nombre_chofer,omitempty" db:"-"`
	ApellidosChofer      string  `json:"apellidos_chofer,omitempty" db:"-"`
	HoraInicio           string  `json:"hora_inicio,omitempty" db:"-"`
	HoraFin              string  `json:"hora_fin,omitempty" db:"-"`
}

// NuevoTourProgramadoRequest representa los datos necesarios para crear un nuevo tour programado
type NuevoTourProgramadoRequest struct {
	IDTipoTour    int       `json:"id_tipo_tour" validate:"required"`
	IDEmbarcacion int       `json:"id_embarcacion" validate:"required"`
	IDHorario     int       `json:"id_horario" validate:"required"`
	Fecha         time.Time `json:"fecha" validate:"required"`
	CupoMaximo    int       `json:"cupo_maximo" validate:"required,min=1"`
	Estado        string    `json:"estado" validate:"omitempty,oneof=PROGRAMADO COMPLETADO CANCELADO"`
}

// ActualizarTourProgramadoRequest representa los datos para actualizar un tour programado
type ActualizarTourProgramadoRequest struct {
	IDTipoTour     int       `json:"id_tipo_tour" validate:"required"`
	IDEmbarcacion  int       `json:"id_embarcacion" validate:"required"`
	IDHorario      int       `json:"id_horario" validate:"required"`
	Fecha          time.Time `json:"fecha" validate:"required"`
	CupoMaximo     int       `json:"cupo_maximo" validate:"required,min=1"`
	CupoDisponible int       `json:"cupo_disponible" validate:"required,min=0"`
	Estado         string    `json:"estado" validate:"required,oneof=PROGRAMADO COMPLETADO CANCELADO"`
}

// CambiarEstadoTourRequest representa los datos para cambiar el estado de un tour programado
type CambiarEstadoTourRequest struct {
	Estado string `json:"estado" validate:"required,oneof=PROGRAMADO COMPLETADO CANCELADO"`
}
