-- Tabla de usuarios (administradores, vendedores, chofer)
CREATE TABLE usuario (
    id_usuario SERIAL PRIMARY KEY,
    nombres VARCHAR(100) NOT NULL,
    apellidos VARCHAR(100) NOT NULL,
    correo VARCHAR(100) UNIQUE,
    telefono VARCHAR(20),
    direccion VARCHAR(255),
    fecha_nacimiento DATE,
    rol VARCHAR(20) NOT NULL,
    nacionalidad VARCHAR(50),
    tipo_de_documento VARCHAR(50) NOT NULL,
    numero_documento VARCHAR(20) NOT NULL,
    fecha_registro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    contrasena VARCHAR(255),
    estado BOOLEAN DEFAULT TRUE,
    UNIQUE (numero_documento)
);

-- Tabla de embarcaciones (con usuario chofer relacionada)
CREATE TABLE embarcacion (
    id_embarcacion SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    capacidad INT NOT NULL,
    descripcion VARCHAR(255),
    estado BOOLEAN DEFAULT TRUE,
    id_usuario INT NOT NULL, -- El chofer asignado
    FOREIGN KEY (id_usuario) REFERENCES usuario(id_usuario)
);

-- Tabla de Paquete_tour (MEJORADA)
CREATE TABLE tipo_tour (
    id_tipo_tour SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    descripcion VARCHAR(255),
    duracion_minutos INT NOT NULL,
    precio_base DECIMAL(10,2) NOT NULL,
    cantidad_pasajeros INT NOT NULL, -- Cambiado de VARCHAR a INT
    url_imagen VARCHAR(255) -- URL o ruta a la imagen del tour
);

-- Tabla de horarios de tour (MEJORADA)
CREATE TABLE horario_tour (
    id_horario SERIAL PRIMARY KEY,
    id_tipo_tour INT NOT NULL,
    hora_inicio TIME NOT NULL,
    hora_fin TIME NOT NULL,
    -- Reemplazado dias_disponible VARCHAR por campos individuales
    disponible_lunes BOOLEAN DEFAULT FALSE,
    disponible_martes BOOLEAN DEFAULT FALSE,
    disponible_miercoles BOOLEAN DEFAULT FALSE,
    disponible_jueves BOOLEAN DEFAULT FALSE,
    disponible_viernes BOOLEAN DEFAULT FALSE,
    disponible_sabado BOOLEAN DEFAULT FALSE,
    disponible_domingo BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (id_tipo_tour) REFERENCES tipo_tour(id_tipo_tour)
);

-- Tabla de horarios de choferes (NUEVA)
CREATE TABLE horario_chofer (
    id_horario_chofer SERIAL PRIMARY KEY,
    id_usuario INT NOT NULL, -- El chofer
    hora_inicio TIME NOT NULL, -- Hora de inicio del turno
    hora_fin TIME NOT NULL, -- Hora de fin del turno
    disponible_lunes BOOLEAN DEFAULT FALSE,
    disponible_martes BOOLEAN DEFAULT FALSE,
    disponible_miercoles BOOLEAN DEFAULT FALSE,
    disponible_jueves BOOLEAN DEFAULT FALSE,
    disponible_viernes BOOLEAN DEFAULT FALSE,
    disponible_sabado BOOLEAN DEFAULT FALSE,
    disponible_domingo BOOLEAN DEFAULT FALSE,
    fecha_inicio DATE NOT NULL, -- Desde cuándo aplica este horario
    fecha_fin DATE, -- Hasta cuándo (NULL si es indefinido)
    FOREIGN KEY (id_usuario) REFERENCES usuario(id_usuario)
);

-- Tabla de tours programados
CREATE TABLE tour_programado (
    id_tour_programado SERIAL PRIMARY KEY,
    id_tipo_tour INT NOT NULL,
    id_embarcacion INT NOT NULL,
    id_horario INT NOT NULL,
    fecha DATE NOT NULL,
    cupo_maximo INT NOT NULL,
    cupo_disponible INT NOT NULL,
    estado VARCHAR(20) DEFAULT 'PROGRAMADO', -- PROGRAMADO, COMPLETADO, CANCELADO
    FOREIGN KEY (id_tipo_tour) REFERENCES tipo_tour(id_tipo_tour),
    FOREIGN KEY (id_embarcacion) REFERENCES embarcacion(id_embarcacion),
    FOREIGN KEY (id_horario) REFERENCES horario_tour(id_horario),
    UNIQUE (id_embarcacion, fecha, id_horario)
);

