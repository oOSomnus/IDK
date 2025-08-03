package entity

type Decision struct {
	ID        int64  `json:"id" gorm:"primaryKey"`
	Situation string `json:"situation" gorm:"not null"`
	Options   string `json:"options" gorm:"not null"`
	Result    string `json:"result" gorm:"not null"`
	Details   string `json:"details" gorm:"type:text"`
	FeatureID int64  `json:"feature_id" gorm:"not null"`
}
