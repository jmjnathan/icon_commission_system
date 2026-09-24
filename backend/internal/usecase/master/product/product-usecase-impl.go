package product

import (
	"context"
	"errors"
	"strings"
	"time"

	productDto "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/master/product"
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/master/product"
	repository "github.com/nathanchristiawan02/icon-commission-system-backend/internal/repository/master/product"
)

type productUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(
	repository repository.ProductRepository,
) ProductUsecase {
	return &productUsecase{
		repository: repository,
	}
}

func (u *productUsecase) GetAll(
	ctx context.Context,
) ([]productDto.ProductResponse, error) {

	products, err := u.repository.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make(
		[]productDto.ProductResponse,
		0,
		len(products),
	)

	for _, product := range products {
		responses = append(
			responses,
			mapProductResponse(product),
		)
	}

	return responses, nil
}

func (u *productUsecase) GetByID(
	ctx context.Context,
	id int,
) (*productDto.ProductResponse, error) {

	if id <= 0 {
		return nil, errors.New("invalid product id")
	}

	product, err := u.repository.FindByID(uint(id))
	if err != nil {
		return nil, err
	}

	response := mapProductResponse(*product)

	return &response, nil
}

func (u *productUsecase) Search(
	ctx context.Context,
	keyword string,
) ([]productDto.ProductResponse, error) {

	keyword = strings.TrimSpace(keyword)

	if keyword == "" {
		return []productDto.ProductResponse{}, nil
	}

	products, err := u.repository.Search(keyword)
	if err != nil {
		return nil, err
	}

	responses := make(
		[]productDto.ProductResponse,
		0,
		len(products),
	)

	for _, product := range products {
		responses = append(
			responses,
			mapProductResponse(product),
		)
	}

	return responses, nil
}

func (u *productUsecase) Create(
	ctx context.Context,
	req productDto.CreateProductRequest,
) (*productDto.ProductResponse, error) {

	name := strings.TrimSpace(req.Name)
	unit := strings.TrimSpace(req.Unit)

	if name == "" {
		return nil, errors.New("product name is required")
	}

	if unit == "" {
		return nil, errors.New("product unit is required")
	}

	product := &entity.MasterProduct{
		Name:        name,
		Unit:        unit,
		Description: req.Description,
		IsActive:    true,
	}

	err := u.repository.Create(product)
	if err != nil {
		return nil, err
	}

	response := mapProductResponse(*product)

	return &response, nil
}

func (u *productUsecase) Update(
	ctx context.Context,
	id int,
	req productDto.UpdateProductRequest,
) (*productDto.ProductResponse, error) {

	if id <= 0 {
		return nil, errors.New("invalid product id")
	}

	product, err := u.repository.FindByID(uint(id))
	if err != nil {
		return nil, err
	}

	name := strings.TrimSpace(req.Name)
	unit := strings.TrimSpace(req.Unit)

	if name == "" {
		return nil, errors.New("product name is required")
	}

	if unit == "" {
		return nil, errors.New("product unit is required")
	}

	product.Name = name
	product.Unit = unit
	product.Description = req.Description
	product.IsActive = req.IsActive

	err = u.repository.Update(product)
	if err != nil {
		return nil, err
	}

	response := mapProductResponse(*product)

	return &response, nil
}

func (u *productUsecase) Delete(
	ctx context.Context,
	id int,
) error {

	if id <= 0 {
		return errors.New("invalid product id")
	}

	_, err := u.repository.FindByID(uint(id))
	if err != nil {
		return err
	}

	return u.repository.Delete(uint(id))
}

func mapProductResponse(
	product entity.MasterProduct,
) productDto.ProductResponse {

	return productDto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Unit:        product.Unit,
		Description: product.Description,
		IsActive:    product.IsActive,
		CreatedAt:   product.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   product.UpdatedAt.Format(time.RFC3339),
	}
}