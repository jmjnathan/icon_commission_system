package client

import (
	"time"

	adminsistrator "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/admin"
)

type Client struct {
	ID                     uint       `gorm:"primaryKey" json:"id"`
	Name                   string     `gorm:"not null" json:"name"`
	Nickname               string     `json:"nickname"`
	Phone                  string     `gorm:"not null" json:"phone"`
	AlternatePhone         string     `json:"alternate_phone"`
	Email                  string     `json:"email"`
	InstagramHandle        string     `json:"instagram_handle"`
	Address                string     `json:"address"`
	City                   string     `json:"city"`
	Province               string     `json:"province"`
	PostalCode             string     `json:"postal_code"`
	Country                string     `gorm:"default:Indonesia" json:"country"`
	Denomination           string     `json:"denomination"`
	PatronSaintPreference  string     `json:"patron_saint_preference"`
	ParishName             string     `json:"parish_name"`
	FirstOrderDate         *time.Time `json:"first_order_date"`
	TotalOrders            int        `gorm:"default:0" json:"total_orders"`
	PreferredPaymentMethod string     `json:"preferred_payment_method"`
	Notes                  string     `json:"notes"`
	adminsistrator.BaseModel
}
