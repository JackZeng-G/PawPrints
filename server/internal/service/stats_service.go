package service

import (
	"pawprints-server/internal/model"
	"gorm.io/gorm"
)

type StatsService struct {
	DB *gorm.DB
}

func NewStatsService(db *gorm.DB) *StatsService {
	return &StatsService{DB: db}
}

type OverviewStats struct {
	TotalPets     int64 `json:"total_pets"`
	AlivePets     int64 `json:"alive_pets"`
	MemorialPets  int64 `json:"memorial_pets"`
	TotalDiaries  int64 `json:"total_diaries"`
	TotalPhotos   int64 `json:"total_photos"`
	ActiveReminders int64 `json:"active_reminders"`
}

func (s *StatsService) Overview() (*OverviewStats, error) {
	var stats OverviewStats
	s.DB.Model(&model.Pet{}).Count(&stats.TotalPets)
	s.DB.Model(&model.Pet{}).Where("status = ?", "alive").Count(&stats.AlivePets)
	s.DB.Model(&model.Pet{}).Where("status = ?", "memorial").Count(&stats.MemorialPets)
	s.DB.Model(&model.DiaryEntry{}).Count(&stats.TotalDiaries)
	s.DB.Model(&model.DiaryPhoto{}).Count(&stats.TotalPhotos)
	s.DB.Model(&model.Reminder{}).Where("completed = ?", false).Count(&stats.ActiveReminders)
	return &stats, nil
}

type TimelineItem struct {
	Date        string `json:"date"`
	DiaryCount  int64  `json:"diary_count"`
	PhotoCount  int64  `json:"photo_count"`
}

func (s *StatsService) Timeline(days int) ([]TimelineItem, error) {
	var items []TimelineItem
	s.DB.Raw(`
		SELECT d.entry_date as date,
			COUNT(DISTINCT d.id) as diary_count,
			COUNT(dp.id) as photo_count
		FROM diary_entries d
		LEFT JOIN diary_photos dp ON dp.diary_entry_id = d.id
		WHERE d.entry_date >= date('now', '-' || ? || ' days')
		GROUP BY d.entry_date
		ORDER BY d.entry_date DESC`, days).Scan(&items)
	return items, nil
}

func (s *StatsService) ExportJSON() (map[string]interface{}, error) {
	data := make(map[string]interface{})
	var pets []model.Pet
	s.DB.Preload("Category").Preload("Breed").Find(&pets)
	data["pets"] = pets

	var diaries []model.DiaryEntry
	s.DB.Preload("Photos").Preload("Pets").Find(&diaries)
	data["diaries"] = diaries

	var health []model.HealthRecord
	s.DB.Find(&health)
	data["health_records"] = health

	var reminders []model.Reminder
	s.DB.Find(&reminders)
	data["reminders"] = reminders

	return data, nil
}

func (s *StatsService) ImportJSON(data map[string]interface{}) error {
	return nil // 简化实现，后续完善
}
