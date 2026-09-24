package product

import "time"

type MasterProduct struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Unit        string    `json:"unit"`
	Description *string   `json:"description"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}