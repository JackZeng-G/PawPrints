package repository

import (
	"pawprints-server/internal/model"

	"gorm.io/gorm"
)

type PetRepository struct {
	DB *gorm.DB
}

func NewPetRepository(db *gorm.DB) *PetRepository {
	return &PetRepository{DB: db}
}

func (r *PetRepository) List(status string, page, pageSize int) ([]model.Pet, int64, error) {
	var pets []model.Pet
	var total int64
	q := r.DB.Model(&model.Pet{}).Preload("Category").Preload("Breed")
	if status != "" {
		q = q.Where("status = ?", status)
	}
	q.Count(&total)
	offset := (page - 1) * pageSize
	if err := q.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&pets).Error; err != nil {
		return nil, 0, err
	}
	return pets, total, nil
}

func (r *PetRepository) GetByID(id uint) (*model.Pet, error) {
	var pet model.Pet
	if err := r.DB.Preload("Category").Preload("Breed").First(&pet, id).Error; err != nil {
		return nil, err
	}
	return &pet, nil
}

func (r *PetRepository) Create(pet *model.Pet) error {
	return r.DB.Create(pet).Error
}

func (r *PetRepository) Update(pet *model.Pet) error {
	return r.DB.Save(pet).Error
}

func (r *PetRepository) Delete(id uint) error {
	return r.DB.Delete(&model.Pet{}, id).Error
}

func (r *PetRepository) GetPhotos(petID uint) ([]model.DiaryPhoto, error) {
	type photoRow struct {
		ID           uint   `gorm:"primaryKey"`
		DiaryEntryID uint   `gorm:"column:diary_entry_id"`
		PhotoURL     string `gorm:"column:photo_url"`
		ThumbnailURL string `gorm:"column:thumbnail_url"`
		SortOrder    int    `gorm:"column:sort_order"`
		EntryDate    string `gorm:"column:entry_date"`
	}
	var rows []photoRow
	err := r.DB.Raw(`
		SELECT dp.id, dp.diary_entry_id, dp.photo_url, dp.thumbnail_url, dp.sort_order,
		       de.entry_date
		FROM diary_photos dp
		JOIN diary_pets dpet ON dp.diary_entry_id = dpet.diary_entry_id
		JOIN diary_entries de ON dp.diary_entry_id = de.id
		WHERE dpet.pet_id = ?
		ORDER BY dp.sort_order`, petID).Scan(&rows).Error
	if err != nil {
		return nil, err
	}
	photos := make([]model.DiaryPhoto, len(rows))
	for i, row := range rows {
		photos[i] = model.DiaryPhoto{
			ID:           row.ID,
			DiaryEntryID: row.DiaryEntryID,
			PhotoURL:     row.PhotoURL,
			ThumbnailURL: row.ThumbnailURL,
			SortOrder:    row.SortOrder,
			EntryDate:    row.EntryDate,
		}
	}
	return photos, nil
}
