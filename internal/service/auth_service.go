package service

import (
	"errors"

	"github.com/msyahruls/dgw-go-test/internal/domain"
	"github.com/msyahruls/dgw-go-test/internal/dto"
	"github.com/msyahruls/dgw-go-test/internal/helper"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(dto.RegisterRequest) (*dto.UserResponse, error)
	Login(dto.LoginRequest) (*domain.User, error)
}

type AuthServiceImpl struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) AuthService {
	return &AuthServiceImpl{db: db}
}

func (s *AuthServiceImpl) Register(req dto.RegisterRequest) (*dto.UserResponse, error) {
	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:     req.Name,
		Username: req.Username,
		Password: hashedPassword,
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	res := &dto.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return res, nil
}

func (s *AuthServiceImpl) Login(req dto.LoginRequest) (*domain.User, error) {
	var user domain.User
	if err := s.db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return nil, errors.New("invalid username or password")
	}

	if !helper.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid username or password")
	}

	return &user, nil
}
