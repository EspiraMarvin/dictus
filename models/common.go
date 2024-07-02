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

// populate UUID
func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if b.UUID == "" {
		b.UUID = uuid.New().String()
	}
	return
}
