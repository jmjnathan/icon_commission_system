package product

import (
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/master/product"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]entity.MasterProduct, error)
	FindByID(id uint) (*entity.MasterProduct, error)
	Search(keyword string) ([]entity.MasterProduct, error)

	Create(product *entity.MasterProduct) error
	Update(product *entity.MasterProduct) error
	Delete(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) FindAll() ([]entity.MasterProduct, error) {
	var products []entity.MasterProduct

	err := r.db.
		Order("name ASC").
		Find(&products).Error

	return products, err
}

func (r *productRepository) FindByID(id uint) (*entity.MasterProduct, error) {
	var product entity.MasterProduct

	err := r.db.
		First(&product, id).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) Search(
	keyword string,
) ([]entity.MasterProduct, error) {

	var products []entity.MasterProduct

	err := r.db.
		Where("name ILIKE ?", "%"+keyword+"%").
		Where("is_active = ?", true).
		Order("name ASC").
		Find(&products).Error

	return products, err
}

func (r *productRepository) Create(
	product *entity.MasterProduct,
) error {
	return r.db.Create(product).Error
}

func (r *productRepository) Update(
	product *entity.MasterProduct,
) error {

	return r.db.
		Model(&entity.MasterProduct{}).
		Where("id = ?", product.ID).
		Updates(map[string]interface{}{
			"name":        product.Name,
			"unit":        product.Unit,
			"description": product.Description,
			"is_active":   product.IsActive,
		}).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.
		Delete(&entity.MasterProduct{}, id).
		Error
}