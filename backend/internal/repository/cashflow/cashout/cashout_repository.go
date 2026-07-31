package cashout

import (
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/cashflow/cashout"
	"gorm.io/gorm"
)

type CashOutRepository interface {
	FindAll() ([]entity.CashOut, error)
	FindByID(id uint) (*entity.CashOut, error)
	Create(cashOut *entity.CashOut) error
	Update(cashOut *entity.CashOut) error
	Delete(id uint) error
}

type cashOutRepository struct {
	db *gorm.DB
}

func NewCashOutRepository(db *gorm.DB) CashOutRepository {
	return &cashOutRepository{db: db}
}

func (r *cashOutRepository) FindAll() ([]entity.CashOut, error) {
	var items []entity.CashOut
	err := r.db.Order("date desc").Find(&items).Error
	return items, err
}

func (r *cashOutRepository) FindByID(id uint) (*entity.CashOut, error) {
	var item entity.CashOut
	err := r.db.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *cashOutRepository) Create(cashOut *entity.CashOut) error {
	return r.db.Create(cashOut).Error
}

func (r *cashOutRepository) Update(cashOut *entity.CashOut) error {
	return r.db.Save(cashOut).Error
}

func (r *cashOutRepository) Delete(id uint) error {
	return r.db.Delete(&entity.CashOut{}, id).Error
}