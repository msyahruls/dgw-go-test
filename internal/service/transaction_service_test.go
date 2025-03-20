package service

import (
	"testing"

	"github.com/msyahruls/kreditplus-go-test/internal/domain"
	"github.com/msyahruls/kreditplus-go-test/internal/repository/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateTransaction_Success(t *testing.T) {
	// Setup Mock Repos
	mockTxRepo := new(mocks.TransactionRepository)
	mockLimitRepo := new(mocks.LimitRepository)

	// In-memory DB
	testDB := getTestDB(t)

	// Transaction data
	txData := &domain.Transaction{
		UserID:            1,
		ContractNumber:    "TX-001",
		InstallmentAmount: 2000000,
	}

	// Limit data
	limit := &domain.Limit{
		ID:          1,
		UserID:      1,
		TenorMonths: 3,
		LimitAmount: 5000000,
	}

	// Mock behaviors
	mockLimitRepo.On("GetLimitForUpdate", mock.Anything, uint(1), 3).
		Return(limit, nil)

	mockLimitRepo.On("UpdateLimit", mock.Anything, mock.Anything).
		Return(nil)

	mockTxRepo.On("Create", txData).Return(nil)

	// Service with real DB and mocks
	svc := &transactionService{
		txRepo:    mockTxRepo,
		limitRepo: mockLimitRepo,
		db:        testDB,
	}

	// Execute
	err := svc.CreateTransaction(txData, 3)

	// Assert
	assert.NoError(t, err)
	mockLimitRepo.AssertExpectations(t)
	mockTxRepo.AssertExpectations(t)
}

func TestCreateTransaction_InsufficientLimit(t *testing.T) {
	mockTxRepo := new(mocks.TransactionRepository)
	mockLimitRepo := new(mocks.LimitRepository)

	limit := &domain.Limit{
		ID:          1,
		UserID:      1,
		TenorMonths: 3,
		LimitAmount: 1000000, // less than transaction amount
	}

	txData := &domain.Transaction{
		UserID:            1,
		ContractNumber:    "TX-002",
		InstallmentAmount: 2000000,
	}

	// Mock LimitRepo behavior
	mockLimitRepo.On("GetLimitForUpdate", mock.Anything, uint(1), 3).Return(limit, nil)

	// In-memory DB
	testDB := getTestDB(t)

	svc := &transactionService{
		txRepo:    mockTxRepo,
		limitRepo: mockLimitRepo,
		db:        testDB,
	}

	err := svc.CreateTransaction(txData, 3)

	assert.Error(t, err)
	assert.Equal(t, "insufficient limit", err.Error())

	mockLimitRepo.AssertExpectations(t)
	mockTxRepo.AssertNotCalled(t, "Create", mock.Anything)
}

func TestCreateTransaction_LimitNotFound(t *testing.T) {
	mockTxRepo := new(mocks.TransactionRepository)
	mockLimitRepo := new(mocks.LimitRepository)

	txData := &domain.Transaction{UserID: 1, InstallmentAmount: 2000000}

	mockLimitRepo.On("GetLimitForUpdate", mock.Anything, uint(1), 3).Return((*domain.Limit)(nil), assert.AnError)

	// In-memory DB
	testDB := getTestDB(t)

	svc := &transactionService{
		txRepo:    mockTxRepo,
		limitRepo: mockLimitRepo,
		db:        testDB,
	}

	err := svc.CreateTransaction(txData, 3)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "limit not found")
	mockLimitRepo.AssertExpectations(t)
	mockTxRepo.AssertNotCalled(t, "Create", mock.Anything)
}

func TestCreateTransaction_UpdateLimitError(t *testing.T) {
	mockTxRepo := new(mocks.TransactionRepository)
	mockLimitRepo := new(mocks.LimitRepository)

	limit := &domain.Limit{ID: 1, UserID: 1, LimitAmount: 5000000}

	txData := &domain.Transaction{UserID: 1, InstallmentAmount: 2000000}

	mockLimitRepo.On("GetLimitForUpdate", mock.Anything, uint(1), 3).Return(limit, nil)
	mockLimitRepo.On("UpdateLimit", mock.Anything, mock.Anything).Return(assert.AnError)

	// In-memory DB
	testDB := getTestDB(t)

	svc := &transactionService{
		txRepo:    mockTxRepo,
		limitRepo: mockLimitRepo,
		db:        testDB,
	}

	err := svc.CreateTransaction(txData, 3)

	assert.Error(t, err)
	assert.Equal(t, assert.AnError, err)
	mockLimitRepo.AssertExpectations(t)
	mockTxRepo.AssertNotCalled(t, "Create", mock.Anything)
}

func TestCreateTransaction_SaveTransactionError(t *testing.T) {
	mockTxRepo := new(mocks.TransactionRepository)
	mockLimitRepo := new(mocks.LimitRepository)

	limit := &domain.Limit{ID: 1, UserID: 1, LimitAmount: 5000000}

	txData := &domain.Transaction{UserID: 1, InstallmentAmount: 2000000}

	mockLimitRepo.On("GetLimitForUpdate", mock.Anything, uint(1), 3).Return(limit, nil)
	mockLimitRepo.On("UpdateLimit", mock.Anything, mock.Anything).Return(nil)
	mockTxRepo.On("Create", txData).Return(assert.AnError)

	// In-memory DB
	testDB := getTestDB(t)

	svc := &transactionService{
		txRepo:    mockTxRepo,
		limitRepo: mockLimitRepo,
		db:        testDB,
	}

	err := svc.CreateTransaction(txData, 3)

	assert.Error(t, err)
	assert.Equal(t, assert.AnError, err)
	mockLimitRepo.AssertExpectations(t)
	mockTxRepo.AssertExpectations(t)
}

// Helper to get in-memory SQLite DB for tests
func getTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test DB: %v", err)
	}
	return db
}
