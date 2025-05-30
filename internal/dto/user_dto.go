package dto

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
}
