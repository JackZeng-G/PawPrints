package model

import "time"

type Pet struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"size:100;not null" json:"name"`
	CategoryID  uint       `gorm:"not null;index" json:"category_id"`
	BreedID     uint       `gorm:"index" json:"breed_id"`
	Gender      string     `gorm:"size:20;default:'unknown'" json:"gender"`
	Birthday    time.Time  `gorm:"not null" json:"birthday"`
	PassingDate *time.Time `json:"passing_date"`
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
