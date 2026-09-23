package entity

import (
	admin "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/admin"
)

type MaterialComponent struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Remark   string `json:"remark"`

	Variants []MaterialComponentVariant `gorm:"foreignKey:MaterialComponentID" json:"variants"`

	admin.BaseModel
}
