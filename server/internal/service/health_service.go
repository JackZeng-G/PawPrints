package service

import (
	"pawprints-server/internal/model"
	"pawprints-server/internal/repository"
)

type HealthService struct {
	Repo *repository.HealthRepository
}

func NewHealthService(repo *repository.HealthRepository) *HealthService {
	return &HealthService{Repo: repo}
}

func (s *HealthService) List(petID uint, recordType string, page, pageSize int) ([]model.HealthRecord, int64, error) {
	return s.Repo.List(petID, recordType, page, pageSize)
}

func (s *HealthService) GetByID(id uint) (*model.HealthRecord, error) {
	return s.Repo.GetByID(id)
}

func (s *HealthService) Create(record *model.HealthRecord) error {
	return s.Repo.Create(record)
}

func (s *HealthService) Update(record *model.HealthRecord) error {
	return s.Repo.Update(record)
}

func (s *HealthService) Delete(id uint) error {
	return s.Repo.Delete(id)
}

func (s *HealthService) GetWeightChart(petID uint) ([]model.HealthRecord, error) {
	return s.Repo.GetWeightChart(petID)
}

func (s *HealthService) GetUpcoming(petID uint, withinDays int) ([]model.HealthRecord, error) {
	return s.Repo.GetUpcoming(petID, withinDays)
}