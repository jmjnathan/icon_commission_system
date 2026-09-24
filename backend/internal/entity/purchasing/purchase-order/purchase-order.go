package purchase_order

import "time"

const (
	StatusDraft     = "draft"
	StatusOrdered   = "ordered"
	StatusReceived  = "received"
	StatusClosed    = "closed"
	StatusCancelled = "cancelled"
)

type PurchaseOrder struct {
	ID           int       `json:"id"`
	DocumentName string    `json:"document_name"`
	DocumentNo   string    `json:"document_no"`
	VendorName   string    `json:"vendor_name"`
	OrderDate    time.Time `json:"order_date"`
	Status       string    `json:"status"`
	Remark       *string   `json:"remark"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}