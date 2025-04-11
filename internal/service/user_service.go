package service

import (
	"github.com/msyahruls/dgw-go-test/internal/domain"
	"github.com/msyahruls/dgw-go-test/internal/dto"
	"github.com/msyahruls/dgw-go-test/internal/helper"
	"github.com/msyahruls/dgw-go-test/internal/repository"

	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(req dto.CreateUserRequest) (*domain.User, error)
	GetUsers() ([]domain.User, error)
	GetUserByID(id uint) (*domain.User, error)
	UpdateUser(id uint, req dto.UpdateUserRequest) (*domain.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		repo: repository.NewUserRepository(db),
	}
}

func (s *userService) CreateUser(req dto.CreateUserRequest) (*domain.User, error) {
	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:     req.Name,
		Username: req.Username,
		Password: hashedPassword,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUsers() ([]domain.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByID(id uint) (*domain.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) UpdateUser(id uint, req dto.UpdateUserRequest) (*domain.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Business logic and domain mapping
	user.Name = req.Name
	user.Username = req.Username

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
