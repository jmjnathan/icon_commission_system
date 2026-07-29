package repository

import (
	"github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/master"
	"gorm.io/gorm"
)

type MasterMaterialRepository interface {
	FindAll() ([]entity.MasterMaterial, error)
	FindByID(id uint) (*entity.MasterMaterial, error)
	Create(size *entity.MasterMaterial) error
	Update(size *entity.MasterMaterial) error
	Delete(id uint) error
}

type masterMaterialRepository struct {
	db *gorm.DB
}

func NewMasterMaterialRepository(db *gorm.DB) MasterMaterialRepository {
	return &masterMaterialRepository{db: db}
}

func (r *masterMaterialRepository) FindAll() ([]entity.MasterMaterial, error) {
	var sizes []entity.MasterMaterial
	err := r.db.Find(&sizes).Error
	return sizes, err
}

func (r *masterMaterialRepository) FindByID(id uint) (*entity.MasterMaterial, error) {
	var size entity.MasterMaterial
	err := r.db.First(&size, id).Error
	if err != nil {
		return nil, err
	}
	return &size, nil
}

func (r *masterMaterialRepository) Create(size *entity.MasterMaterial) error {
	return r.db.Create(size).Error
}

func (r *masterMaterialRepository) Update(size *entity.MasterMaterial) error {
	return r.db.Save(size).Error
}

func (r *masterMaterialRepository) Delete(id uint) error {
	return r.db.Delete(&entity.MasterMaterial{}, id).Error
}