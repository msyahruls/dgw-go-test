package handler

import (
	"net/http"

	"github.com/msyahruls/dgw-go-test/internal/dto"
	"github.com/msyahruls/dgw-go-test/internal/helper"
	"github.com/msyahruls/dgw-go-test/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductHandler struct {
	ProductService service.ProductService
}

func NewProductHandler(db *gorm.DB) *ProductHandler {
	return &ProductHandler{
		ProductService: service.NewProductService(db),
	}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with input data
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateProductRequest true "Product Request Body"
// @Success 200 {object} helper.APIResponse
// @Failure 400 {object} helper.APIResponse
// @Router /api/products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	product, err := h.ProductService.CreateProduct(req)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to create product", err.Error())
		return
	}

	helper.Success(c, "Product created successfully", product)
}

// GetProducts godoc
// @Summary Get list of products
// @Tags products
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helper.APIResponse
// @Failure 500 {object} helper.APIResponse
// @Router /api/products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.ProductService.GetProducts()
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to retrieve products", err.Error())
		return
	}
	helper.Success(c, "Products retrieved successfully", products)
}

// GetProductByID godoc
// @Summary Get product by ID
// @Tags products
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200 {object} helper.APIResponse
// @Failure 404 {object} helper.APIResponse
// @Router /api/products/{id} [get]
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := helper.ParseUintParam(idParam)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}

	product, err := h.ProductService.GetProductByID(id)
	if err != nil {
		helper.Error(c, http.StatusNotFound, "Product not found", err.Error())
		return
	}
	helper.Success(c, "Product retrieved successfully", product)
}

// UpdateProduct godoc
// @Summary Update product by ID
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Param request body dto.UpdateProductRequest true "Update Product Body"
// @Success 200 {object} helper.APIResponse
// @Failure 400 {object} helper.APIResponse
// @Failure 404 {object} helper.APIResponse
// @Router /api/products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := helper.ParseUintParam(idParam)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}

	var req dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	product, err := h.ProductService.UpdateProduct(id, req)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to update product", err.Error())
		return
	}

	helper.Success(c, "Product updated successfully", product)
}

// DeleteProduct godoc
// @Summary Delete product by ID
// @Tags products
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200 {object} helper.APIResponse
// @Failure 404 {object} helper.APIResponse
// @Router /api/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := helper.ParseUintParam(idParam)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}

	if err := h.ProductService.DeleteProduct(id); err != nil {
		helper.Error(c, http.StatusInternalServerError, "Failed to delete product", err.Error())
		return
	}
	helper.Success(c, "Product deleted successfully", nil)
}
