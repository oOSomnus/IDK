package feature

import (
	"I_Dev_Kit/cmd/web/components"
	"I_Dev_Kit/internal/entity"
	"I_Dev_Kit/internal/repo"
	"fmt"
)

type UseCase struct {
	repo repo.FeatureRepository
}

func New(r repo.FeatureRepository) *UseCase {
	return &UseCase{repo: r}
}

func (u *UseCase) GetStats() (components.QuickStats, error) {
	p, err := u.repo.GetByState(entity.FeatureStatePlanned)
	if err != nil {
		return components.QuickStats{}, fmt.Errorf("failed to get planned features: %w", err)
	}
	c, err := u.repo.GetByState(entity.FeatureStateCompleted)
	if err != nil {
		return components.QuickStats{}, fmt.Errorf("failed to get completed features: %w",
			err)
	}
	ip, err := u.repo.GetByState(entity.FeatureStateInProgress)
	if err != nil {
		return components.QuickStats{}, fmt.Errorf("failed to get in-progress features: %w", err)
	}
	overdue, err := u.repo.GetOverdue()
	if err != nil {
		return components.QuickStats{}, fmt.Errorf("failed to get overdue features: %w",
			err)
	}
	return components.QuickStats{
		Planned:     len(p),
		Completed:   len(c),
		Overdue:     len(overdue),
		In_Progress: len(ip),
	}, nil
}
