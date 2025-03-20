package dto

type CreateLimitRequest struct {
	UserID      uint    `json:"user_id" binding:"required"`
	TenorMonths int     `json:"tenor_months" binding:"required,oneof=1 2 3 4"`
	LimitAmount float64 `json:"limit_amount" binding:"required,gt=0"`
}
