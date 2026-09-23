package dto

type MaterialComponentRequest struct {
	Name     string `json:"name" binding:"required"`
	Category string `json:"category" binding:"required"`
	Remark   string `json:"remark"`
	Status   string `json:"status"`
}

type MaterialComponentVariantRequest struct {
	Specification string  `json:"specification"`
	Length        float64 `json:"length"`
	Width         float64 `json:"width"`
	Thickness     float64 `json:"thickness"`
	Unit          string  `json:"unit" binding:"required"`
	UnitPrice     float64 `json:"unit_price" binding:"required,min=0"`
	Remark        string  `json:"remark"`
	Status        string  `json:"status"`
}
