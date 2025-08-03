package entity

type Project struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	URL         string    `json:"url" gorm:"type:text"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description" gorm:"type:text"`
	Features    []Feature `json:"features" gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
