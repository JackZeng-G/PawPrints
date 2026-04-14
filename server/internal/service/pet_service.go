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
	if err := s.Repo.Create(pet); err != nil {
		return err
	}
	// 重新加载以包含关联数据
	loaded, err := s.Repo.GetByID(pet.ID)
	if err != nil {
		return nil // 创建已成功，仅加载失败
	}
	*pet = *loaded
	return nil
}

func (s *PetService) Update(pet *model.Pet) error {
	if err := s.Repo.Update(pet); err != nil {
		return err
	}
	// 重新加载以包含关联数据
	loaded, err := s.Repo.GetByID(pet.ID)
	if err != nil {
		return nil // 更新已成功，仅加载失败
	}
	*pet = *loaded
	return nil
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
