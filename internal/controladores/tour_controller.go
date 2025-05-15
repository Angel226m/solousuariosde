package controladores

import (
	"net/http"
	"sistema-tours/internal/entidades"
	"sistema-tours/internal/servicios"
	"sistema-tours/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TourController maneja los endpoints de tipos de tour y horarios
type TourController struct {
	tourService *servicios.TourService
}

// NewTourController crea una nueva instancia de TourController
func NewTourController(tourService *servicios.TourService) *TourController {
	return &TourController{
		tourService: tourService,
	}
}

// CreateTipoTour crea un nuevo tipo de tour
func (c *TourController) CreateTipoTour(ctx *gin.Context) {
	var tipoTourReq entidades.NuevoTipoTourRequest

	// Parsear request
	if err := ctx.ShouldBindJSON(&tipoTourReq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Datos inválidos", err))
		return
	}

	// Validar datos
	if err := utils.ValidateStruct(tipoTourReq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Error de validación", err))
		return
	}

	// Crear tipo de tour
	id, err := c.tourService.CreateTipoTour(&tipoTourReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Error al crear tipo de tour", err))
		return
	}

	// Respuesta exitosa
	ctx.JSON(http.StatusCreated, utils.SuccessResponse("Tipo de tour creado exitosamente", gin.H{"id": id}))
}

// GetTipoTourByID obtiene un tipo de tour por su ID
func (c *TourController) GetTipoTourByID(ctx *gin.Context) {
	// Parsear ID de la URL
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("ID inválido", err))
		return
	}

	// Obtener tipo de tour
	tipoTour, err := c.tourService.GetTipoTourByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Tipo de tour no encontrado", err))
		return
	}

	// Respuesta exitosa
	ctx.JSON(http.StatusOK, utils.SuccessResponse("Tipo de tour obtenido", tipoTour))
}

// UpdateTipoTour actualiza un tipo de tour
func (c *TourController) UpdateTipoTour(ctx *gin.Context) {
	// Parsear ID de la URL
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("ID inválido", err))
		return
	}

	var tipoTourReq entidades.ActualizarTipoTourRequest

	// Parsear request
	if err := ctx.ShouldBindJSON(&tipoTourReq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Datos inválidos", err))
		return
	}

	// Validar datos
	if err := utils.ValidateStruct(tipoTourReq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Error de validación", err))
		return
	}

	// Actualizar tipo de tour
	err = c.tourService.UpdateTipoTour(id, &tipoTourReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Error al actualizar tipo de tour", err))
		return
	}

	// Respuesta exitosa
	ctx.JSON(http.StatusOK, utils.SuccessResponse("Tipo de tour actualizado exitosamente", nil))
}

// DeleteTipoTour elimina un tipo de tour
func (c *TourController) DeleteTipoTour(ctx *gin.Context) {
	// Parsear ID de la URL
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("ID inválido", err))
		return
	}

	// Eliminar tipo de tour
	err = c.tourService.DeleteTipoTour(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Error al eliminar tipo de tour", err))
		return
	}

	// Respuesta exitosa
	ctx.JSON(http.StatusOK, utils.SuccessResponse("Tipo de tour eliminado exitosamente", nil))
}

// ListTiposTour lista todos los tipos de tour
func (c *TourController) ListTiposTour(ctx *gin.Context) {
	// Listar tipos de tour
	tiposTour, err := c.tourService.ListTiposTour()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Error al listar tipos de tour", err))
		return
	}

	// Respuesta exitosa
	ctx.JSON(http.StatusOK, utils.SuccessResponse("Tipos de tour listados exitosamente", tiposTour))
}

// CreateHorario crea un nuevo horario de tour
func (c *TourController) CreateHorario(ctx *gin.Context) {
	var horarioReq entidades.NuevoHorarioTourRequest

	// Parsear request
	if err := ctx.ShouldBindJSON(&horarioReq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Datos inválidos", err))
		return
	}

	// Validar datos
	if err := utils.ValidateStruct(horarioReq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Error de validación", err))
		return
	}

	// Crear horario
	id, err := c.tourService.CreateHorario(&horarioReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Error al crear horario", err))
		return
	}

	// Respuesta exitosa
	ctx.JSON(http.StatusCreated, utils.SuccessResponse("Horario creado exitosamente", gin.H{"id": id}))
}

// GetHorarioByID obtiene un horario por su ID
func (c *TourController) GetHorarioByID(ctx *gin.Context) {
	// Parsear ID de la URL
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("ID inválido", err))
		return
	}

	// Obtener horario
	horario, err := c.tourService.GetHorarioByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse("Horario no encontrado", err))
		return
	}

	// Respuesta exitosa
	ctx.JSON(http.StatusOK, utils.SuccessResponse("Horario obtenido", horario))
}

// UpdateHorario actualiza un horario
func (c *TourController) UpdateHorario(ctx *gin.Context) {
	// Parsear ID de la URL
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("ID inválido", err))
		return
	}

	var horarioReq entidades.ActualizarHorarioTourRequest

	// Parsear request
	if err := ctx.ShouldBindJSON(&horarioReq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Datos inválidos", err))
		return
	}

	// Validar datos
	if err := utils.ValidateStruct(horarioReq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Error de validación", err))
		return
	}

	// Actualizar horario
	err = c.tourService.UpdateHorario(id, &horarioReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Error al actualizar horario", err))
		return
	}

	// Respuesta exitosa
	ctx.JSON(http.StatusOK, utils.SuccessResponse("Horario actualizado exitosamente", nil))
}

// DeleteHorario elimina un horario
func (c *TourController) DeleteHorario(ctx *gin.Context) {
	// Parsear ID de la URL
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("ID inválido", err))
		return
	}

	// Eliminar horario
	err = c.tourService.DeleteHorario(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Error al eliminar horario", err))
		return
	}

	// Respuesta exitosa
	ctx.JSON(http.StatusOK, utils.SuccessResponse("Horario eliminado exitosamente", nil))
}

// ListHorarios lista todos los horarios
func (c *TourController) ListHorarios(ctx *gin.Context) {
	// Listar horarios
	horarios, err := c.tourService.ListHorarios()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse("Error al listar horarios", err))
		return
	}

	// Respuesta exitosa
	ctx.JSON(http.StatusOK, utils.SuccessResponse("Horarios listados exitosamente", horarios))
}

// GetHorariosByTipoTourID obtiene los horarios de un tipo de tour
func (c *TourController) GetHorariosByTipoTourID(ctx *gin.Context) {
	// Parsear ID de la URL
	// Cambiamos de ctx.Param("idTipoTour") a ctx.Param("id")
	idTipoTour, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("ID inválido", err))
		return
	}

	// Obtener horarios
	horarios, err := c.tourService.GetHorariosByTipoTourID(idTipoTour)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse("Error al obtener horarios", err))
		return
	}

	// Respuesta exitosa
	ctx.JSON(http.StatusOK, utils.SuccessResponse("Horarios obtenidos exitosamente", horarios))
}
