package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	NIK         string         `gorm:"unique;not null" json:"nik"`
	FullName    string         `json:"full_name"`
	LegalName   string         `json:"legal_name"`
	BirthPlace  string         `json:"birth_place"`
	BirthDate   time.Time      `gorm:"type:date" json:"birth_date"`
	Salary      float64        `json:"salary"`
	PhotoIDCard string         `json:"photo_id_card"`
	PhotoSelfie string         `json:"photo_selfie"`
	Limits      []Limit        `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"limits"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type Limit struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserID      uint           `gorm:"not null" json:"user_id"`
	TenorMonths int            `json:"tenor_months"`
	LimitAmount float64        `json:"limit_amount"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
