package model

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
	"time"
)

type DateOnly time.Time

func (d DateOnly) Value() (driver.Value, error) {
	t := time.Time(d)
	if t.IsZero() {
		return nil, nil
	}
	return t.Format("2006-01-02"), nil
}

func (d *DateOnly) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case string:
		t, err := time.Parse("2006-01-02", v)
		if err != nil {
			return err
		}
		*d = DateOnly(t)
	case []byte:
		t, err := time.Parse("2006-01-02", string(v))
		if err != nil {
			return err
		}
		t, err = time.Parse(time.RFC3339Nano, string(v))
		if err == nil {
			*d = DateOnly(t)
		}
	case time.Time:
		*d = DateOnly(v)
	}
	return nil
}

func (d *DateOnly) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" || s == "null" {
		return nil
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		t, err = time.Parse(time.RFC3339, s)
	}
	if err != nil {
		return err
	}
	*d = DateOnly(t)
	return nil
}

func (d DateOnly) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	if t.IsZero() {
		return json.Marshal(nil)
	}
	return json.Marshal(t.Format("2006-01-02"))
}

type Pet struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"size:100;not null" json:"name"`
	CategoryID  uint       `gorm:"not null;index" json:"category_id"`
	BreedID     uint       `gorm:"index" json:"breed_id"`
	Gender      string     `gorm:"size:20;default:'unknown'" json:"gender"`
	Birthday    DateOnly   `gorm:"not null" json:"birthday"`
	PassingDate *DateOnly  `json:"passing_date"`
	AvatarURL   string     `gorm:"size:255" json:"avatar_url"`
	Bio         string     `gorm:"type:text" json:"bio"`
	Status      string     `gorm:"size:20;default:'alive'" json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`

	Category PetCategory `gorm:"foreignKey:CategoryID" json:"category"`
	Breed    *PetBreed   `gorm:"foreignKey:BreedID" json:"breed"`
}

func (Pet) TableName() string {
	return "pets"
}
