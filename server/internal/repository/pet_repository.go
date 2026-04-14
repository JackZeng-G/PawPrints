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
	var photos []model.DiaryPhoto
	err := r.DB.Raw(`
		SELECT dp.* FROM diary_photos dp
		JOIN diary_pets dpet ON dp.diary_entry_id = dpet.diary_entry_id
		WHERE dpet.pet_id = ?
		ORDER BY dp.sort_order`, petID).Scan(&photos).Error
	return photos, err
}
