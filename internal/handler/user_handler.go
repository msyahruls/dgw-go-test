package handler

import (
	"net/http"

	"github.com/msyahruls/kreditplus-go-test/internal/domain"
	"github.com/msyahruls/kreditplus-go-test/internal/dto"
	"github.com/msyahruls/kreditplus-go-test/internal/helper"
	"github.com/msyahruls/kreditplus-go-test/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		UserService: service.NewUserService(db),
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	user := domain.User{
		NIK:         req.NIK,
		FullName:    req.FullName,
		LegalName:   req.LegalName,
		BirthPlace:  req.BirthPlace,
		BirthDate:   parseDate(req.BirthDate), // helper func (parse time.Time)
		Salary:      req.Salary,
		PhotoIDCard: req.PhotoIDCard,
		PhotoSelfie: req.PhotoSelfie,
	}

	err := h.UserService.CreateUser(&user)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to create user", err.Error())
		return
	}
	helper.Success(c, "User created successfully", user)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.UserService.GetUsers()
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to retrieve users", err.Error())
		return
	}
	helper.Success(c, "Users retrieved successfully", users)
}
