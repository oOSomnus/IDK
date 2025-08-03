package entity

type Diagram struct {
	ID        int64       `json:"id" gorm:"primaryKey"`
	Data      string      `json:"data" gorm:"type:text"`
	URL       string      `json:"url" gorm:"type:text"`
	Type      DiagramType `json:"type" gorm:"type:varchar(20);not null"`
	FeatureID int64       `json:"feature_id" gorm:"not null"`
}

type DiagramType string

const (
	SequenceDiagram     DiagramType = "sequence"
	ClassDiagram        DiagramType = "class"
	StateMachineDiagram DiagramType = "state_machine"
)
