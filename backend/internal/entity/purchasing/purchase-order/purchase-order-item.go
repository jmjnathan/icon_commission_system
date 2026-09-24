package purchase_order

import "time"

type PurchaseOrderItem struct {
	ID              int       `json:"id"`
	PurchaseOrderID int       `json:"purchase_order_id"`
	ProductID       int       `json:"product_id"`
	ProductName     string    `json:"product_name"`
	Brand           *string   `json:"brand"`
	Quantity        float64   `json:"quantity"`
	UnitPrice       float64   `json:"unit_price"`
	Subtotal        float64   `json:"subtotal"`
	Remark          *string   `json:"remark"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}