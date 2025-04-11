package handler

import (
	"github.com/msyahruls/dgw-go-test/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(router *gin.Engine, db *gorm.DB) {
	authHandler := NewAuthHandler(db)
	userHandler := NewUserHandler(db)

	api := router.Group("/api")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)

		// Protected Routes
		protected := api.Group("/")
		protected.Use(middleware.JWTAuthMiddleware())

		protected.POST("/users", userHandler.CreateUser)
		protected.GET("/users", userHandler.GetUsers)
		protected.GET("/users/:id", userHandler.GetUserByID)
		protected.PATCH("/users/:id", userHandler.UpdateUser)
		protected.DELETE("/users/:id", userHandler.DeleteUser)
	}
}
