package cashout

type CashOutRequest struct {
	Category      string  `json:"category" binding:"required"`
	Description   string  `json:"description" binding:"required"`
	Qty           float64 `json:"qty" binding:"required"`
	UnitPrice     float64 `json:"unit_price" binding:"required"`
	Vendor        string  `json:"vendor"`
	PaymentMethod string  `json:"payment_method"`
	ReceiptURL    string  `json:"receipt_url"`
	Notes         string  `json:"notes"`
	Date          string  `json:"date" binding:"required"`
}