package purchase_order

type CreatePurchaseOrderRequest struct {
	DocumentNo string                           `json:"document_no"`
	VendorName string                           `json:"vendor_name" validate:"required"`
	OrderDate  string                           `json:"order_date" validate:"required"`
	Remark     *string                          `json:"remark"`
	Items      []CreatePurchaseOrderItemRequest `json:"items" validate:"required,min=1"`
}

type UpdatePurchaseOrderRequest struct {
	DocumentNo string                           `json:"document_no"`
	VendorName string                           `json:"vendor_name" validate:"required"`
	OrderDate  string                           `json:"order_date" validate:"required"`
	Remark     *string                          `json:"remark"`
	Items      []UpdatePurchaseOrderItemRequest `json:"items" validate:"required,min=1"`
}

type PurchaseOrderResponse struct {
	ID           int     `json:"id"`
	DocumentName string  `json:"document_name"`
	DocumentNo   string  `json:"document_no"`
	VendorName   string  `json:"vendor_name"`
	OrderDate    string  `json:"order_date"`
	Status       string  `json:"status"`
	Remark       *string `json:"remark"`
	Total        float64 `json:"total"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}