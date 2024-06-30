package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	UUID      string         `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"column:created;not null"`
	UpdatedAt time.Time      `gorm:"column:updated;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// BeforeCreate ensures a UUID and createdAt data is inserted
func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if b.UUID == "" {
		b.UUID = uuid.New().String()
	}
	return
}

type Post struct {
	Title string `gorm:"not null"`
	Body  string `gorm:"not null"`
	Base  `gorm:"embedded"`
}
