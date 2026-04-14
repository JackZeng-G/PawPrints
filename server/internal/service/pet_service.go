package service

import (
	"pawprints-server/internal/model"
	"pawprints-server/internal/repository"
)

type PetService struct {
	Repo *repository.PetRepository
}

func NewPetService(repo *repository.PetRepository) *PetService {
	return &PetService{Repo: repo}
}

func (s *PetService) List(status string, page, pageSize int) ([]model.Pet, int64, error) {
	return s.Repo.List(status, page, pageSize)
}

func (s *PetService) GetByID(id uint) (*model.Pet, error) {
	return s.Repo.GetByID(id)
}

func (s *PetService) Create(pet *model.Pet) error {
	return s.Repo.Create(pet)
}

func (s *PetService) Update(pet *model.Pet) error {
	return s.Repo.Update(pet)
}

func (s *PetService) Delete(id uint) error {
	return s.Repo.Delete(id)
}

func (s *PetService) SetMemorial(id uint) error {
	pet, err := s.Repo.GetByID(id)
	if err != nil {
		return err
	}
	pet.Status = "memorial"
	return s.Repo.Update(pet)
}

func (s *PetService) GetPhotos(petID uint) ([]model.DiaryPhoto, error) {
	return s.Repo.GetPhotos(petID)
}
