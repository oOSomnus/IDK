package feature

import (
	"I_Dev_Kit/cmd/web/components"
	"I_Dev_Kit/internal/repo"
)

type UseCase struct {
	repo repo.FeatureRepository
}

func New(r repo.FeatureRepository) *UseCase {
	return &UseCase{repo: r}
}

func (u *UseCase) GetStats() components.QuickStats {
	return components.QuickStats{
		Active:      100,
		Completed:   100,
		Overdue:     10,
		In_Progress: 1000,
	}
}
