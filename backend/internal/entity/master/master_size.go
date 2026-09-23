package entity

import (
	adminsistrator "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/admin"
)

type MasterSize struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Size string `json:"size"`
	adminsistrator.BaseModel
}
