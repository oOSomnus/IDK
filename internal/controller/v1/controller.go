package v1

import (
	"I_Dev_Kit/internal/usecase"
	"I_Dev_Kit/pkg/logger"

	"github.com/go-playground/validator/v10"
)

// V1 -.
type V1 struct {
	// p usecase.Project
	f usecase.Feature
	l logger.Interface
	v *validator.Validate
}

func New(f usecase.Feature, l logger.Interface) *V1 {
	return &V1{
		// p: p,
		f: f,
		l: l,
		v: validator.New(),
	}
}
