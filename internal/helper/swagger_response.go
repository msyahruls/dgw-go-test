package helper

import "github.com/msyahruls/dgw-go-test/internal/dto"

type APIResponseUser struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Data    dto.UserResponse `json:"data"`
}

type APIResponseUsers struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Data    dto.UserResponses `json:"data"`
}

type APIResponseLogin struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Data    dto.LoginResponse `json:"data"`
}

type APIResponseCategory struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Data    dto.CategoryResponse `json:"data"`
}

type APIResponseCategories struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Data    dto.CategoryResponses `json:"data"`
}
