package handler

import (
	"github.com/msyahruls/kreditplus-go-test/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(router *gin.Engine, db *gorm.DB) {
	userHandler := NewUserHandler(db)
	txHandler := NewTransactionHandler(db)
	limitHandler := NewLimitHandler(db)

	api := router.Group("/api")
	{
		api.POST("/login", LoginHandler)

		// Protected Routes
		protected := api.Group("/")
		protected.Use(middleware.JWTAuthMiddleware())

		protected.POST("/users", userHandler.CreateUser)
		protected.GET("/users", userHandler.GetUsers)

		protected.POST("/transactions", txHandler.CreateTransaction)
		protected.GET("/transactions", txHandler.GetTransactions)
		protected.GET("/transactions/:id/schedules", txHandler.GetPaymentSchedules)
		protected.POST("/payments", txHandler.PayInstallment)

		protected.POST("/limits", limitHandler.CreateOrUpdateLimit)
		protected.GET("/limits/:user_id", limitHandler.GetLimits)
	}
}
