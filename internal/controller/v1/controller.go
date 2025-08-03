package v1

import (
	"I_Dev_Kit/internal/usecase"
	"I_Dev_Kit/pkg/logger"

	"github.com/go-playground/validator/v10"
)

// V1 -.
type V1 struct {
	p usecase.Project
	l logger.Interface
	v *validator.Validate
}
