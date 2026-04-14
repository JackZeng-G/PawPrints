package service

import (
	"encoding/json"
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
	// 填充 pet_name
	petNames := make(map[uint]string)
	for i := range diaries {
		for j := range diaries[i].Pets {
			pid := diaries[i].Pets[j].PetID
			if name, ok := petNames[pid]; ok {
				diaries[i].Pets[j].PetName = name
			} else {
				var name string
				s.DB.Model(&model.Pet{}).Where("id = ?", pid).Select("name").Scan(&name)
				petNames[pid] = name
				diaries[i].Pets[j].PetName = name
			}
		}
	}
	data["diaries"] = diaries

	var health []model.HealthRecord
	s.DB.Preload("Pet").Find(&health)
	data["health_records"] = health

	var reminders []model.Reminder
	s.DB.Preload("Pet").Find(&reminders)
	data["reminders"] = reminders

	return data, nil
}

func (s *StatsService) ImportJSON(data map[string]interface{}) error {
	return s.DB.Transaction(func(tx *gorm.DB) error {
		// 清空现有数据（保持外键检查顺序）
		if err := tx.Where("1=1").Delete(&model.Reminder{}).Error; err != nil {
			return err
		}
		if err := tx.Where("1=1").Delete(&model.HealthRecord{}).Error; err != nil {
			return err
		}
		if err := tx.Where("1=1").Delete(&model.DiaryPhoto{}).Error; err != nil {
			return err
		}
		if err := tx.Where("1=1").Delete(&model.DiaryPet{}).Error; err != nil {
			return err
		}
		if err := tx.Where("1=1").Delete(&model.DiaryEntry{}).Error; err != nil {
			return err
		}
		if err := tx.Where("1=1").Delete(&model.Pet{}).Error; err != nil {
			return err
		}

		// 导入宠物
		if petsRaw, ok := data["pets"].([]interface{}); ok {
			for _, p := range petsRaw {
				petJSON, _ := json.Marshal(p)
				var pet model.Pet
				if err := json.Unmarshal(petJSON, &pet); err != nil {
					continue
				}
				// 清理关联避免外键冲突
				pet.Category = model.PetCategory{}
				pet.Breed = nil
				if err := tx.Create(&pet).Error; err != nil {
					return err
				}
			}
		}

		// 导入日记
		if diariesRaw, ok := data["diaries"].([]interface{}); ok {
			for _, d := range diariesRaw {
				diaryJSON, _ := json.Marshal(d)
				var entry model.DiaryEntry
				if err := json.Unmarshal(diaryJSON, &entry); err != nil {
					continue
				}
				// 清理关联
				entry.Pets = nil
				entry.Photos = nil
				if err := tx.Create(&entry).Error; err != nil {
					continue
				}
				// 重新解析关联数据
				var rawMap map[string]interface{}
				json.Unmarshal(diaryJSON, &rawMap)
				if petsRaw, ok := rawMap["pets"].([]interface{}); ok {
					for _, p := range petsRaw {
						petJSON, _ := json.Marshal(p)
						var dp model.DiaryPet
						json.Unmarshal(petJSON, &dp)
						dp.DiaryEntryID = entry.ID
						tx.Create(&dp)
					}
				}
				if photosRaw, ok := rawMap["photos"].([]interface{}); ok {
					for _, ph := range photosRaw {
						photoJSON, _ := json.Marshal(ph)
						var photo model.DiaryPhoto
						json.Unmarshal(photoJSON, &photo)
						photo.DiaryEntryID = entry.ID
						tx.Create(&photo)
					}
				}
			}
		}

		// 导入健康记录
		if healthRaw, ok := data["health_records"].([]interface{}); ok {
			for _, h := range healthRaw {
				healthJSON, _ := json.Marshal(h)
				var record model.HealthRecord
				if err := json.Unmarshal(healthJSON, &record); err != nil {
					continue
				}
				record.Pet = model.Pet{} // 清理关联
				if err := tx.Create(&record).Error; err != nil {
					continue
				}
			}
		}

		// 导入提醒
		if remindersRaw, ok := data["reminders"].([]interface{}); ok {
			for _, r := range remindersRaw {
				reminderJSON, _ := json.Marshal(r)
				var reminder model.Reminder
				if err := json.Unmarshal(reminderJSON, &reminder); err != nil {
					continue
				}
				reminder.Pet = nil
				reminder.HealthRecord = nil
				if err := tx.Create(&reminder).Error; err != nil {
					continue
				}
			}
		}

		return nil
	})
}
