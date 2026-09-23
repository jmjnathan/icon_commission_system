package repository

import (
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/master"
	"gorm.io/gorm"
)

type MaterialComponentVariantRepository interface {
	FindAll() ([]entity.MaterialComponentVariant, error)

	FindByMaterialComponentID(
		materialComponentID uint,
	) ([]entity.MaterialComponentVariant, error)

	FindByID(
		id uint,
	) (*entity.MaterialComponentVariant, error)

	Create(
		variant *entity.MaterialComponentVariant,
	) error

	Update(
		variant *entity.MaterialComponentVariant,
	) error

	Delete(id uint) error
}

type materialComponentVariantRepository struct {
	db *gorm.DB
}

func (r *materialComponentVariantRepository) FindAll() ([]entity.MaterialComponentVariant, error) {
	var variants []entity.MaterialComponentVariant

	err := r.db.
		Preload("MaterialComponent").
		Where("status = ?", "Active").
		Order("material_component_id ASC").
		Order("id ASC").
		Find(&variants).Error

	return variants, err
}

func NewMaterialComponentVariantRepository(
	db *gorm.DB,
) MaterialComponentVariantRepository {
	return &materialComponentVariantRepository{db: db}
}

func (r *materialComponentVariantRepository) FindByMaterialComponentID(
	materialComponentID uint,
) ([]entity.MaterialComponentVariant, error) {

	var variants []entity.MaterialComponentVariant

	err := r.db.
		Where("material_component_id = ?", materialComponentID).
		Order("id ASC").
		Find(&variants).Error

	return variants, err
}

func (r *materialComponentVariantRepository) FindByID(
	id uint,
) (*entity.MaterialComponentVariant, error) {

	var variant entity.MaterialComponentVariant

	err := r.db.First(&variant, id).Error

	if err != nil {
		return nil, err
	}

	return &variant, nil
}

func (r *materialComponentVariantRepository) Create(
	variant *entity.MaterialComponentVariant,
) error {
	return r.db.Create(variant).Error
}

func (r *materialComponentVariantRepository) Update(
	variant *entity.MaterialComponentVariant,
) error {
	return r.db.Save(variant).Error
}

func (r *materialComponentVariantRepository) Delete(
	id uint,
) error {
	return r.db.Delete(
		&entity.MaterialComponentVariant{},
		id,
	).Error
}
