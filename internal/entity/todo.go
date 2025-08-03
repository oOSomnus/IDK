package entity

type Todo struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Status    TodoState `json:"status" gorm:"type:varchar(20);not null"`
	FeatureID int64     `json:"feature_id" gorm:"not null"`
}

type TodoState string

const (
	TodoStatePending    TodoState = "pending"
	TodoStateInProgress TodoState = "in_progress"
	TodoStateCompleted  TodoState = "completed"
	TodoStateCancelled  TodoState = "cancelled"
	TodoStateCreated    TodoState = "created"
)
