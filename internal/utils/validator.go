package utils

import (
	"fmt" // Add this import
	"reflect"
	"strings"

	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	es_translations "github.com/go-playground/validator/v10/translations/es"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

// ValidationError representa un error de validación
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Error implementa la interfaz error para ValidationError
func (v *ValidationError) Error() string {
	return fmt.Sprintf("Campo '%s': %s", v.Field, v.Message)
}

// InitValidator inicializa el validador con traducciones al español
func InitValidator() {
	// Crear validador
	validate = validator.New()

	// Configurar para usar nombres de JSON
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// Configurar traductor español
	spanish := es.New()
	uni := ut.New(spanish, spanish)
	trans, _ = uni.GetTranslator("es")

	// Registrar traducciones
	es_translations.RegisterDefaultTranslations(validate, trans)
}

// ValidateStruct valida una estructura utilizando etiquetas de validación
func ValidateStruct(s interface{}) error {
	// Inicializar validador si no se ha hecho
	if validate == nil {
		InitValidator()
	}

	return validate.Struct(s)
}

// FormatValidationErrors formatea errores de validación para una respuesta amigable
func FormatValidationErrors(err error) []ValidationError {
	if err == nil {
		return nil
	}

	var errors []ValidationError
	validationErrors := err.(validator.ValidationErrors)

	for _, e := range validationErrors {
		errors = append(errors, ValidationError{
			Field:   e.Field(),
			Message: e.Translate(trans),
		})
	}

	return errors
}
