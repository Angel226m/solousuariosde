package entidades

import "time"

// Pago representa la estructura de un pago en el sistema
type Pago struct {
	ID           int       `json:"id_pago" db:"id_pago"`
	IDReserva    int       `json:"id_reserva" db:"id_reserva"`
	IDMetodoPago int       `json:"id_metodo_pago" db:"id_metodo_pago"`
	IDCanal      int       `json:"id_canal" db:"id_canal"`
	Monto        float64   `json:"monto" db:"monto"`
	FechaPago    time.Time `json:"fecha_pago" db:"fecha_pago"`
	Comprobante  string    `json:"comprobante" db:"comprobante"`
	Estado       string    `json:"estado" db:"estado"`

	// Campos adicionales para mostrar información relacionada
	NumeroReserva    int    `json:"numero_reserva,omitempty" db:"-"`
	NombreCliente    string `json:"nombre_cliente,omitempty" db:"-"`
	ApellidosCliente string `json:"apellidos_cliente,omitempty" db:"-"`
	DocumentoCliente string `json:"documento_cliente,omitempty" db:"-"`
	NombreMetodoPago string `json:"nombre_metodo_pago,omitempty" db:"-"`
	NombreCanalVenta string `json:"nombre_canal_venta,omitempty" db:"-"`
}

// NuevoPagoRequest representa los datos necesarios para crear un nuevo pago
type NuevoPagoRequest struct {
	IDReserva    int     `json:"id_reserva" validate:"required"`
	IDMetodoPago int     `json:"id_metodo_pago" validate:"required"`
	IDCanal      int     `json:"id_canal" validate:"required"`
	Monto        float64 `json:"monto" validate:"required,min=0"`
	Comprobante  string  `json:"comprobante"`
}

// ActualizarPagoRequest representa los datos para actualizar un pago
type ActualizarPagoRequest struct {
	IDReserva    int     `json:"id_reserva" validate:"required"`
	IDMetodoPago int     `json:"id_metodo_pago" validate:"required"`
	IDCanal      int     `json:"id_canal" validate:"required"`
	Monto        float64 `json:"monto" validate:"required,min=0"`
	Comprobante  string  `json:"comprobante"`
	Estado       string  `json:"estado" validate:"required,oneof=PROCESADO ANULADO"`
}

// CambiarEstadoPagoRequest representa los datos para cambiar el estado de un pago
type CambiarEstadoPagoRequest struct {
	Estado string `json:"estado" validate:"required,oneof=PROCESADO ANULADO"`
}
