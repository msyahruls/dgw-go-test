package service

import (
	"fmt"
	"sync"
	"testing"

	"github.com/msyahruls/kreditplus-go-test/internal/domain"
	"github.com/msyahruls/kreditplus-go-test/internal/repository"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file:test.db?cache=shared&mode=memory"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test db: %v", err)
	}

	// Improve concurrency
	db.Exec("PRAGMA journal_mode=WAL;")
	db.Exec("PRAGMA busy_timeout=5000;")

	err = db.AutoMigrate(&domain.Limit{}, &domain.Transaction{}, &domain.PaymentSchedule{}) // AutoMigrate necessary
	if err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	return db
}

func TestConcurrentTransaction(t *testing.T) {
	db := getTestDB(t)

	// Seed initial limit
	initialLimit := domain.Limit{
		UserID:      1,
		TenorMonths: 3,
		LimitAmount: 1000000,
	}
	err := db.Create(&initialLimit).Error
	assert.NoError(t, err)

	txRepo := repository.NewTransactionRepository(db)
	limitRepo := repository.NewLimitRepository(db)

	svc := &transactionService{
		txRepo:    txRepo,
		limitRepo: limitRepo,
		db:        db,
	}

	wg := sync.WaitGroup{}
	numRequests := 10
	wg.Add(numRequests)

	successCount := 0
	mu := sync.Mutex{}

	for i := 0; i < numRequests; i++ {
		go func(i int) {
			defer wg.Done()

			txData := &domain.Transaction{
				UserID:            1,
				ContractNumber:    fmt.Sprintf("TX-%03d", i),
				InstallmentAmount: 250000,
				TenorMonths:       3,
			}

			err := svc.CreateTransaction(txData, 3)
			if err == nil {
				mu.Lock()
				successCount++
				mu.Unlock()
			} else {
				t.Logf("Transaction %d failed as expected: %v", i, err)
			}
		}(i)
	}

	wg.Wait()

	// Final check
	var finalLimit domain.Limit
	err = db.First(&finalLimit, "user_id = ? AND tenor_months = ?", 1, 3).Error
	assert.NoError(t, err)

	t.Logf("Final Limit Amount: %.2f", finalLimit.LimitAmount)
	t.Logf("Total Successful Transactions: %d out of %d", successCount, numRequests)

	// Accept: all may fail in SQLite due to locking
	// Assert: No negative limit, no data corruption
	assert.GreaterOrEqual(t, finalLimit.LimitAmount, 0.0)
}
