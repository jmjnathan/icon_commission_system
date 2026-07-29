package repository

import (
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/master"
	"gorm.io/gorm"
)

type MasterSaintsRepository interface {
	FindAll() ([]entity.MasterSaints, error)
	FindByID(id uint) (*entity.MasterSaints, error)
	Create(size *entity.MasterSaints) error
	Update(size *entity.MasterSaints) error
	Delete(id uint) error
}

type masterSaintsRepository struct {
	db *gorm.DB
}

func NewMasterSaintsRepository(db *gorm.DB) MasterSaintsRepository {
	return &masterSaintsRepository{db: db}
}

func (r *masterSaintsRepository) FindAll() ([]entity.MasterSaints, error) {
	var sizes []entity.MasterSaints
	err := r.db.Find(&sizes).Error
	return sizes, err
}

func (r *masterSaintsRepository) FindByID(id uint) (*entity.MasterSaints, error) {
	var size entity.MasterSaints
	err := r.db.First(&size, id).Error
	if err != nil {
		return nil, err
	}
	return &size, nil
}

func (r *masterSaintsRepository) Create(size *entity.MasterSaints) error {
	return r.db.Create(size).Error
}

func (r *masterSaintsRepository) Update(size *entity.MasterSaints) error {
	return r.db.Save(size).Error
}

func (r *masterSaintsRepository) Delete(id uint) error {
	return r.db.Delete(&entity.MasterSaints{}, id).Error
}