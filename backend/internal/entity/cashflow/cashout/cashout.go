package cashout

import (
	"time"

	shared "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/admin"
)

type CashOut struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Category      string    `json:"category"`
	Description   string    `json:"description"`
	Qty           float64   `json:"qty"`
	UnitPrice     float64   `json:"unit_price"`
	Amount        float64   `json:"amount"`
	Vendor        string    `json:"vendor"`
	PaymentMethod string    `json:"payment_method"`
	ReceiptURL    string    `json:"receipt_url"`
	Notes         string    `json:"notes"`
	Date          time.Time `json:"date"`
	shared.BaseModel
}
