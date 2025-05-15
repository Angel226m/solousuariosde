package entidades

import "time"

// TipoTour representa la estructura de un tipo de tour en el sistema
type TipoTour struct {
	ID                int     `json:"id_tipo_tour" db:"id_tipo_tour"`
	Nombre            string  `json:"nombre" db:"nombre"`
	Descripcion       string  `json:"descripcion" db:"descripcion"`
	DuracionMinutos   int     `json:"duracion_minutos" db:"duracion_minutos"`
	PrecioBase        float64 `json:"precio_base" db:"precio_base"`
	CantidadPasajeros int     `json:"cantidad_pasajeros" db:"cantidad_pasajeros"`
	URLImagen         string  `json:"url_imagen" db:"url_imagen"`
	// Campos para mostrar información relacionada
	Horarios         []*HorarioTour `json:"horarios,omitempty" db:"-"`
	ToursProgramados int            `json:"tours_programados,omitempty" db:"-"`
}

// NuevoTipoTourRequest representa los datos necesarios para crear un nuevo tipo de tour
type NuevoTipoTourRequest struct {
	Nombre            string  `json:"nombre" validate:"required"`
	Descripcion       string  `json:"descripcion"`
	DuracionMinutos   int     `json:"duracion_minutos" validate:"required,min=1"`
	PrecioBase        float64 `json:"precio_base" validate:"required,min=0"`
	CantidadPasajeros int     `json:"cantidad_pasajeros" validate:"required,min=1"`
	URLImagen         string  `json:"url_imagen"`
}

// ActualizarTipoTourRequest representa los datos para actualizar un tipo de tour
type ActualizarTipoTourRequest struct {
	Nombre            string  `json:"nombre" validate:"required"`
	Descripcion       string  `json:"descripcion"`
	DuracionMinutos   int     `json:"duracion_minutos" validate:"required,min=1"`
	PrecioBase        float64 `json:"precio_base" validate:"required,min=0"`
	CantidadPasajeros int     `json:"cantidad_pasajeros" validate:"required,min=1"`
	URLImagen         string  `json:"url_imagen"`
}

// HorarioTour representa la estructura de un horario de tour
type HorarioTour struct {
	ID                  int       `json:"id_horario" db:"id_horario"`
	IDTipoTour          int       `json:"id_tipo_tour" db:"id_tipo_tour"`
	HoraInicio          time.Time `json:"hora_inicio" db:"hora_inicio"`
	HoraFin             time.Time `json:"hora_fin" db:"hora_fin"`
	DisponibleLunes     bool      `json:"disponible_lunes" db:"disponible_lunes"`
	DisponibleMartes    bool      `json:"disponible_martes" db:"disponible_martes"`
	DisponibleMiercoles bool      `json:"disponible_miercoles" db:"disponible_miercoles"`
	DisponibleJueves    bool      `json:"disponible_jueves" db:"disponible_jueves"`
	DisponibleViernes   bool      `json:"disponible_viernes" db:"disponible_viernes"`
	DisponibleSabado    bool      `json:"disponible_sabado" db:"disponible_sabado"`
	DisponibleDomingo   bool      `json:"disponible_domingo" db:"disponible_domingo"`
	// Campos para mostrar información relacionada
	NombreTipoTour string `json:"nombre_tipo_tour,omitempty" db:"-"`
}

// NuevoHorarioTourRequest representa los datos necesarios para crear un nuevo horario
type NuevoHorarioTourRequest struct {
	IDTipoTour          int    `json:"id_tipo_tour" validate:"required"`
	HoraInicio          string `json:"hora_inicio" validate:"required"`
	HoraFin             string `json:"hora_fin" validate:"required"`
	DisponibleLunes     bool   `json:"disponible_lunes"`
	DisponibleMartes    bool   `json:"disponible_martes"`
	DisponibleMiercoles bool   `json:"disponible_miercoles"`
	DisponibleJueves    bool   `json:"disponible_jueves"`
	DisponibleViernes   bool   `json:"disponible_viernes"`
	DisponibleSabado    bool   `json:"disponible_sabado"`
	DisponibleDomingo   bool   `json:"disponible_domingo"`
}

// ActualizarHorarioTourRequest representa los datos para actualizar un horario
type ActualizarHorarioTourRequest struct {
	IDTipoTour          int    `json:"id_tipo_tour" validate:"required"`
	HoraInicio          string `json:"hora_inicio" validate:"required"`
	HoraFin             string `json:"hora_fin" validate:"required"`
	DisponibleLunes     bool   `json:"disponible_lunes"`
	DisponibleMartes    bool   `json:"disponible_martes"`
	DisponibleMiercoles bool   `json:"disponible_miercoles"`
	DisponibleJueves    bool   `json:"disponible_jueves"`
	DisponibleViernes   bool   `json:"disponible_viernes"`
	DisponibleSabado    bool   `json:"disponible_sabado"`
	DisponibleDomingo   bool   `json:"disponible_domingo"`
}

// ResumenDisponibilidad muestra un resumen de la disponibilidad de un horario
func (h *HorarioTour) ResumenDisponibilidad() string {
	dias := ""
	if h.DisponibleLunes {
		dias += "Lu "
	}
	if h.DisponibleMartes {
		dias += "Ma "
	}
	if h.DisponibleMiercoles {
		dias += "Mi "
	}
	if h.DisponibleJueves {
		dias += "Ju "
	}
	if h.DisponibleViernes {
		dias += "Vi "
	}
	if h.DisponibleSabado {
		dias += "Sa "
	}
	if h.DisponibleDomingo {
		dias += "Do "
	}
	return dias
}
