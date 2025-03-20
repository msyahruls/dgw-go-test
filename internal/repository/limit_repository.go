package repository

import (
	"github.com/msyahruls/kreditplus-go-test/internal/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LimitRepository interface {
	CreateOrUpdate(limit *domain.Limit) error
	GetLimitsByUserID(userID uint) ([]domain.Limit, error)
	GetLimitForUpdate(tx *gorm.DB, userID uint, tenor int) (*domain.Limit, error)
	UpdateLimit(tx *gorm.DB, limit *domain.Limit) error
}

type limitRepository struct {
	db *gorm.DB
}

func NewLimitRepository(db *gorm.DB) LimitRepository {
	return &limitRepository{db: db}
}

func (r *limitRepository) CreateOrUpdate(limit *domain.Limit) error {
	var existing domain.Limit
	err := r.db.Where("user_id = ? AND tenor_months = ?", limit.UserID, limit.TenorMonths).First(&existing).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if existing.ID != 0 {
		// Update existing limit
		existing.LimitAmount = limit.LimitAmount
		return r.db.Save(&existing).Error
	}

	// Create new limit
	return r.db.Create(limit).Error
}

func (r *limitRepository) GetLimitsByUserID(userID uint) ([]domain.Limit, error) {
	var limits []domain.Limit
	err := r.db.Where("user_id = ?", userID).Find(&limits).Error
	return limits, err
}

func (r *limitRepository) GetLimitForUpdate(tx *gorm.DB, userID uint, tenor int) (*domain.Limit, error) {
	var limit domain.Limit
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("user_id = ? AND tenor_months = ?", userID, tenor).
		First(&limit).Error
	if err != nil {
		return nil, err
	}
	return &limit, nil
}

func (r *limitRepository) UpdateLimit(tx *gorm.DB, limit *domain.Limit) error {
	return tx.Save(limit).Error
}
