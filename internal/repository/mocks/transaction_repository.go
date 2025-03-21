package mocks

import (
	"github.com/msyahruls/kreditplus-go-test/internal/domain"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	mock.Mock
}

func (m *TransactionRepository) GetDB() *gorm.DB {
	return nil // not testing actual DB
}

func (m *TransactionRepository) CreateTransaction(tx *domain.Transaction) error {
	args := m.Called(tx)
	return args.Error(0)
}

func (m *TransactionRepository) FindAllTransaction() ([]domain.Transaction, error) {
	args := m.Called()
	return args.Get(0).([]domain.Transaction), args.Error(1)
}
