package mocks

import (
	"github.com/msyahruls/kreditplus-go-test/internal/domain"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type LimitRepository struct {
	mock.Mock
}

func (m *LimitRepository) CreateOrUpdate(limit *domain.Limit) error {
	args := m.Called(limit)
	return args.Error(0)
}

func (m *LimitRepository) GetLimitsByUserID(userID uint) ([]domain.Limit, error) {
	args := m.Called(userID)
	return args.Get(0).([]domain.Limit), args.Error(1)
}

func (m *LimitRepository) GetLimitForUpdate(tx *gorm.DB, userID uint, tenor int) (*domain.Limit, error) {
	args := m.Called(tx, userID, tenor)
	return args.Get(0).(*domain.Limit), args.Error(1)
}

func (m *LimitRepository) UpdateLimit(tx *gorm.DB, limit *domain.Limit) error {
	args := m.Called(tx, limit)
	return args.Error(0)
}
