package commission

import (
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/commission"
	"gorm.io/gorm"
)

type CommissionRepository interface {
	FindAll() ([]entity.Commission, error)
	FindByID(id uint) (*entity.Commission, error)
	Create(commission *entity.Commission) error
	Update(commission *entity.Commission) error
	Delete(id uint) error
}

type commissionRepository struct {
	db *gorm.DB
}

func NewCommissionRepository(db *gorm.DB) CommissionRepository {
	return &commissionRepository{db: db}
}

func (r *commissionRepository) preloadAll(db *gorm.DB) *gorm.DB {
	return db.
		Preload("Client").
		Preload("Items").
		Preload("Items.Saint").
		Preload("Items.Size").
		Preload("Items.Material").
		Preload("Items.Style")
}

func (r *commissionRepository) FindAll() ([]entity.Commission, error) {
	var commissions []entity.Commission
	err := r.preloadAll(r.db).Find(&commissions).Error
	return commissions, err
}

func (r *commissionRepository) FindByID(id uint) (*entity.Commission, error) {
	var commission entity.Commission
	err := r.preloadAll(r.db).First(&commission, id).Error
	if err != nil {
		return nil, err
	}
	return &commission, nil
}

func (r *commissionRepository) Create(commission *entity.Commission) error {
	return r.db.Create(commission).Error
}

func (r *commissionRepository) Update(commission *entity.Commission) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: false}).Save(commission).Error
}

func (r *commissionRepository) Delete(id uint) error {
	// Hapus item dulu (child), baru header (parent)
	if err := r.db.Where("commission_id = ?", id).Delete(&entity.CommissionItem{}).Error; err != nil {
		return err
	}
	return r.db.Delete(&entity.Commission{}, id).Error
}