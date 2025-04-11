package handler

import (
	"net/http"

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

	user, err := h.UserService.CreateUser(req)
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

// GetUserByID godoc
// @Summary Get user by ID
// @Tags users
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} helper.APIResponse
// @Failure 404 {object} helper.APIResponse
// @Router /api/users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := helper.ParseUintParam(idParam)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}

	user, err := h.UserService.GetUserByID(id)
	if err != nil {
		helper.Error(c, http.StatusNotFound, "User not found", err.Error())
		return
	}
	helper.Success(c, "User retrieved successfully", user)
}

// UpdateUser godoc
// @Summary Update user by ID
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Param request body dto.UpdateUserRequest true "Update User Body"
// @Success 200 {object} helper.APIResponse
// @Failure 400 {object} helper.APIResponse
// @Failure 404 {object} helper.APIResponse
// @Router /api/users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := helper.ParseUintParam(idParam)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	user, err := h.UserService.UpdateUser(id, req)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to update user", err.Error())
		return
	}

	helper.Success(c, "User updated successfully", user)
}

// DeleteUser godoc
// @Summary Delete user by ID
// @Tags users
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} helper.APIResponse
// @Failure 404 {object} helper.APIResponse
// @Router /api/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := helper.ParseUintParam(idParam)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}

	if err := h.UserService.DeleteUser(id); err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to delete user", err.Error())
		return
	}
	helper.Success(c, "User deleted successfully", nil)
}
