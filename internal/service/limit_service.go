package service

import (
	"github.com/msyahruls/kreditplus-go-test/internal/domain"
	"github.com/msyahruls/kreditplus-go-test/internal/repository"

	"gorm.io/gorm"
)

type LimitService interface {
	CreateOrUpdateLimit(limit *domain.Limit) error
	GetLimits(userID uint) ([]domain.Limit, error)
}

type limitService struct {
	repo repository.LimitRepository
}

func NewLimitService(db *gorm.DB) LimitService {
	return &limitService{
		repo: repository.NewLimitRepository(db),
	}
}

func (s *limitService) CreateOrUpdateLimit(limit *domain.Limit) error {
	return s.repo.CreateOrUpdate(limit)
}

func (s *limitService) GetLimits(userID uint) ([]domain.Limit, error) {
	return s.repo.GetLimitsByUserID(userID)
}
