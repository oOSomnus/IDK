package entity

import "time"

type Project struct {
	ID          int64        `json:"id" gorm:"primaryKey"`
	URL         string       `json:"url" gorm:"type:text"`
	Title       string       `json:"title" gorm:"not null"`
	Description string       `json:"description" gorm:"type:text"`
	TeamMembers int          `json:"team_members" gorm:"default:0"`
	Progress    int          `json:"progress" gorm:"default:0"`
	Status      string       `json:"status" gorm:"default:'active'"`
	Features    []Feature    `json:"features" gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	State       ProjectState `json:"state" gorm:"type:varchar(20);not null"`
	DueDate     time.Time    `json:"due_date" gorm:"type:date"`
	UpdatedAt   time.Time    `json:"update_at" gorm:"autoUpdateTime"`
	CreatedAt   time.Time    `json:"create_at"`
}

type ProjectState string

const (
	ProjectStatePending    ProjectState = "pending"
	ProjectStateInProgress ProjectState = "in_progress"
	ProjectStateCompleted  ProjectState = "completed"
	ProjectStateCancelled  ProjectState = "cancelled"
	ProjectStatePlanned    ProjectState = "planned"
)
