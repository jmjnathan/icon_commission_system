package commission

type CommissionItemMaterialRequest struct {
	MaterialComponentVariantID uint    `json:"material_component_variant_id" binding:"required"`
	Quantity                   float64 `json:"quantity" binding:"required,gt=0"`
	Remark                     string  `json:"remark"`
}
