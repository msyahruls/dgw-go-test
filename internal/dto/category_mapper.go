package dto

import "github.com/msyahruls/dgw-go-test/internal/domain"

func ToCategoryResponse(c *domain.Category) CategoryResponse {
	return CategoryResponse{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func ToCategoryResponses(categories []domain.Category) []CategoryResponse {
	res := make([]CategoryResponse, len(categories))
	for i, u := range categories {
		res[i] = ToCategoryResponse(&u)
	}
	return res
}
