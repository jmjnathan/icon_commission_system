package commission

import (
	masterTrx "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/master"
)

type CommissionItem struct {
	ID                         uint                     `gorm:"primaryKey" json:"id"`
	CommissionID               uint                     `json:"commission_id"`
	SaintID                    uint                     `json:"saint_id"`
	Saint                      masterTrx.MasterSaints   `json:"saint" gorm:"foreignKey:SaintID"`
	SizeID                     uint                     `json:"size_id"`
	Size                       masterTrx.MasterSize     `json:"size" gorm:"foreignKey:SizeID"`
	MaterialID                 uint                     `json:"material_id"`
	Material                   masterTrx.MasterMaterial `json:"material" gorm:"foreignKey:MaterialID"`
	MaterialComponentVariantID uint                     `json:"material_component_variant_id"`
	MaterialQuantity           float64                  `json:"material_quantity"`
	MaterialUnitPrice          float64                  `json:"material_unit_price"`
	MaterialCost               float64                  `json:"material_cost"`
	StyleID                    uint                     `json:"style_id"`
	Style                      masterTrx.MasterStyle    `json:"style" gorm:"foreignKey:StyleID"`
	Price                      float64                  `json:"price"`
	Notes                      string                   `json:"notes"`
}
