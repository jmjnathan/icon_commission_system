package repository

import (
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/master"
	"gorm.io/gorm"
)

type MaterialComponentRepository interface {
	FindAll() ([]entity.MaterialComponent, error)
	FindByID(id uint) (*entity.MaterialComponent, error)
	Create(component *entity.MaterialComponent) error
	Update(component *entity.MaterialComponent) error
	Delete(id uint) error
}

type materialComponentRepository struct {
	db *gorm.DB
}

func NewMaterialComponentRepository(db *gorm.DB) MaterialComponentRepository {
	return &materialComponentRepository{db: db}
}

func (r *materialComponentRepository) FindAll() ([]entity.MaterialComponent, error) {
	var components []entity.MaterialComponent

	err := r.db.Find(&components).Error

	return components, err
}

func (r *materialComponentRepository) FindByID(id uint) (*entity.MaterialComponent, error) {
	var component entity.MaterialComponent

	err := r.db.First(&component, id).Error
	if err != nil {
		return nil, err
	}

	return &component, nil
}

func (r *materialComponentRepository) Create(
	component *entity.MaterialComponent,
) error {
	return r.db.Create(component).Error
}

func (r *materialComponentRepository) Update(
	component *entity.MaterialComponent,
) error {
	return r.db.Save(component).Error
}

func (r *materialComponentRepository) Delete(id uint) error {
	return r.db.Delete(&entity.MaterialComponent{}, id).Error
}
