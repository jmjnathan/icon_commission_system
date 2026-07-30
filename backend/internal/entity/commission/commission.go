package commission

import (
	"time"

	shared "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/admin"
	clientEntity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/client"
)

type Commission struct {
	ID         uint             `gorm:"primaryKey" json:"id"`
	ClientID   uint             `json:"client_id"`
	Client     clientEntity.Client `json:"client" gorm:"foreignKey:ClientID"`
	OrderDate  time.Time        `json:"order_date"`
	Deadline   time.Time        `json:"deadline"`
	Notes      string           `json:"notes"`
	Status     string           `gorm:"default:pending" json:"status"`
	TotalPrice float64          `json:"total_price"`
	Items      []CommissionItem `json:"items" gorm:"foreignKey:CommissionID"`
	Photos []CommissionPhoto `json:"photos" gorm:"foreignKey:CommissionID"`
	shared.BaseModel
}