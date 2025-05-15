package servicios

import (
	"errors"
	"sistema-tours/internal/entidades"
	"sistema-tours/internal/repositorios"
)

// TourService maneja la lógica de negocio para tipos de tour y horarios
type TourService struct {
	tipoTourRepo *repositorios.TipoTourRepository
}

// NewTourService crea una nueva instancia de TourService
func NewTourService(tipoTourRepo *repositorios.TipoTourRepository) *TourService {
	return &TourService{
		tipoTourRepo: tipoTourRepo,
	}
}

// CreateTipoTour crea un nuevo tipo de tour
func (s *TourService) CreateTipoTour(tipoTour *entidades.NuevoTipoTourRequest) (int, error) {
	// Verificar si ya existe un tipo de tour con el mismo nombre
	existingNombre, err := s.tipoTourRepo.GetByNombre(tipoTour.Nombre)
	if err == nil && existingNombre != nil {
		return 0, errors.New("ya existe un tipo de tour con ese nombre")
	}

	// Validar que la duración sea válida
	if tipoTour.DuracionMinutos <= 0 {
		return 0, errors.New("la duración debe ser mayor a 0 minutos")
	}

	// Validar que el precio base sea válido
	if tipoTour.PrecioBase < 0 {
		return 0, errors.New("el precio base no puede ser negativo")
	}

	// Validar que la cantidad de pasajeros sea válida
	if tipoTour.CantidadPasajeros <= 0 {
		return 0, errors.New("la cantidad de pasajeros debe ser mayor a 0")
	}

	// Crear tipo de tour
	return s.tipoTourRepo.Create(tipoTour)
}

// GetTipoTourByID obtiene un tipo de tour por su ID
func (s *TourService) GetTipoTourByID(id int) (*entidades.TipoTour, error) {
	return s.tipoTourRepo.GetByID(id)
}

// UpdateTipoTour actualiza un tipo de tour existente
func (s *TourService) UpdateTipoTour(id int, tipoTour *entidades.ActualizarTipoTourRequest) error {
	// Verificar que el tipo de tour existe
	existing, err := s.tipoTourRepo.GetByID(id)
	if err != nil {
		return err
	}

	// Verificar si ya existe otro tipo de tour con el mismo nombre
	if tipoTour.Nombre != existing.Nombre {
		existingNombre, err := s.tipoTourRepo.GetByNombre(tipoTour.Nombre)
		if err == nil && existingNombre != nil && existingNombre.ID != id {
			return errors.New("ya existe otro tipo de tour con ese nombre")
		}
	}

	// Validar que la duración sea válida
	if tipoTour.DuracionMinutos <= 0 {
		return errors.New("la duración debe ser mayor a 0 minutos")
	}

	// Validar que el precio base sea válido
	if tipoTour.PrecioBase < 0 {
		return errors.New("el precio base no puede ser negativo")
	}

	// Validar que la cantidad de pasajeros sea válida
	if tipoTour.CantidadPasajeros <= 0 {
		return errors.New("la cantidad de pasajeros debe ser mayor a 0")
	}

	// Actualizar tipo de tour
	return s.tipoTourRepo.Update(id, tipoTour)
}

// DeleteTipoTour elimina un tipo de tour
func (s *TourService) DeleteTipoTour(id int) error {
	// Verificar que el tipo de tour existe
	_, err := s.tipoTourRepo.GetByID(id)
	if err != nil {
		return err
	}

	// Eliminar tipo de tour
	return s.tipoTourRepo.Delete(id)
}

// ListTiposTour lista todos los tipos de tour
func (s *TourService) ListTiposTour() ([]*entidades.TipoTour, error) {
	return s.tipoTourRepo.List()
}

// CreateHorario crea un nuevo horario de tour
func (s *TourService) CreateHorario(horario *entidades.NuevoHorarioTourRequest) (int, error) {
	// Verificar que el tipo de tour existe
	_, err := s.tipoTourRepo.GetByID(horario.IDTipoTour)
	if err != nil {
		return 0, errors.New("el tipo de tour especificado no existe")
	}

	// Verificar que al menos un día esté seleccionado
	if !horario.DisponibleLunes && !horario.DisponibleMartes && !horario.DisponibleMiercoles &&
		!horario.DisponibleJueves && !horario.DisponibleViernes && !horario.DisponibleSabado &&
		!horario.DisponibleDomingo {
		return 0, errors.New("debe seleccionar al menos un día de la semana")
	}

	// Crear horario
	return s.tipoTourRepo.CreateHorario(horario)
}

// GetHorarioByID obtiene un horario por su ID
func (s *TourService) GetHorarioByID(id int) (*entidades.HorarioTour, error) {
	return s.tipoTourRepo.GetHorarioByID(id)
}

// UpdateHorario actualiza un horario de tour
func (s *TourService) UpdateHorario(id int, horario *entidades.ActualizarHorarioTourRequest) error {
	// Verificar que el horario existe
	existing, err := s.tipoTourRepo.GetHorarioByID(id)
	if err != nil {
		return err
	}

	// Verificar que el tipo de tour existe
	if horario.IDTipoTour != existing.IDTipoTour {
		_, err := s.tipoTourRepo.GetByID(horario.IDTipoTour)
		if err != nil {
			return errors.New("el tipo de tour especificado no existe")
		}
	}

	// Verificar que al menos un día esté seleccionado
	if !horario.DisponibleLunes && !horario.DisponibleMartes && !horario.DisponibleMiercoles &&
		!horario.DisponibleJueves && !horario.DisponibleViernes && !horario.DisponibleSabado &&
		!horario.DisponibleDomingo {
		return errors.New("debe seleccionar al menos un día de la semana")
	}

	// Actualizar horario
	return s.tipoTourRepo.UpdateHorario(id, horario)
}

// DeleteHorario elimina un horario de tour
func (s *TourService) DeleteHorario(id int) error {
	// Verificar que el horario existe
	_, err := s.tipoTourRepo.GetHorarioByID(id)
	if err != nil {
		return err
	}

	// Eliminar horario
	return s.tipoTourRepo.DeleteHorario(id)
}

// ListHorarios lista todos los horarios de tour
func (s *TourService) ListHorarios() ([]*entidades.HorarioTour, error) {
	return s.tipoTourRepo.ListHorarios()
}

// GetHorariosByTipoTourID obtiene los horarios de un tipo de tour
func (s *TourService) GetHorariosByTipoTourID(idTipoTour int) ([]*entidades.HorarioTour, error) {
	// Verificar que el tipo de tour existe
	_, err := s.tipoTourRepo.GetByID(idTipoTour)
	if err != nil {
		return nil, err
	}

	// Obtener horarios
	return s.tipoTourRepo.GetHorariosByTipoTourID(idTipoTour)
}
