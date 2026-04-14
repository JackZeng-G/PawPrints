package model

import "time"

type HealthRecord struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	PetID       uint       `gorm:"not null;index" json:"pet_id"`
	Type        string     `gorm:"size:30;not null;index" json:"type"`
	Title       string     `gorm:"size:200;not null" json:"title"`
	Value       float64    `json:"value"`
	Unit        string     `gorm:"size:20" json:"unit"`
	Notes       string     `gorm:"type:text" json:"notes"`
	RecordDate  time.Time  `gorm:"not null;index" json:"record_date"`
	NextDueDate *time.Time `gorm:"index" json:"next_due_date"`
	VetName     string     `gorm:"size:100" json:"vet_name"`
	ClinicName   string     `gorm:"size:200" json:"clinic_name"`
	Medication  string     `gorm:"type:text" json:"medication"`
	Diagnosis   string     `gorm:"type:text" json:"diagnosis"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`

	Pet Pet `gorm:"foreignKey:PetID" json:"pet"`
}

func (HealthRecord) TableName() string { return "health_records" }
