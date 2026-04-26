package models

import (
	"time"
	"gorm.io/gorm"
)

type News struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"type:varchar(255);not null" json:"title"`
	Slug      string         `gorm:"uniqueIndex;not null" json:"slug"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	Thumbnail string         `json:"thumbnail"`
	Category  string         `gorm:"type:varchar(50)" json:"category"`
	AuthorID  uint           `json:"author_id"`
	Author    User           `gorm:"foreignKey:AuthorID" json:"author"` // Relasi ke User
	IsDraft   bool           `gorm:"default:true" json:"is_draft"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
