package dto

type CreateTransactionRequest struct {
	UserID            uint    `json:"user_id" binding:"required"`
	ContractNumber    string  `json:"contract_number" binding:"required"`
	OTR               float64 `json:"otr" binding:"required,gt=0"`
	AdminFee          float64 `json:"admin_fee" binding:"required"`
	InstallmentAmount float64 `json:"installment_amount" binding:"required,gt=0"`
	InterestAmount    float64 `json:"interest_amount" binding:"required"`
	AssetName         string  `json:"asset_name" binding:"required"`
}

type PaymentRequest struct {
	ScheduleID  uint   `json:"schedule_id" binding:"required"`
	PaymentDate string `json:"payment_date" binding:"required"`
}
