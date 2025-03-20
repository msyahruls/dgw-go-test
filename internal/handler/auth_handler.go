package handler

import (
	"net/http"
	"time"

	"github.com/msyahruls/kreditplus-go-test/internal/config"
	"github.com/msyahruls/kreditplus-go-test/internal/helper"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	// Dummy user check
	if req.Username != "admin" || req.Password != "password" {
		helper.Error(c, http.StatusBadRequest, "Invalid credentials", nil)
		return
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "Token creation failed", nil)
		return
	}

	helper.Success(c, "User login successfully", gin.H{"token": tokenString})
}
