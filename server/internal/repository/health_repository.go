package repository

import (
	"time"
	"pawprints-server/internal/model"
	"gorm.io/gorm"
)

type HealthRepository struct {
	DB *gorm.DB
}

func NewHealthRepository(db *gorm.DB) *HealthRepository {
	return &HealthRepository{DB: db}
}

func (r *HealthRepository) List(petID uint, recordType string, page, pageSize int) ([]model.HealthRecord, int64, error) {
	var records []model.HealthRecord
	var total int64
	q := r.DB.Model(&model.HealthRecord{})
	if petID > 0 {
		q = q.Where("pet_id = ?", petID)
	}
	if recordType != "" {
		q = q.Where("type = ?", recordType)
	}
	q.Count(&total)
	offset := (page - 1) * pageSize
	if err := q.Offset(offset).Limit(pageSize).Order("record_date DESC").Find(&records).Error; err != nil {
		return nil, 0, err
	}
	return records, total, nil
}

func (r *HealthRepository) GetByID(id uint) (*model.HealthRecord, error) {
	var record model.HealthRecord
	if err := r.DB.First(&record, id).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *HealthRepository) Create(record *model.HealthRecord) error {
	return r.DB.Create(record).Error
}

func (r *HealthRepository) Update(record *model.HealthRecord) error {
	return r.DB.Save(record).Error
}

func (r *HealthRepository) Delete(id uint) error {
	return r.DB.Delete(&model.HealthRecord{}, id).Error
}

func (r *HealthRepository) GetWeightChart(petID uint) ([]model.HealthRecord, error) {
	var records []model.HealthRecord
	err := r.DB.Where("pet_id = ? AND type = 'weight'", petID).Order("record_date ASC").Find(&records).Error
	return records, err
}

func (r *HealthRepository) GetUpcoming(petID uint, withinDays int) ([]model.HealthRecord, error) {
	var records []model.HealthRecord
	now := time.Now()
	endDate := now.AddDate(0, 0, withinDays)
	err := r.DB.Where("next_due_date >= ? AND next_due_date <= ?", now, endDate).
		Where("pet_id = ? OR pet_id = 0", petID).
		Order("next_due_date ASC").Find(&records).Error
	return records, err
}