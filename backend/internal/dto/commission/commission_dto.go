package commission

type CommissionItemRequest struct {
	SaintID    uint    `json:"saint_id" binding:"required"`
	SizeID     uint    `json:"size_id" binding:"required"`
	MaterialID uint    `json:"material_id" binding:"required"`
	StyleID    uint    `json:"style_id" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	Notes      string  `json:"notes"`
}

type CommissionRequest struct {
	ClientID      uint                    `json:"client_id" binding:"required"`
	Deadline      string                  `json:"deadline" binding:"required"`
	Notes         string                  `json:"notes"`
	Items         []CommissionItemRequest `json:"items" binding:"required,min=1,dive"`
	PhotoUrls     []string                `json:"photo_urls"`
	DiscountType  string                  `json:"discount_type"`  // opsional: "percent" / "fixed"
	DiscountValue float64                 `json:"discount_value"` // opsional
}

type CommissionPaymentRequest struct {
	Amount      float64 `json:"amount" binding:"required"`
	PaymentType string  `json:"payment_type" binding:"required"` // dp, cicilan, pelunasan
	Method      string  `json:"method" binding:"required"`
	Notes       string  `json:"notes"`
}

type CommissionStatusRequest struct {
	Status string `json:"status" binding:"required"`
}
