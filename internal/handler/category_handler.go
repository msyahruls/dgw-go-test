package handler

import (
	"net/http"

	"github.com/msyahruls/dgw-go-test/internal/dto"
	"github.com/msyahruls/dgw-go-test/internal/helper"
	"github.com/msyahruls/dgw-go-test/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryHandler struct {
	CategoryService service.CategoryService
}

func NewCategoryHandler(db *gorm.DB) *CategoryHandler {
	return &CategoryHandler{
		CategoryService: service.NewCategoryService(db),
	}
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category with input data
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateCategoryRequest true "Category Request Body"
// @Success 200 {object} helper.APIResponseCategory
// @Failure 400 {object} helper.APIResponse
// @Router /api/categories [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	category, err := h.CategoryService.CreateCategory(req)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to create category", err.Error())
		return
	}

	helper.Success(c, "Category created successfully", category)
}

// GetCategories godoc
// @Summary Get list of categories
// @Tags categories
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helper.APIResponseCategories
// @Failure 500 {object} helper.APIResponse
// @Router /api/categories [get]
func (h *CategoryHandler) GetCategories(c *gin.Context) {
	categories, err := h.CategoryService.GetCategories()
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to retrieve categories", err.Error())
		return
	}
	helper.Success(c, "Categories retrieved successfully", categories)
}

// GetCategoryByID godoc
// @Summary Get category by ID
// @Tags categories
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Success 200 {object} helper.APIResponseCategory
// @Failure 404 {object} helper.APIResponse
// @Router /api/categories/{id} [get]
func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := helper.ParseUintParam(idParam)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}

	category, err := h.CategoryService.GetCategoryByID(id)
	if err != nil {
		helper.Error(c, http.StatusNotFound, "Category not found", err.Error())
		return
	}
	helper.Success(c, "Category retrieved successfully", category)
}

// UpdateCategory godoc
// @Summary Update category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Param request body dto.UpdateCategoryRequest true "Update Category Body"
// @Success 200 {object} helper.APIResponseCategory
// @Failure 400 {object} helper.APIResponse
// @Failure 404 {object} helper.APIResponse
// @Router /api/categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := helper.ParseUintParam(idParam)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}

	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	category, err := h.CategoryService.UpdateCategory(id, req)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to update category", err.Error())
		return
	}

	helper.Success(c, "Category updated successfully", category)
}

// DeleteCategory godoc
// @Summary Delete category by ID
// @Tags categories
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Success 200 {object} helper.APIResponse
// @Failure 404 {object} helper.APIResponse
// @Router /api/categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := helper.ParseUintParam(idParam)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}

	if err := h.CategoryService.DeleteCategory(id); err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to delete category", err.Error())
		return
	}
	helper.Success(c, "Category deleted successfully", nil)
}
