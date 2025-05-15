package servicios

import (
	"errors"
	"sistema-tours/internal/entidades"
	"sistema-tours/internal/repositorios"
	"sistema-tours/internal/utils"
)

// UsuarioService maneja la lógica de negocio para usuarios
type UsuarioService struct {
	usuarioRepo *repositorios.UsuarioRepository
}

// NewUsuarioService crea una nueva instancia de UsuarioService
func NewUsuarioService(usuarioRepo *repositorios.UsuarioRepository) *UsuarioService {
	return &UsuarioService{
		usuarioRepo: usuarioRepo,
	}
}

// Create crea un nuevo usuario
func (s *UsuarioService) Create(usuario *entidades.NuevoUsuarioRequest) (int, error) {
	// Verificar si ya existe usuario con el mismo correo
	existingEmail, err := s.usuarioRepo.GetByEmail(usuario.Correo)
	if err == nil && existingEmail != nil {
		return 0, errors.New("ya existe un usuario con ese correo electrónico")
	}

	// Verificar si ya existe usuario con el mismo documento
	existingDoc, err := s.usuarioRepo.GetByDocumento(usuario.TipoDocumento, usuario.NumeroDocumento)
	if err == nil && existingDoc != nil {
		return 0, errors.New("ya existe un usuario con ese documento")
	}

	// Hash de la contraseña
	hashedPassword, err := utils.HashPassword(usuario.Contrasena)
	if err != nil {
		return 0, err
	}

	// Crear usuario
	return s.usuarioRepo.Create(usuario, hashedPassword)
}

// GetByID obtiene un usuario por su ID
func (s *UsuarioService) GetByID(id int) (*entidades.Usuario, error) {
	return s.usuarioRepo.GetByID(id)
}

// Update actualiza un usuario existente
func (s *UsuarioService) Update(id int, usuario *entidades.Usuario) error {
	// Verificar que el usuario existe
	existing, err := s.usuarioRepo.GetByID(id)
	if err != nil {
		return err
	}

	// Verificar si ya existe otro usuario con el mismo correo
	if usuario.Correo != existing.Correo {
		existingEmail, err := s.usuarioRepo.GetByEmail(usuario.Correo)
		if err == nil && existingEmail != nil && existingEmail.ID != id {
			return errors.New("ya existe otro usuario con ese correo electrónico")
		}
	}

	// Verificar si ya existe otro usuario con el mismo documento
	if usuario.NumeroDocumento != existing.NumeroDocumento || usuario.TipoDocumento != existing.TipoDocumento {
		existingDoc, err := s.usuarioRepo.GetByDocumento(usuario.TipoDocumento, usuario.NumeroDocumento)
		if err == nil && existingDoc != nil && existingDoc.ID != id {
			return errors.New("ya existe otro usuario con ese documento")
		}
	}

	// Actualizar ID para asegurar que sea el correcto
	usuario.ID = id

	// Actualizar usuario
	return s.usuarioRepo.Update(usuario)
}

// Delete elimina un usuario (borrado lógico)
func (s *UsuarioService) Delete(id int) error {
	// Verificar que el usuario existe
	_, err := s.usuarioRepo.GetByID(id)
	if err != nil {
		return err
	}

	// Eliminar usuario
	return s.usuarioRepo.Delete(id)
}

// ListByRol lista usuarios por rol
func (s *UsuarioService) ListByRol(rol string) ([]*entidades.Usuario, error) {
	return s.usuarioRepo.ListByRol(rol)
}

// List lista todos los usuarios
func (s *UsuarioService) List() ([]*entidades.Usuario, error) {
	return s.usuarioRepo.List()
}
