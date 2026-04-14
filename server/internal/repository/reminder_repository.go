package repository

import (
	"time"
	"pawprints-server/internal/model"
	"gorm.io/gorm"
)

type ReminderRepository struct {
	DB *gorm.DB
}

func NewReminderRepository(db *gorm.DB) *ReminderRepository {
	return &ReminderRepository{DB: db}
}

func (r *ReminderRepository) List(completed bool, page, pageSize int) ([]model.Reminder, int64, error) {
	var reminders []model.Reminder
	var total int64
	q := r.DB.Model(&model.Reminder{}).Where("completed = ?", completed).Preload("Pet")
	q.Count(&total)
	offset := (page - 1) * pageSize
	if err := q.Offset(offset).Limit(pageSize).Order("remind_at ASC").Find(&reminders).Error; err != nil {
		return nil, 0, err
	}
	return reminders, total, nil
}

func (r *ReminderRepository) GetActive() ([]model.Reminder, error) {
	var reminders []model.Reminder
	now := time.Now()
	err := r.DB.Where("completed = ? AND remind_at <= ?", false, now).
		Where("snoozed_until IS NULL OR snoozed_until <= ?", now).
		Order("remind_at ASC").Find(&reminders).Error
	return reminders, err
}

func (r *ReminderRepository) GetByID(id uint) (*model.Reminder, error) {
	var reminder model.Reminder
	if err := r.DB.First(&reminder, id).Error; err != nil {
		return nil, err
	}
	return &reminder, nil
}

func (r *ReminderRepository) Create(reminder *model.Reminder) error {
	return r.DB.Create(reminder).Error
}

func (r *ReminderRepository) Update(reminder *model.Reminder) error {
	return r.DB.Save(reminder).Error
}

func (r *ReminderRepository) Delete(id uint) error {
	return r.DB.Delete(&model.Reminder{}, id).Error
}

func (r *ReminderRepository) MarkComplete(id uint) error {
	now := time.Now()
	return r.DB.Model(&model.Reminder{}).Where("id = ?", id).Updates(map[string]interface{}{
		"completed":    true,
		"completed_at": now,
	}).Error
}

func (r *ReminderRepository) Snooze(id uint, until time.Time) error {
	return r.DB.Model(&model.Reminder{}).Where("id = ?", id).Update("snoozed_until", until).Error
}