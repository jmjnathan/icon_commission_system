package purchase_order

type CreatePurchaseOrderItemRequest struct {
	ProductID 	int     `json:"product_id" validate:"required"`
	ProductName string  `json:"product_name" validate:"required"`
	Quantity  	float64 `json:"quantity" validate:"required,gt=0"`
	UnitPrice 	float64 `json:"unit_price" validate:"required,gte=0"`
	Brand     	*string `json:"brand"`
	Remark    	*string `json:"remark"`
}

type UpdatePurchaseOrderItemRequest struct {
	ProductID 	int     `json:"product_id" validate:"required"`
	ProductName string  `json:"product_name" validate:"required"`
	Quantity  	float64 `json:"quantity" validate:"required,gt=0"`
	UnitPrice 	float64 `json:"unit_price" validate:"required,gte=0"`
	Brand     	*string `json:"brand"`
	Remark    	*string `json:"remark"`
}

type PurchaseOrderItemResponse struct {
	ID              int      `json:"id"`
	PurchaseOrderID int      `json:"purchase_order_id"`
	ProductID       int      `json:"product_id"`
	ProductName     string   `json:"product_name"`
	Brand           *string  `json:"brand"`
	Quantity        float64  `json:"quantity"`
	UnitPrice       float64  `json:"unit_price"`
	Subtotal        float64  `json:"subtotal"`
	Remark          *string  `json:"remark"`
}