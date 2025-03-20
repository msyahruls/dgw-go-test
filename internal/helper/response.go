package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// Success Response
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
		Error:   nil,
	})
}

// Error Response
func Error(c *gin.Context, statusCode int, message string, err interface{}) {
	c.JSON(statusCode, APIResponse{
		Success: false,
		Message: message,
		Data:    nil,
		Error:   err,
	})
}
