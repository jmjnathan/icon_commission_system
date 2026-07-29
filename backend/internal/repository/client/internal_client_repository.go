package repository

import (
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/client"
	"gorm.io/gorm"
)

type ClientRepository interface {
	FindAll() ([]entity.Client, error)
	FindByID(id uint) (*entity.Client, error)
	Create(client *entity.Client) error
	Update(client *entity.Client) error
	Delete(id uint) error
}

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return &clientRepository{db: db}
}

func (r *clientRepository) FindAll() ([]entity.Client, error) {
	var clients []entity.Client
	err := r.db.Find(&clients).Error
	return clients, err
}

func (r *clientRepository) FindByID(id uint) (*entity.Client, error) {
	var client entity.Client
	err := r.db.First(&client, id).Error
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *clientRepository) Create(client *entity.Client) error {
	return r.db.Create(client).Error
}

func (r *clientRepository) Update(client *entity.Client) error {
	return r.db.Save(client).Error
}

func (r *clientRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Client{}, id).Error
}