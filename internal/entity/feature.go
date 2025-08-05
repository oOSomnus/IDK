package entity

type Feature struct {
	ID          int64        `json:"id" gorm:"primaryKey"`
	Title       string       `json:"title" gorm:"not null"`
	Description string       `json:"description" gorm:"type:text"`
	ProjectID   int64        `json:"project_id" gorm:"not null"`
	Decisions   []Decision   `json:"decisions" gorm:"foreignKey:FeatureID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Diagrams    []Diagram    `json:"diagrams" gorm:"foreignKey:FeatureID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Todos       []Todo       `json:"todos" gorm:"foreignKey:FeatureID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	State       FeatureState `json:"state" gorm:"type:varchar(20);not null"`
}

type FeatureState string

const (
	FeatureStatePending    FeatureState = "pending"
	FeatureStateInProgress FeatureState = "in_progress"
	FeatureStateCompleted  FeatureState = "completed"
	FeatureStateCancelled  FeatureState = "cancelled"
	FeatureStateCreated    FeatureState = "created"
)
