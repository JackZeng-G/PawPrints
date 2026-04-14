package model

type PetCategory struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"size:50;not null" json:"name"`
	Icon      string `gorm:"size:50" json:"icon"`
	SortOrder int    `gorm:"default:0" json:"sort_order"`
}

func (PetCategory) TableName() string {
	return "pet_categories"
}
