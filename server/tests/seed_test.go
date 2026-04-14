package tests

import (
	"testing"
	"pawprints-server/internal/bootstrap"
	"pawprints-server/internal/model"
)

func TestSeed_PetCategories(t *testing.T) {
	app := bootstrap.NewApp()
	defer app.Close()

	var categories []model.PetCategory
	app.DB.Find(&categories)

	if len(categories) < 7 {
		t.Errorf("Should have at least 7 categories, got %d", len(categories))
	}

	var cat model.PetCategory
	app.DB.Where("name = ?", "猫").First(&cat)
	if cat.Name != "猫" {
		t.Error("Cat category should exist")
	}
}

func TestSeed_PetBreeds(t *testing.T) {
	app := bootstrap.NewApp()
	defer app.Close()

	var breeds []model.PetBreed
	app.DB.Find(&breeds)

	if len(breeds) < 20 {
		t.Errorf("Should have at least 20 breeds, got %d", len(breeds))
	}

	var ragdoll model.PetBreed
	app.DB.Where("name = ?", "布偶猫").First(&ragdoll)
	if ragdoll.Name != "布偶猫" {
		t.Error("Ragdoll breed should exist")
	}
}
