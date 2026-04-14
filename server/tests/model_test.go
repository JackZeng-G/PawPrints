package tests

import (
	"testing"
	"time"
	"pawprints-server/internal/model"
)

func TestModel_PetCategory(t *testing.T) {
	cat := model.PetCategory{Name: "猫", Icon: "cat", SortOrder: 1}
	if cat.Name != "猫" {
		t.Error("Category name should be set")
	}
}

func TestModel_PetBreed(t *testing.T) {
	breed := model.PetBreed{CategoryID: 1, Name: "布偶猫", ExpectedLifespan: "12-15年", Size: "中型"}
	if breed.Name != "布偶猫" {
		t.Error("Breed name should be set")
	}
}

func TestModel_Pet(t *testing.T) {
	pet := model.Pet{
		Name:       "小橘",
		CategoryID: 1,
		Gender:     "male",
		Birthday:   model.DateOnly(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
		Bio:        "一只可爱的小橘猫",
		Status:     "alive",
	}
	if pet.Name != "小橘" {
		t.Error("Pet name should be set")
	}
	if pet.Status != "alive" {
		t.Error("Default status should be alive")
	}
}

func TestModel_DiaryEntry(t *testing.T) {
	diary := model.DiaryEntry{Title: "今天很开心", Content: "小橘今天学会了跳高", EntryDate: model.DateOnly(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)), Mood: "happy"}
	if diary.Title != "今天很开心" {
		t.Error("Diary title should be set")
	}
}

func TestModel_HealthRecord(t *testing.T) {
	hr := model.HealthRecord{PetID: 1, Type: "vaccine", Title: "狂犬疫苗", RecordDate: model.DateOnly(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))}
	if hr.Type != "vaccine" {
		t.Error("Health record type should be set")
	}
}

func TestModel_Reminder(t *testing.T) {
	rem := model.Reminder{Title: "疫苗接种提醒", RemindAt: model.DateTime(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, 7)), Type: "vaccine_due", PetID: 1, Completed: false}
	if rem.Title != "疫苗接种提醒" {
		t.Error("Reminder title should be set")
	}
	if rem.Completed {
		t.Error("New reminder should not be completed")
	}
}