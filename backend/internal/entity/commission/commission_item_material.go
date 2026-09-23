package commission

import (
	admin "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/admin"
	master "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/master"
)

type CommissionItemMaterial struct {
	ID uint `gorm:"primaryKey" json:"id"`

	CommissionItemID uint `json:"commission_item_id"`

	MaterialComponentVariantID uint `json:"material_component_variant_id"`

	Quantity  float64 `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	TotalCost float64 `json:"total_cost"`

	Remark string `json:"remark"`

	MaterialComponentVariant master.MaterialComponentVariant `gorm:"foreignKey:MaterialComponentVariantID" json:"material_component_variant"`

	admin.BaseModel
}
