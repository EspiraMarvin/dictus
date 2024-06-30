package models

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	UUID      string         `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"column:created;not null"`
	UpdatedAt time.Time      `gorm:"column:updated;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Post struct {
	// gorm.Model
	Title string
	Body  string
	Base
}
