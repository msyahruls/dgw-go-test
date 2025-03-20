package service

import (
	"testing"

	"github.com/msyahruls/kreditplus-go-test/internal/domain"
	"github.com/msyahruls/kreditplus-go-test/internal/repository/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreateOrUpdateLimit_Success(t *testing.T) {
	// Mock repo
	mockRepo := new(mocks.LimitRepository)

	limit := domain.Limit{
		UserID:      1,
		TenorMonths: 3,
		LimitAmount: 5000000,
	}

	mockRepo.On("CreateOrUpdate", &limit).Return(nil)

	// Create service with mock repo
	svc := limitService{
		repo: mockRepo,
	}

	err := svc.CreateOrUpdateLimit(&limit)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetLimits_Success(t *testing.T) {
	mockRepo := new(mocks.LimitRepository)

	expectedLimits := []domain.Limit{
		{ID: 1, UserID: 1, TenorMonths: 3, LimitAmount: 5000000},
	}

	mockRepo.On("GetLimitsByUserID", uint(1)).Return(expectedLimits, nil)

	svc := limitService{
		repo: mockRepo,
	}

	limits, err := svc.GetLimits(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedLimits, limits)
	mockRepo.AssertExpectations(t)
}

func TestCreateOrUpdateLimit_Failure(t *testing.T) {
	mockRepo := new(mocks.LimitRepository)

	limit := &domain.Limit{UserID: 1}
	mockRepo.On("CreateOrUpdate", limit).Return(assert.AnError)

	svc := &limitService{repo: mockRepo}

	err := svc.CreateOrUpdateLimit(limit)

	assert.Error(t, err)
	assert.Equal(t, assert.AnError, err)
	mockRepo.AssertExpectations(t)
}

func TestGetLimits_Failure(t *testing.T) {
	mockRepo := new(mocks.LimitRepository)

	mockRepo.On("GetLimitsByUserID", uint(1)).Return([]domain.Limit{}, assert.AnError)

	svc := &limitService{repo: mockRepo}

	limits, err := svc.GetLimits(1)

	assert.Error(t, err)
	assert.Empty(t, limits)
	assert.Equal(t, assert.AnError, err)
	mockRepo.AssertExpectations(t)
}
