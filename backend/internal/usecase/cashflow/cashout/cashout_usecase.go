package cashout

import (
	"time"

	dto "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/cashflow/cashout"
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/cashflow/cashout"
	shared"github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/admin"
	repository "github.com/nathanchristiawan02/icon-commission-system-backend/internal/repository/cashflow/cashout"
)

type CashOutUsecase interface {
	GetAll() ([]entity.CashOut, error)
	Create(req dto.CashOutRequest, username string) (*entity.CashOut, error)
	Update(id uint, req dto.CashOutRequest, username string) (*entity.CashOut, error)
	Delete(id uint) error
}

type cashOutUsecase struct {
	repo repository.CashOutRepository
}

func NewCashOutUsecase(repo repository.CashOutRepository) CashOutUsecase {
	return &cashOutUsecase{repo: repo}
}

func (u *cashOutUsecase) GetAll() ([]entity.CashOut, error) {
	return u.repo.FindAll()
}

func (u *cashOutUsecase) Create(req dto.CashOutRequest, username string) (*entity.CashOut, error) {
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, err
	}

	item := &entity.CashOut{
		Category:      req.Category,
		Description:   req.Description,
		Qty:           req.Qty,
		UnitPrice:     req.UnitPrice,
		Amount:        req.Qty * req.UnitPrice,
		Vendor:        req.Vendor,
		PaymentMethod: req.PaymentMethod,
		ReceiptURL:    req.ReceiptURL,
		Notes:         req.Notes,
		Date:          date,
		BaseModel: shared.BaseModel{
			Status:          "Active",
			CreatedUsername: username,
			UpdatedUsername: username,
		},
	}
	err = u.repo.Create(item)
	return item, err
}

func (u *cashOutUsecase) Update(id uint, req dto.CashOutRequest, username string) (*entity.CashOut, error) {
	item, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, err
	}

	item.Category = req.Category
	item.Description = req.Description
	item.Qty = req.Qty
	item.UnitPrice = req.UnitPrice
	item.Amount = req.Qty * req.UnitPrice
	item.Vendor = req.Vendor
	item.PaymentMethod = req.PaymentMethod
	item.ReceiptURL = req.ReceiptURL
	item.Notes = req.Notes
	item.Date = date
	item.UpdatedUsername = username
	item.UpdatedAt = time.Now()

	err = u.repo.Update(item)
	return item, err
}

func (u *cashOutUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}