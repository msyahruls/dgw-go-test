package service

import (
	"fmt"
	"time"

	"github.com/msyahruls/kreditplus-go-test/internal/domain"
	"github.com/msyahruls/kreditplus-go-test/internal/repository"

	"gorm.io/gorm"
)

type TransactionService interface {
	CreateTransaction(txData *domain.Transaction, tenor int) error
	GetTransactions() ([]domain.Transaction, error)
	GetPaymentSchedules(id uint) ([]domain.PaymentSchedule, error)
	PayInstallment(id uint, paymentDate time.Time) error
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

		txData.TenorMonths = tenor

		// Save transaction to get the TransactionID
		if err := s.txRepo.CreateTransaction(txData); err != nil {
			return err
		}

		// Create payment schedules
		for i := 1; i <= tenor; i++ {
			schedule := domain.PaymentSchedule{
				TransactionID: txData.ID,
				DueDate:       time.Now().AddDate(0, i, 0), // Add i months
				Amount:        txData.InstallmentAmount,
				Status:        "UNPAID",
			}
			if err := tx.Create(&schedule).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *transactionService) GetTransactions() ([]domain.Transaction, error) {
	return s.txRepo.FindAllTransaction()
}

func (s *transactionService) GetPaymentSchedules(txID uint) ([]domain.PaymentSchedule, error) {
	var schedules []domain.PaymentSchedule
	err := s.db.Where("transaction_id = ?", txID).Find(&schedules).Error
	return schedules, err
}

func (s *transactionService) PayInstallment(scheduleID uint, paymentDate time.Time) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var schedule domain.PaymentSchedule
		err := tx.First(&schedule, scheduleID).Error
		if err != nil {
			return err
		}

		if schedule.Status == "PAID" {
			return fmt.Errorf("installment already paid")
		}

		schedule.Status = "PAID"
		schedule.PaymentDate = &paymentDate
		if err := tx.Save(&schedule).Error; err != nil {
			return err
		}

		// check if all schedule is paid
		var unpaidCount int64
		err = tx.Model(&domain.PaymentSchedule{}).
			Where("transaction_id = ? AND status = ?", schedule.TransactionID, "UNPAID").
			Count(&unpaidCount).Error
		if err != nil {
			return err
		}

		// if paid → update limit
		if unpaidCount == 0 {
			var transaction domain.Transaction
			if err := tx.First(&transaction, schedule.TransactionID).Error; err != nil {
				return err
			}

			limit, err := s.limitRepo.GetLimitForUpdate(tx, transaction.UserID, transaction.TenorMonths)
			if err != nil {
				return err
			}

			// Tambahkan kembali total transaksi
			total := transaction.InstallmentAmount
			limit.LimitAmount += total
			if err := s.limitRepo.UpdateLimit(tx, limit); err != nil {
				return err
			}
		}

		return nil
		// return tx.Save(&schedule).Error
	})
}
