package repository

import (
	"github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity"
	"gorm.io/gorm"
)

type MasterSizeRepository interface {
	FindAll() ([]entity.MasterSize, error)
	FindByID(id uint) (*entity.MasterSize, error)
	Create(size *entity.MasterSize) error
	Update(size *entity.MasterSize) error
	Delete(id uint) error
}

type masterSizeRepository struct {
	db *gorm.DB
}

func NewMasterSizeRepository(db *gorm.DB) MasterSizeRepository {
	return &masterSizeRepository{db: db}
}

func (r *masterSizeRepository) FindAll() ([]entity.MasterSize, error) {
	var sizes []entity.MasterSize
	err := r.db.Find(&sizes).Error
	return sizes, err
}

func (r *masterSizeRepository) FindByID(id uint) (*entity.MasterSize, error) {
	var size entity.MasterSize
	err := r.db.First(&size, id).Error
	if err != nil {
		return nil, err
	}
	return &size, nil
}

func (r *masterSizeRepository) Create(size *entity.MasterSize) error {
	return r.db.Create(size).Error
}

func (r *masterSizeRepository) Update(size *entity.MasterSize) error {
	return r.db.Save(size).Error
}

func (r *masterSizeRepository) Delete(id uint) error {
	return r.db.Delete(&entity.MasterSize{}, id).Error
}