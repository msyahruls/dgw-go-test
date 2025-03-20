package service

import (
	"github.com/msyahruls/kreditplus-go-test/internal/domain"
	"github.com/msyahruls/kreditplus-go-test/internal/repository"

	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(user *domain.User) error
	GetUsers() ([]domain.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		repo: repository.NewUserRepository(db),
	}
}

func (s *userService) CreateUser(user *domain.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetUsers() ([]domain.User, error) {
	return s.repo.FindAll()
}
