package handler

import (
	"net/http"

	"github.com/msyahruls/dgw-go-test/internal/domain"
	"github.com/msyahruls/dgw-go-test/internal/dto"
	"github.com/msyahruls/dgw-go-test/internal/helper"
	"github.com/msyahruls/dgw-go-test/internal/service"

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

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with input data
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateUserRequest true "User Request Body"
// @Success 200 {object} helper.APIResponse
// @Failure 400 {object} helper.APIResponse
// @Router /api/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	user := domain.User{
		Name: req.Name,
	}

	err := h.UserService.CreateUser(&user)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to create user", err.Error())
		return
	}
	helper.Success(c, "User created successfully", user)
}

// GetUsers godoc
// @Summary Get list of users
// @Tags users
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helper.APIResponse
// @Failure 500 {object} helper.APIResponse
// @Router /api/users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.UserService.GetUsers()
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to retrieve users", err.Error())
		return
	}
	helper.Success(c, "Users retrieved successfully", users)
}
