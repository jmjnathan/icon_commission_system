package commission

import "time"

type CommissionPayment struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CommissionID uint      `json:"commission_id"`
	Amount       float64   `json:"amount"`
	PaymentType  string    `json:"payment_type"` // "dp", "cicilan", "pelunasan"
	Method       string    `json:"method"`        // Transfer, Cash, QRIS
	PaidAt       time.Time `json:"paid_at"`
	Notes        string    `json:"notes"`
	CreatedUsername string `json:"created_username"`
}