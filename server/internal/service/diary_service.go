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

type PhotoInput struct {
	PhotoURL     string `json:"photo_url"`
	ThumbnailURL string `json:"thumbnail_url"`
}

type CreateDiaryInput struct {
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	Mood      string          `json:"mood"`
	EntryDate model.DateOnly  `json:"entry_date"`
	PetIDs    []uint          `json:"pet_ids"`
	Photos    []PhotoInput    `json:"photos"`
}

func (s *DiaryService) List(petID uint, keyword string, page, pageSize int) ([]model.DiaryEntry, int64, error) {
	return s.Repo.List(petID, keyword, page, pageSize)
}

func (s *DiaryService) GetByID(id uint) (*model.DiaryEntry, error) {
	return s.Repo.GetByID(id)
}

func (s *DiaryService) Create(input CreateDiaryInput) (*model.DiaryEntry, error) {
	entry := &model.DiaryEntry{
		Title:     input.Title,
		Content:   input.Content,
		Mood:      input.Mood,
		EntryDate: input.EntryDate,
	}
	photos := make([]model.DiaryPhoto, 0, len(input.Photos))
	for i, p := range input.Photos {
		photos = append(photos, model.DiaryPhoto{
			PhotoURL:     p.PhotoURL,
			ThumbnailURL: p.ThumbnailURL,
			SortOrder:    i,
		})
	}
	if err := s.Repo.Create(entry, input.PetIDs, photos); err != nil {
		return nil, err
	}
	// 重新加载以包含关联数据
	return s.Repo.GetByID(entry.ID)
}

func (s *DiaryService) Update(id uint, input CreateDiaryInput) (*model.DiaryEntry, error) {
	entry, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	entry.Title = input.Title
	entry.Content = input.Content
	entry.Mood = input.Mood
	entry.EntryDate = input.EntryDate
	photos := make([]model.DiaryPhoto, 0, len(input.Photos))
	for i, p := range input.Photos {
		photos = append(photos, model.DiaryPhoto{
			PhotoURL:     p.PhotoURL,
			ThumbnailURL: p.ThumbnailURL,
			SortOrder:    i,
		})
	}
	if err := s.Repo.Update(entry, input.PetIDs, photos); err != nil {
		return nil, err
	}
	// 重新加载以包含关联数据
	return s.Repo.GetByID(entry.ID)
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
