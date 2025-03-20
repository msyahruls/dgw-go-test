package repository

import (
	"github.com/msyahruls/kreditplus-go-test/internal/domain"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(tx *domain.Transaction) error
	FindAll() ([]domain.Transaction, error)
	GetDB() *gorm.DB
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *transactionRepository) Create(tx *domain.Transaction) error {
	return r.db.Create(tx).Error
}

func (r *transactionRepository) FindAll() ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := r.db.Find(&transactions).Error
	return transactions, err
}
