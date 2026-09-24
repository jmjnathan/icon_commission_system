package product

import (
	"context"

	productDto "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/master/product"
)

type ProductUsecase interface {
	GetAll(ctx context.Context) ([]productDto.ProductResponse, error)
	GetByID(ctx context.Context, id int) (*productDto.ProductResponse, error)
	Search(ctx context.Context, keyword string) ([]productDto.ProductResponse, error)

	Create(
		ctx context.Context,
		req productDto.CreateProductRequest,
	) (*productDto.ProductResponse, error)

	Update(
		ctx context.Context,
		id int,
		req productDto.UpdateProductRequest,
	) (*productDto.ProductResponse, error)

	Delete(ctx context.Context, id int) error
}