package entity

import (
	admin "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/admin"
)

type MaterialComponentVariant struct {
	ID uint `gorm:"primaryKey" json:"id"`

	MaterialComponentID uint `json:"material_component_id"`

	Specification string `json:"specification"`

	Length    float64 `json:"length"`
	Width     float64 `json:"width"`
	Thickness float64 `json:"thickness"`

	Unit      string  `json:"unit"`
	UnitPrice float64 `json:"unit_price"`

	Remark string `json:"remark"`

	MaterialComponent MaterialComponent `gorm:"foreignKey:MaterialComponentID" json:"material_component"`

	admin.BaseModel
}
