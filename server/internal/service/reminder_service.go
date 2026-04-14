package service

import (
	"time"
	"pawprints-server/internal/model"
	"pawprints-server/internal/repository"
)

type ReminderService struct {
	Repo *repository.ReminderRepository
}

func NewReminderService(repo *repository.ReminderRepository) *ReminderService {
	return &ReminderService{Repo: repo}
}

func (s *ReminderService) List(completed bool, page, pageSize int) ([]model.Reminder, int64, error) {
	return s.Repo.List(completed, page, pageSize)
}

func (s *ReminderService) GetActive() ([]model.Reminder, error) {
	return s.Repo.GetActive()
}

func (s *ReminderService) Create(reminder *model.Reminder) error {
	return s.Repo.Create(reminder)
}

func (s *ReminderService) Update(reminder *model.Reminder) error {
	return s.Repo.Update(reminder)
}

func (s *ReminderService) Delete(id uint) error {
	return s.Repo.Delete(id)
}

func (s *ReminderService) Complete(id uint) error {
	return s.Repo.MarkComplete(id)
}

func (s *ReminderService) Snooze(id uint, duration time.Duration) error {
	return s.Repo.Snooze(id, time.Now().Add(duration))
}