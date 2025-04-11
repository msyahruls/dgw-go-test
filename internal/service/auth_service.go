package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/msyahruls/dgw-go-test/internal/config"
	"github.com/msyahruls/dgw-go-test/internal/domain"
	"github.com/msyahruls/dgw-go-test/internal/dto"
	"github.com/msyahruls/dgw-go-test/internal/helper"
	"github.com/msyahruls/dgw-go-test/internal/repository"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(dto.RegisterRequest) (*domain.User, error)
	Login(dto.LoginRequest) (*domain.User, string, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(db *gorm.DB) AuthService {
	return &authService{
		repo: repository.NewUserRepository(db),
	}
}

func (s *authService) Register(req dto.RegisterRequest) (*domain.User, error) {
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

func (s *authService) Login(req dto.LoginRequest) (*domain.User, string, error) {
	user, err := s.repo.FindByUsername(req.Username)
	if err != nil {
		return nil, "", err
	}

	if !helper.CheckPasswordHash(req.Password, user.Password) {
		return nil, "", errors.New("invalid username or password")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"name":     user.Name,
		"exp":      time.Now().Add(72 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		return nil, "", err
	}

	return user, tokenString, nil
}
