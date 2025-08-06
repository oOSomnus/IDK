package usecase

import (
	"I_Dev_Kit/cmd/web/components"
	"I_Dev_Kit/internal/entity"
)

type (
	Project interface {
		GetProjectsByPage(page int) ([]entity.Project, error)
	}
	Feature interface {
		GetStats() (components.QuickStats, error)
	}
)
