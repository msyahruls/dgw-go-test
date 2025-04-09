package handler

import (
	"github.com/msyahruls/dgw-go-test/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(router *gin.Engine, db *gorm.DB) {
	userHandler := NewUserHandler(db)

	api := router.Group("/api")
	{
		api.POST("/login", LoginHandler)

		// Protected Routes
		protected := api.Group("/")
		protected.Use(middleware.JWTAuthMiddleware())

		protected.POST("/users", userHandler.CreateUser)
		protected.GET("/users", userHandler.GetUsers)
	}
}
