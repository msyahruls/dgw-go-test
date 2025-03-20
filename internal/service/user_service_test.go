package service

import (
	"testing"

	"github.com/msyahruls/kreditplus-go-test/internal/domain"
	"github.com/msyahruls/kreditplus-go-test/internal/repository/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	expectedUser := &domain.User{
		NIK:         "123456",
		FullName:    "John Doe",
		LegalName:   "John Doe",
		BirthPlace:  "Jakarta",
		Salary:      5000000,
		PhotoIDCard: "ktp.jpg",
		PhotoSelfie: "selfie.jpg",
	}

	mockRepo.On("Create", expectedUser).Return(nil)

	svc := &userService{repo: mockRepo}

	err := svc.CreateUser(expectedUser)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetUsers_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	expectedUsers := []domain.User{
		{ID: 1, FullName: "John Doe"},
		{ID: 2, FullName: "Jane Doe"},
	}

	mockRepo.On("FindAll").Return(expectedUsers, nil)

	svc := &userService{repo: mockRepo}

	users, err := svc.GetUsers()

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_Failure(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	user := &domain.User{NIK: "123456"}
	mockRepo.On("Create", user).Return(assert.AnError)

	svc := &userService{repo: mockRepo}

	err := svc.CreateUser(user)

	assert.Error(t, err)
	assert.Equal(t, assert.AnError, err)
	mockRepo.AssertExpectations(t)
}

func TestGetUsers_Failure(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	mockRepo.On("FindAll").Return([]domain.User{}, assert.AnError)

	svc := &userService{repo: mockRepo}

	users, err := svc.GetUsers()

	assert.Error(t, err)
	assert.Empty(t, users)
	assert.Equal(t, assert.AnError, err)
	mockRepo.AssertExpectations(t)
}
