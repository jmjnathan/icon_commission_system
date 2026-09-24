package product

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Unit        string  `json:"unit" validate:"required"`
	Description *string `json:"description"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Unit        string  `json:"unit" validate:"required"`
	Description *string `json:"description"`
	IsActive    bool    `json:"is_active"`
}

type ProductResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Unit        string  `json:"unit"`
	Description *string `json:"description"`
	IsActive    bool    `json:"is_active"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}