package model

import "time"

type Reminder struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	Title          string     `gorm:"size:200;not null" json:"title"`
	Description    string     `gorm:"type:text" json:"description"`
	RemindAt       time.Time  `gorm:"not null;index" json:"remind_at"`
	Type           string     `gorm:"size:30;index" json:"type"`
	PetID          uint       `gorm:"index" json:"pet_id"`
	HealthRecordID uint       `gorm:"index" json:"health_record_id"`
	Completed      bool       `gorm:"default:false;index" json:"completed"`
	CompletedAt    *time.Time `json:"completed_at"`
	SnoozedUntil   *time.Time `json:"snoozed_until"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`

	Pet          *Pet          `gorm:"foreignKey:PetID" json:"pet"`
	HealthRecord *HealthRecord `gorm:"foreignKey:HealthRecordID" json:"health_record"`
}

func (Reminder) TableName() string { return "reminders" }
