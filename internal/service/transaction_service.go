package service

import (
	"fmt"

	"github.com/msyahruls/kreditplus-go-test/internal/domain"
	"github.com/msyahruls/kreditplus-go-test/internal/repository"

	"gorm.io/gorm"
)

type TransactionService interface {
	CreateTransaction(txData *domain.Transaction, tenor int) error
	GetTransactions() ([]domain.Transaction, error)
}

type transactionService struct {
	txRepo    repository.TransactionRepository
	limitRepo repository.LimitRepository
	db        *gorm.DB
}

func NewTransactionService(db *gorm.DB, txRepo repository.TransactionRepository, limitRepo repository.LimitRepository) TransactionService {
	return &transactionService{
		db:        db,
		txRepo:    txRepo,
		limitRepo: limitRepo,
	}
}

func (s *transactionService) CreateTransaction(txData *domain.Transaction, tenor int) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// Use injected limitRepo
		limit, err := s.limitRepo.GetLimitForUpdate(tx, txData.UserID, tenor)
		if err != nil {
			return fmt.Errorf("limit not found: %w", err)
		}

		// Check if enough limit
		if limit.LimitAmount < txData.InstallmentAmount {
			return fmt.Errorf("insufficient limit")
		}

		// Deduct limit
		limit.LimitAmount -= txData.InstallmentAmount
		err = s.limitRepo.UpdateLimit(tx, limit)
		if err != nil {
			return err
		}

		// Save transaction
		return s.txRepo.Create(txData)
	})
}

func (s *transactionService) GetTransactions() ([]domain.Transaction, error) {
	return s.txRepo.FindAll()
}
