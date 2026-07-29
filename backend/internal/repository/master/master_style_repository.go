package repository

import (
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/master"
	"gorm.io/gorm"
)

type MasterStyleRepository interface {
	FindAll() ([]entity.MasterStyle, error)
	FindByID(id uint) (*entity.MasterStyle, error)
	Create(size *entity.MasterStyle) error
	Update(size *entity.MasterStyle) error
	Delete(id uint) error
}

type masterStyleRepository struct {
	db *gorm.DB
}

func NewMasterStyleRepository(db *gorm.DB) MasterStyleRepository {
	return &masterStyleRepository{db: db}
}

func (r *masterStyleRepository) FindAll() ([]entity.MasterStyle, error) {
	var sizes []entity.MasterStyle
	err := r.db.Find(&sizes).Error
	return sizes, err
}

func (r *masterStyleRepository) FindByID(id uint) (*entity.MasterStyle, error) {
	var size entity.MasterStyle
	err := r.db.First(&size, id).Error
	if err != nil {
		return nil, err
	}
	return &size, nil
}

func (r *masterStyleRepository) Create(size *entity.MasterStyle) error {
	return r.db.Create(size).Error
}

func (r *masterStyleRepository) Update(size *entity.MasterStyle) error {
	return r.db.Save(size).Error
}

func (r *masterStyleRepository) Delete(id uint) error {
	return r.db.Delete(&entity.MasterStyle{}, id).Error
}