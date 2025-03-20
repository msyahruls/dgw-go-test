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

// CreateOrUpdateLimit godoc
// @Summary Create or update limit for user
// @Description Create a new limit or update existing limit for a specific user and tenor
// @Tags limits
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateLimitRequest true "Limit Request Body"
// @Success 200 {object} helper.APIResponse
// @Failure 400 {object} helper.APIResponse
// @Router /api/limits [post]
func (h *LimitHandler) CreateOrUpdateLimit(c *gin.Context) {
	var req dto.CreateLimitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	limit := domain.Limit{
		UserID:      req.UserID,
		TenorMonths: req.TenorMonths,
		LimitAmount: req.LimitAmount,
	}

	err := h.LimitService.CreateOrUpdateLimit(&limit)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to save limit", err.Error())
		return
	}

	helper.Success(c, "Limit saved successfully", limit)
}

// GetUserLimits godoc
// @Summary Get user limits
// @Tags limits
// @Produce json
// @Security BearerAuth
// @Param user_id path int true "User ID"
// @Success 200 {object} helper.APIResponse
// @Failure 400 {object} helper.APIResponse
// @Router /api/limits/{user_id} [get]
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
