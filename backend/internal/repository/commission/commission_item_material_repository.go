package commission

import (
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/commission"
	"gorm.io/gorm"
)

type CommissionItemMaterialRepository interface {
	FindAllByCommissionItemID(commissionItemID uint) ([]entity.CommissionItemMaterial, error)
	FindByID(id uint) (*entity.CommissionItemMaterial, error)
	Create(material *entity.CommissionItemMaterial) error
	Update(material *entity.CommissionItemMaterial) error
	Delete(id uint) error
}

type commissionItemMaterialRepository struct {
	db *gorm.DB
}

func NewCommissionItemMaterialRepository(
	db *gorm.DB,
) CommissionItemMaterialRepository {
	return &commissionItemMaterialRepository{
		db: db,
	}
}

func (r *commissionItemMaterialRepository) FindAllByCommissionItemID(
	commissionItemID uint,
) ([]entity.CommissionItemMaterial, error) {

	var materials []entity.CommissionItemMaterial

	err := r.db.
		Preload("MaterialComponentVariant").
		Where("commission_item_id = ?", commissionItemID).
		Order("id ASC").
		Find(&materials).Error

	return materials, err
}

func (r *commissionItemMaterialRepository) FindByID(
	id uint,
) (*entity.CommissionItemMaterial, error) {

	var material entity.CommissionItemMaterial

	err := r.db.
		Preload("MaterialComponentVariant").
		First(&material, id).Error

	if err != nil {
		return nil, err
	}

	return &material, nil
}

func (r *commissionItemMaterialRepository) Create(
	material *entity.CommissionItemMaterial,
) error {
	return r.db.Create(material).Error
}

func (r *commissionItemMaterialRepository) Update(
	material *entity.CommissionItemMaterial,
) error {
	return r.db.Save(material).Error
}

func (r *commissionItemMaterialRepository) Delete(
	id uint,
) error {
	return r.db.Delete(
		&entity.CommissionItemMaterial{},
		id,
	).Error
}
