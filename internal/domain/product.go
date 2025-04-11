package domain

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID         uint           `gorm:"primaryKey"`
	Name       string         `gorm:"not null"`
	Price      float64        `gorm:"not null"`
	CategoryID uint           `json:"category_id"`
	Category   Category       `gorm:"foreignKey:CategoryID" json:"category"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
