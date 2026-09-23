package commission

import (
	"time"

	shared "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/admin"
	clientEntity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/client"
)

type Commission struct {
	ID            uint              `gorm:"primaryKey" json:"id"`
	ClientID      uint              `json:"client_id"`
	Client        clientEntity.Client `json:"client" gorm:"foreignKey:ClientID"`
	OrderDate     time.Time         `json:"order_date"`
	Deadline      time.Time         `json:"deadline"`
	Notes         string            `json:"notes"`
	Status        string            `gorm:"default:pending" json:"status"`
	Subtotal      float64           `json:"subtotal"`       // total harga item sebelum diskon
	DiscountType  string            `json:"discount_type"`   // "percent" atau "fixed"
	DiscountValue float64           `json:"discount_value"`  // nilai diskon (misal 10 untuk 10%, atau 50000 untuk potongan tetap)
	TotalPrice    float64           `json:"total_price"`     // subtotal - diskon (harga final)
	PaymentStatus string            `gorm:"default:unpaid" json:"payment_status"` // unpaid, partial, paid
	Items         []CommissionItem  `json:"items" gorm:"foreignKey:CommissionID"`
	Photos        []CommissionPhoto `json:"photos" gorm:"foreignKey:CommissionID"`
	Payments      []CommissionPayment `json:"payments" gorm:"foreignKey:CommissionID"`
	shared.BaseModel
}