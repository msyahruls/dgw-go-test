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

// CreateTransaction godoc
// @Summary Create a transaction
// @Description Create a new transaction and deduct user limit
// @Tags transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param tenor query int true "Tenor"
// @Param request body dto.CreateTransactionRequest true "Transaction Request Body"
// @Success 200 {object} helper.APIResponse
// @Failure 400 {object} helper.APIResponse
// @Failure 500 {object} helper.APIResponse
// @Router /api/transactions [post]
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

// GetTransactions godoc
// @Summary Get all transactions
// @Tags transactions
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helper.APIResponse
// @Failure 500 {object} helper.APIResponse
// @Router /api/transactions [get]
func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	transactions, err := h.TransactionService.GetTransactions()
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to retrieve transactions", err.Error())
		return
	}
	helper.Success(c, "Transactions retrieved successfully", transactions)
}

// GetPaymentSchedules godoc
// @Summary Get payment schedules for a transaction
// @Description Retrieve list of payment schedules (installments) by transaction ID
// @Tags transactions
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} helper.APIResponse
// @Failure 500 {object} helper.APIResponse
// @Router /transactions/{id}/schedules [get]
func (h *TransactionHandler) GetPaymentSchedules(c *gin.Context) {
	idParam := c.Param("id")
	txID, _ := strconv.Atoi(idParam)

	schedules, err := h.TransactionService.GetPaymentSchedules(uint(txID))
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to retrieve schedules", err.Error())
		return
	}

	helper.Success(c, "Payment schedules retrieved successfully", schedules)
}

// PayInstallment godoc
// @Summary Pay specific installment
// @Description Mark an installment schedule as paid
// @Tags payments
// @Accept json
// @Produce json
// @Param request body dto.PaymentRequest true "Payment Request"
// @Success 200 {object} helper.APIResponse
// @Failure 400 {object} helper.APIResponse
// @Failure 500 {object} helper.APIResponse
// @Router /payments [post]
func (h *TransactionHandler) PayInstallment(c *gin.Context) {
	var req dto.PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	err := h.TransactionService.PayInstallment(req.ScheduleID, helper.ParseDate(req.PaymentDate))
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to process payment", err.Error())
		return
	}

	helper.Success(c, "Installment paid successfully", nil)
}
