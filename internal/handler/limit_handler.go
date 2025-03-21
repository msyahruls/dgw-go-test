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
// @Param tenor query int true "Tenor"
// @Success 200 {object} helper.APIResponse
// @Failure 400 {object} helper.APIResponse
// @Router /api/limits [get]
func (h *LimitHandler) GetLimits(c *gin.Context) {
	userIDStr := c.Query("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		userID = 0
	}

	limits, err := h.LimitService.GetLimits(uint(userID))
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to retrieve limits", err.Error())
		return
	}

	helper.Success(c, "Limits retrieved successfully", limits)
}

// GetUserLimits godoc
// @Summary Get user limits by user ID
// @Description Retrieve remaining limits per tenor for a user
// @Tags limits
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} helper.APIResponse
// @Failure 500 {object} helper.APIResponse
// @Router /users/{id}/limits [get]
func (h *LimitHandler) GetUserLimits(c *gin.Context) {
	idParam := c.Param("id")
	userID, _ := strconv.Atoi(idParam)

	limits, err := h.LimitService.GetUserLimits(uint(userID))
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to retrieve limits", err.Error())
		return
	}

	helper.Success(c, "User limits retrieved successfully", limits)
}
