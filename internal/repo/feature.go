package repo

import (
	"I_Dev_Kit/internal/entity"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type FeatureRepository interface {
	GetByState(state entity.FeatureState) ([]entity.Feature, error)
	GetOverdue() ([]entity.Feature, error)
}

type featureRepo struct {
	db *gorm.DB
}

func NewFeatureRepo(db *gorm.DB) *featureRepo {
	return &featureRepo{db: db}
}

// GetByState returns all features with the specified state w/o foreign keys
func (r *featureRepo) GetByState(state entity.FeatureState) ([]entity.Feature, error) {
	var features []entity.Feature
	result := r.db.Where("state = ?", state).Find(&features)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get features by state: %w", result.Error)
	}
	return features, nil
}

// GetOverdue return all features that are overdue w/o foreign keys
func (r *featureRepo) GetOverdue() ([]entity.Feature, error) {
	var features []entity.Feature
	now := time.Now()
	result := r.db.Where("deadline < ?", now).Find(&features)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get overdue features: %w", result.Error)
	}
	return features, nil
}
