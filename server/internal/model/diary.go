package model

import "time"

type DiaryEntry struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"size:200;not null" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	EntryDate time.Time `gorm:"not null;index" json:"entry_date"`
	Mood      string    `gorm:"size:20" json:"mood"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Pets   []DiaryPet   `gorm:"foreignKey:DiaryEntryID" json:"pets"`
	Photos []DiaryPhoto `gorm:"foreignKey:DiaryEntryID" json:"photos"`
}

type DiaryPet struct {
	ID           uint `gorm:"primaryKey" json:"id"`
	DiaryEntryID uint `gorm:"not null;index" json:"diary_entry_id"`
	PetID        uint `gorm:"not null;index" json:"pet_id"`
}

type DiaryPhoto struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	DiaryEntryID uint   `gorm:"not null;index" json:"diary_entry_id"`
	PhotoURL     string `gorm:"size:255;not null" json:"photo_url"`
	ThumbnailURL string `gorm:"size:255" json:"thumbnail_url"`
	SortOrder    int    `gorm:"default:0" json:"sort_order"`
}

func (DiaryEntry) TableName() string { return "diary_entries" }
func (DiaryPet) TableName() string   { return "diary_pets" }
func (DiaryPhoto) TableName() string { return "diary_photos" }
