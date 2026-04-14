package tests

import (
	"testing"
	"pawprints-server/internal/bootstrap"
	"pawprints-server/internal/model"
)

func TestMigrate_AutoMigrate(t *testing.T) {
	app := bootstrap.NewApp()
	defer app.Close()

	tables := []struct {
		name  string
		model interface{}
	}{
		{"pet_categories", &model.PetCategory{}},
		{"pet_breeds", &model.PetBreed{}},
		{"pets", &model.Pet{}},
		{"diary_entries", &model.DiaryEntry{}},
		{"diary_pets", &model.DiaryPet{}},
		{"diary_photos", &model.DiaryPhoto{}},
		{"health_records", &model.HealthRecord{}},
		{"reminders", &model.Reminder{}},
	}

	for _, tt := range tables {
		if !app.DB.Migrator().HasTable(tt.model) {
			t.Errorf("%s table should exist", tt.name)
		}
	}
}
