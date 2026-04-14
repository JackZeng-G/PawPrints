package model

type PetBreed struct {
	ID               uint   `gorm:"primaryKey" json:"id"`
	CategoryID       uint   `gorm:"not null;index" json:"category_id"`
	Name             string `gorm:"size:100;not null" json:"name"`
	ExpectedLifespan string `gorm:"size:50" json:"expected_lifespan"`
	Size             string `gorm:"size:20" json:"size"`
}

func (PetBreed) TableName() string {
	return "pet_breeds"
}
