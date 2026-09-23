package repository

import (
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/admin"
	"gorm.io/gorm"
)

type AdminRepository interface {
	FindByUsername(username string) (*entity.Admin, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) FindByUsername(username string) (*entity.Admin, error) {
	var admin entity.Admin
	err := r.db.Where("username = ?", username).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
