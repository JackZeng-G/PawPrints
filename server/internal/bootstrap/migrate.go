package bootstrap

import (
	"pawprints-server/internal/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.PetCategory{},
		&model.PetBreed{},
		&model.Pet{},
		&model.DiaryEntry{},
		&model.DiaryPet{},
		&model.DiaryPhoto{},
		&model.HealthRecord{},
		&model.Reminder{},
	)
}
