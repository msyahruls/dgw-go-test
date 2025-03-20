package handler

import (
	"net/http"
	"strconv"

	"github.com/msyahruls/kreditplus-go-test/internal/domain"
	"github.com/msyahruls/kreditplus-go-test/internal/dto"
	"github.com/msyahruls/kreditplus-go-test/internal/helper"
	"github.com/msyahruls/kreditplus-go-test/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LimitHandler struct {
	LimitService service.LimitService
}

func NewLimitHandler(db *gorm.DB) *LimitHandler {
	return &LimitHandler{
		LimitService: service.NewLimitService(db),
	}
}

func (h *LimitHandler) CreateOrUpdateLimit(c *gin.Context) {
	var req dto.CreateLimitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	limit := domain.Limit{
		UserID:      req.UserID,
		TenorMonths: req.TenorMonths,
		LimitAmount: req.LimitAmount,
	}

	err := h.LimitService.CreateOrUpdateLimit(&limit)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	helper.Success(c, "Limit saved successfully", limit)
}

func (h *LimitHandler) GetLimits(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid user_id", err.Error())
		return
	}

	limits, err := h.LimitService.GetLimits(uint(userID))
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to retrieve limits", err.Error())
		return
	}

	helper.Success(c, "Limits retrieved successfully", limits)
}