-- Tabla de métodos de pago
CREATE TABLE metodo_pago (
    id_metodo_pago SERIAL PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL,
    descripcion VARCHAR(255)
);

-- Tabla de canales de venta
CREATE TABLE canal_venta (
    id_canal SERIAL PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL, -- WEB, LOCAL
    descripcion VARCHAR(255)
);

-- Tabla de reservas (Un cliente puede tener MÚLTIPLES reservas)
CREATE TABLE reserva (
    id_reserva SERIAL PRIMARY KEY,
    id_usuario INT NOT NULL, -- Cliente que hace la reserva
    id_tour_programado INT NOT NULL,
    id_canal INT NOT NULL,
    id_vendedor INT, -- Vendedor que registra la reserva (NULL si es web)
    fecha_reserva TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    cantidad_adultos INT NOT NULL DEFAULT 0,
    cantidad_ninos INT NOT NULL DEFAULT 0,
    total_pagar DECIMAL(10,2) NOT NULL,
    notas TEXT,
    estado VARCHAR(20) DEFAULT 'RESERVADO', -- RESERVADO, CANCELADA
    FOREIGN KEY (id_usuario) REFERENCES usuario(id_usuario),
    FOREIGN KEY (id_tour_programado) REFERENCES tour_programado(id_tour_programado),
    FOREIGN KEY (id_canal) REFERENCES canal_venta(id_canal),
    FOREIGN KEY (id_vendedor) REFERENCES usuario(id_usuario)
);

-- Tabla de tipo de pasaje
CREATE TABLE tipo_pasaje (
    id_tipo_pasaje SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    costo DECIMAL(10,2) NOT NULL,
    edad VARCHAR(50)
);

-- Tabla de pasajeros (adicionales a quien reserva)
CREATE TABLE pasajero (
    id_pasajero SERIAL PRIMARY KEY,
    id_reserva INT NOT NULL,
    tipo_documento VARCHAR(50) NOT NULL,
    numero_documento VARCHAR(20) NOT NULL,
    nombres VARCHAR(100) NOT NULL,
    apellidos VARCHAR(100) NOT NULL,
    id_tipo_pasaje INT NOT NULL, -- Añadido el tipo de pasaje
    FOREIGN KEY (id_reserva) REFERENCES reserva(id_reserva),
    FOREIGN KEY (id_tipo_pasaje) REFERENCES tipo_pasaje(id_tipo_pasaje)
);

-- Tabla de pagos
CREATE TABLE pago (
    id_pago SERIAL PRIMARY KEY,
    id_reserva INT NOT NULL,
    id_metodo_pago INT NOT NULL,
    id_canal INT NOT NULL, -- Añadido el canal de venta
    monto DECIMAL(10,2) NOT NULL,
    fecha_pago TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    comprobante VARCHAR(100), -- Número de comprobante o transacción
    estado VARCHAR(20) DEFAULT 'PROCESADO', -- PROCESADO, ANULADO
    FOREIGN KEY (id_reserva) REFERENCES reserva(id_reserva),
    FOREIGN KEY (id_metodo_pago) REFERENCES metodo_pago(id_metodo_pago),
    FOREIGN KEY (id_canal) REFERENCES canal_venta(id_canal)
);

-- Tabla de comprobantes de pago
CREATE TABLE comprobante_pago (
    id_comprobante SERIAL PRIMARY KEY,
    id_reserva INT NOT NULL,
    tipo VARCHAR(20) NOT NULL, -- BOLETA, FACTURA
    numero_comprobante VARCHAR(20) NOT NULL,
    fecha_emision TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    subtotal DECIMAL(10,2) NOT NULL,
    igv DECIMAL(10,2) NOT NULL,
    total DECIMAL(10,2) NOT NULL,
    estado VARCHAR(20) DEFAULT 'EMITIDO', -- EMITIDO, ANULADO
    FOREIGN KEY (id_reserva) REFERENCES reserva(id_reserva),
    UNIQUE (tipo, numero_comprobante)
);