package service

import (
	"pawprints-server/internal/model"
	"pawprints-server/internal/repository"
)

type DiaryService struct {
	Repo *repository.DiaryRepository
}

func NewDiaryService(repo *repository.DiaryRepository) *DiaryService {
	return &DiaryService{Repo: repo}
}

type CreateDiaryInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Mood    string `json:"mood"`
	PetIDs  []uint `json:"pet_ids"`
}

func (s *DiaryService) List(petID uint, keyword string, page, pageSize int) ([]model.DiaryEntry, int64, error) {
	return s.Repo.List(petID, keyword, page, pageSize)
}

func (s *DiaryService) GetByID(id uint) (*model.DiaryEntry, error) {
	return s.Repo.GetByID(id)
}

func (s *DiaryService) Create(input CreateDiaryInput) (*model.DiaryEntry, error) {
	entry := &model.DiaryEntry{
		Title:   input.Title,
		Content: input.Content,
		Mood:    input.Mood,
	}
	if err := s.Repo.Create(entry, input.PetIDs); err != nil {
		return nil, err
	}
	return entry, nil
}

func (s *DiaryService) Update(id uint, input CreateDiaryInput) (*model.DiaryEntry, error) {
	entry, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	entry.Title = input.Title
	entry.Content = input.Content
	entry.Mood = input.Mood
	if err := s.Repo.Update(entry, input.PetIDs); err != nil {
		return nil, err
	}
	return entry, nil
}

func (s *DiaryService) Delete(id uint) error {
	return s.Repo.Delete(id)
}

func (s *DiaryService) AddPhoto(photo *model.DiaryPhoto) error {
	return s.Repo.AddPhoto(photo)
}

func (s *DiaryService) DeletePhoto(id uint) error {
	return s.Repo.DeletePhoto(id)
}
