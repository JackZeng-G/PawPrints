package bootstrap

import (
	"pawprints-server/internal/seed"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	return seed.SeedPetData(db)
}
