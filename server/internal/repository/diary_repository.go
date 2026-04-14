package repository

import (
	"pawprints-server/internal/model"
	"gorm.io/gorm"
)

type DiaryRepository struct {
	DB *gorm.DB
}

func NewDiaryRepository(db *gorm.DB) *DiaryRepository {
	return &DiaryRepository{DB: db}
}

func (r *DiaryRepository) List(petID uint, keyword string, page, pageSize int) ([]model.DiaryEntry, int64, error) {
	var entries []model.DiaryEntry
	var total int64
	q := r.DB.Model(&model.DiaryEntry{}).Preload("Photos").Preload("Pets")
	if petID > 0 {
		q = q.Joins("JOIN diary_pets ON diary_pets.diary_entry_id = diary_entries.id").Where("diary_pets.pet_id = ?", petID)
	}
	if keyword != "" {
		q = q.Where("title LIKE ? OR content LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	q.Count(&total)
	offset := (page - 1) * pageSize
	if err := q.Offset(offset).Limit(pageSize).Order("entry_date DESC").Find(&entries).Error; err != nil {
		return nil, 0, err
	}
	r.populatePetNames(entries)
	return entries, total, nil
}

func (r *DiaryRepository) GetByID(id uint) (*model.DiaryEntry, error) {
	var entry model.DiaryEntry
	if err := r.DB.Preload("Photos").Preload("Pets").First(&entry, id).Error; err != nil {
		return nil, err
	}
	r.populatePetNames([]model.DiaryEntry{entry})
	return &entry, nil
}

func (r *DiaryRepository) populatePetNames(entries []model.DiaryEntry) {
	petNames := make(map[uint]string)
	for i := range entries {
		for j := range entries[i].Pets {
			pid := entries[i].Pets[j].PetID
			if name, ok := petNames[pid]; ok {
				entries[i].Pets[j].PetName = name
			} else {
				var name string
				r.DB.Model(&model.Pet{}).Where("id = ?", pid).Select("name").Scan(&name)
				petNames[pid] = name
				entries[i].Pets[j].PetName = name
			}
		}
	}
}

func (r *DiaryRepository) Create(entry *model.DiaryEntry, petIDs []uint, photos []model.DiaryPhoto) error {
	tx := r.DB.Begin()
	if err := tx.Create(entry).Error; err != nil {
		tx.Rollback()
		return err
	}
	for _, petID := range petIDs {
		if err := tx.Create(&model.DiaryPet{DiaryEntryID: entry.ID, PetID: petID}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	for i := range photos {
		photos[i].DiaryEntryID = entry.ID
		if err := tx.Create(&photos[i]).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (r *DiaryRepository) Update(entry *model.DiaryEntry, petIDs []uint, photos []model.DiaryPhoto) error {
	tx := r.DB.Begin()
	if err := tx.Save(entry).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Where("diary_entry_id = ?", entry.ID).Delete(&model.DiaryPet{})
	for _, petID := range petIDs {
		if err := tx.Create(&model.DiaryPet{DiaryEntryID: entry.ID, PetID: petID}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Where("diary_entry_id = ?", entry.ID).Delete(&model.DiaryPhoto{})
	for i := range photos {
		photos[i].DiaryEntryID = entry.ID
		if err := tx.Create(&photos[i]).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (r *DiaryRepository) Delete(id uint) error {
	tx := r.DB.Begin()
	tx.Where("diary_entry_id = ?", id).Delete(&model.DiaryPet{})
	tx.Where("diary_entry_id = ?", id).Delete(&model.DiaryPhoto{})
	if err := tx.Delete(&model.DiaryEntry{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *DiaryRepository) AddPhoto(photo *model.DiaryPhoto) error {
	return r.DB.Create(photo).Error
}

func (r *DiaryRepository) DeletePhoto(id uint) error {
	return r.DB.Delete(&model.DiaryPhoto{}, id).Error
}
