package entity

import (
	adminsistrator "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/admin"
)

type MasterStyle struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
	adminsistrator.BaseModel
}
