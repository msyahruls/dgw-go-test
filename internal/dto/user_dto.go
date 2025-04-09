package dto

type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
}
