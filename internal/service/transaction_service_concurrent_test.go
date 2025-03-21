package service

import (
	"log"
	"strconv"
	"sync"
	"testing"

	"github.com/msyahruls/kreditplus-go-test/internal/domain"
	"github.com/msyahruls/kreditplus-go-test/internal/repository"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test db: %v", err)
	}
	err = db.AutoMigrate(&domain.Limit{}, &domain.Transaction{}) // AutoMigrate necessary
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
		LimitAmount: 1000000, // 1,000,000
	}
	err := db.Create(&initialLimit).Error
	assert.NoError(t, err)

	// Repositories
	txRepo := repository.NewTransactionRepository(db)
	limitRepo := repository.NewLimitRepository(db)

	// Service
	svc := &transactionService{
		txRepo:    txRepo,
		limitRepo: limitRepo,
		db:        db,
	}

	wg := sync.WaitGroup{}
	numRequests := 10
	wg.Add(numRequests)

	successCount := 0
	mu := sync.Mutex{} // Protect successCount

	for i := 0; i < numRequests; i++ {
		go func(i int) {
			defer wg.Done()

			txData := &domain.Transaction{
				UserID:            1,
				ContractNumber:    "TX-00" + strconv.Itoa(i), // simple unique contract number
				InstallmentAmount: 250000,                    // Each transaction deducts 250,000
			}

			err := svc.CreateTransaction(txData, 3)
			if err == nil {
				mu.Lock()
				successCount++
				mu.Unlock()
			} else {
				log.Printf("Transaction %d failed: %v", i, err)
			}
		}(i)
	}

	wg.Wait()

	// Check final limit & number of successful transactions
	var finalLimit domain.Limit
	_ = db.First(&finalLimit, "user_id = ? AND tenor_months = ?", 1, 3)

	t.Logf("Final Limit Amount: %v", finalLimit.LimitAmount)
	t.Logf("Total Successful Transactions: %d", successCount)

	// Expected: only 4 transactions of 250k can succeed (total limit = 1M)
	assert.Equal(t, float64(0), finalLimit.LimitAmount)
	assert.Equal(t, 4, successCount)
}
