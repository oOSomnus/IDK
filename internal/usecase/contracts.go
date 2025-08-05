package usecase

import "I_Dev_Kit/cmd/web/components"

type (
	Project interface {
	}
	Feature interface {
		GetStats() components.QuickStats
	}
)
