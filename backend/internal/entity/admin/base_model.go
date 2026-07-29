package admin

import "time"

type BaseModel struct {
	Status          string    `gorm:"default:Active" json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	CreatedUsername string    `json:"created_username"`
	UpdatedAt       time.Time `json:"updated_at"`
	UpdatedUsername string    `json:"updated_username"`
}