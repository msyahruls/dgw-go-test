package handler

import (
	"net/http"
	"strconv"

	"github.com/msyahruls/kreditplus-go-test/internal/domain"
	"github.com/msyahruls/kreditplus-go-test/internal/dto"
	"github.com/msyahruls/kreditplus-go-test/internal/helper"
	"github.com/msyahruls/kreditplus-go-test/internal/repository"
	"github.com/msyahruls/kreditplus-go-test/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TransactionHandler struct {
	TransactionService service.TransactionService
}

func NewTransactionHandler(db *gorm.DB) *TransactionHandler {
	// Initialize repositories
	txRepo := repository.NewTransactionRepository(db)
	limitRepo := repository.NewLimitRepository(db)

	// Inject into service
	txService := service.NewTransactionService(db, txRepo, limitRepo)

	return &TransactionHandler{
		TransactionService: txService,
	}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var req dto.CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	// Receive tenor in query or body (depending on implementation)
	tenorQuery := c.Query("tenor")
	tenor, err := strconv.Atoi(tenorQuery)
	if err != nil || tenor <= 0 {
		helper.Error(c, http.StatusBadRequest, "Invalid tenor", err.Error())
		return
	}

	transaction := domain.Transaction{
		UserID:            req.UserID,
		ContractNumber:    req.ContractNumber,
		OTR:               req.OTR,
		AdminFee:          req.AdminFee,
		InstallmentAmount: req.InstallmentAmount,
		InterestAmount:    req.InterestAmount,
		AssetName:         req.AssetName,
	}

	err = h.TransactionService.CreateTransaction(&transaction, tenor)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to create transaction", err.Error())
		return
	}
	helper.Success(c, "User created successfully", transaction)
}

func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	transactions, err := h.TransactionService.GetTransactions()
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to retrieve transactions", err.Error())
		return
	}
	helper.Success(c, "Transactions retrieved successfully", transactions)
}
