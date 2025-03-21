package domain

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	UserID            uint           `gorm:"not null" json:"user_id"`
	ContractNumber    string         `gorm:"unique;not null" json:"contract_number"`
	OTR               float64        `json:"otr"`
	AdminFee          float64        `json:"admin_fee"`
	InstallmentAmount float64        `json:"installment_amount"`
	InterestAmount    float64        `json:"interest_amount"`
	AssetName         string         `json:"asset_name"`
	TenorMonths       int            `json:"tenor_months"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type PaymentSchedule struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	TransactionID uint           `gorm:"not null" json:"transaction_id"`
	DueDate       time.Time      `json:"due_date"`
	Amount        float64        `json:"amount"`
	Status        string         `json:"status"` // "UNPAID", "PAID", "OVERDUE"
	PaymentDate   *time.Time     `json:"payment_date"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
