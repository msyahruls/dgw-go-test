package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/msyahruls/dgw-go-test/internal/config"
	"github.com/msyahruls/dgw-go-test/internal/dto"
	"github.com/msyahruls/dgw-go-test/internal/helper"
	"github.com/msyahruls/dgw-go-test/internal/service"
	"gorm.io/gorm"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		AuthService: service.NewAuthService(db),
	}
}

// Register godoc
// @Summary Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Register Request"
// @Success 200 {object} helper.APIResponse
// @Failure 400 {object} helper.APIResponse
// @Router /api/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	user, err := h.AuthService.Register(req)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "Failed to register", err.Error())
		return
	}

	helper.Success(c, "User registered successfully", user)
}

// Login godoc
// @Summary Login and get JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login Credentials"
// @Success 200 {object} helper.APIResponse
// @Failure 400 {object} helper.APIResponse
// @Router /api/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	user, err := h.AuthService.Login(req)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid credentials", err.Error())
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"name":     user.Name,
		"exp":      time.Now().Add(72 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "Token creation failed", err.Error())
		return
	}

	helper.Success(c, "Login successful", gin.H{"token": tokenString})
}
