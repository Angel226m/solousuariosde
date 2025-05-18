package entidades

import "time"

// ComprobantePago representa la estructura de un comprobante de pago en el sistema
type ComprobantePago struct {
	ID                int       `json:"id_comprobante" db:"id_comprobante"`
	IDReserva         int       `json:"id_reserva" db:"id_reserva"`
	Tipo              string    `json:"tipo" db:"tipo"`
	NumeroComprobante string    `json:"numero_comprobante" db:"numero_comprobante"`
	FechaEmision      time.Time `json:"fecha_emision" db:"fecha_emision"`
	Subtotal          float64   `json:"subtotal" db:"subtotal"`
	IGV               float64   `json:"igv" db:"igv"`
	Total             float64   `json:"total" db:"total"`
	Estado            string    `json:"estado" db:"estado"`

	// Campos adicionales para mostrar información relacionada
	NombreCliente    string    `json:"nombre_cliente,omitempty" db:"-"`
	ApellidosCliente string    `json:"apellidos_cliente,omitempty" db:"-"`
	DocumentoCliente string    `json:"documento_cliente,omitempty" db:"-"`
	NombreTour       string    `json:"nombre_tour,omitempty" db:"-"`
	FechaTour        time.Time `json:"fecha_tour,omitempty" db:"-"`
}

// NuevoComprobantePagoRequest representa los datos necesarios para crear un nuevo comprobante de pago
type NuevoComprobantePagoRequest struct {
	IDReserva         int     `json:"id_reserva" validate:"required"`
	Tipo              string  `json:"tipo" validate:"required,oneof=BOLETA FACTURA"`
	NumeroComprobante string  `json:"numero_comprobante" validate:"required"`
	Subtotal          float64 `json:"subtotal" validate:"required,min=0"`
	IGV               float64 `json:"igv" validate:"required,min=0"`
	Total             float64 `json:"total" validate:"required,min=0"`
}

// ActualizarComprobantePagoRequest representa los datos para actualizar un comprobante de pago
type ActualizarComprobantePagoRequest struct {
	IDReserva         int     `json:"id_reserva" validate:"required"`
	Tipo              string  `json:"tipo" validate:"required,oneof=BOLETA FACTURA"`
	NumeroComprobante string  `json:"numero_comprobante" validate:"required"`
	Subtotal          float64 `json:"subtotal" validate:"required,min=0"`
	IGV               float64 `json:"igv" validate:"required,min=0"`
	Total             float64 `json:"total" validate:"required,min=0"`
	Estado            string  `json:"estado" validate:"required,oneof=EMITIDO ANULADO"`
}

// CambiarEstadoComprobanteRequest representa los datos para cambiar el estado de un comprobante
type CambiarEstadoComprobanteRequest struct {
	Estado string `json:"estado" validate:"required,oneof=EMITIDO ANULADO"`
}
