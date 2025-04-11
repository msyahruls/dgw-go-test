package dto

type CreateProductRequest struct {
	Name       string  `json:"name" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	CategoryID uint    `json:"category_id" binding:"required"`
}

type UpdateProductRequest struct {
	Name       string  `json:"name" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	CategoryID uint    `json:"category_id" binding:"required"`
}
