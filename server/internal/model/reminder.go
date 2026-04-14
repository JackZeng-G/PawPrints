package model

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
	"time"
)

type DateTime time.Time

func (d DateTime) Value() (driver.Value, error) {
	t := time.Time(d)
	if t.IsZero() {
		return nil, nil
	}
	return t, nil
}

func (d *DateTime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*d = DateTime(v)
	}
	return nil
}

func (d *DateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" || s == "null" {
		return nil
	}
	t, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		t, err = time.Parse(time.RFC3339, s)
	}
	if err != nil {
		return err
	}
	*d = DateTime(t)
	return nil
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	if t.IsZero() {
		return json.Marshal(nil)
	}
	return json.Marshal(t.Format(time.RFC3339))
}

type Reminder struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	Title          string     `gorm:"size:200;not null" json:"title"`
	Description    string     `gorm:"type:text" json:"description"`
	RemindAt       DateTime   `gorm:"not null;index" json:"remind_at"`
	Type           string     `gorm:"size:30;index" json:"type"`
	PetID          uint       `gorm:"index" json:"pet_id"`
	HealthRecordID uint       `gorm:"index" json:"health_record_id"`
	Completed      bool       `gorm:"default:false;index" json:"completed"`
	CompletedAt    *DateTime  `json:"completed_at"`
	SnoozedUntil   *DateTime  `json:"snoozed_until"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`

	Pet          *Pet          `gorm:"foreignKey:PetID" json:"pet"`
	HealthRecord *HealthRecord `gorm:"foreignKey:HealthRecordID" json:"health_record"`
}

func (Reminder) TableName() string { return "reminders" }
