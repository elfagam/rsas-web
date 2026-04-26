package models

import (
	"time"
	"gorm.io/gorm"
)

type Doctor struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Name       string         `gorm:"type:varchar(255);not null" json:"name"`
	Specialty  string         `gorm:"type:varchar(100);not null" json:"specialty"`
	Expertise  string         `gorm:"type:text" json:"expertise"`
	Image      string         `json:"image"`
	IsActive   bool           `gorm:"default:true" json:"is_active"`
	Schedules  []Schedule     `gorm:"foreignKey:DoctorID" json:"schedules"` // Relasi One-to-Many
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

type Schedule struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	DoctorID  uint           `json:"doctor_id"`
	Day       string         `gorm:"type:varchar(20);not null" json:"day"` // Senin, Selasa, dst
	TimeRange string         `gorm:"type:varchar(50);not null" json:"time_range"` // 08:00 - 12:00
	Quota     int            `gorm:"default:20" json:"quota"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
