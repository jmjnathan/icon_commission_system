package commission

import (
	"time"

	dto "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/commission"
	shared "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/admin"
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/commission"
	repository "github.com/nathanchristiawan02/icon-commission-system-backend/internal/repository/commission"
)

type CommissionUsecase interface {
	GetAll() ([]entity.Commission, error)
	Create(req dto.CommissionRequest, username string) (*entity.Commission, error)
	Update(id uint, req dto.CommissionRequest, username string) (*entity.Commission, error)
	UpdateStatus(id uint, req dto.CommissionStatusRequest, username string) (*entity.Commission, error)
	Delete(id uint) error
}

type commissionUsecase struct {
	repo repository.CommissionRepository
}

func NewCommissionUsecase(repo repository.CommissionRepository) CommissionUsecase {
	return &commissionUsecase{repo: repo}
}

func buildItems(reqItems []dto.CommissionItemRequest) ([]entity.CommissionItem, float64) {
	var items []entity.CommissionItem
	var total float64

	for _, i := range reqItems {
		items = append(items, entity.CommissionItem{
			SaintID:    i.SaintID,
			SizeID:     i.SizeID,
			MaterialID: i.MaterialID,
			StyleID:    i.StyleID,
			Price:      i.Price,
			Notes:      i.Notes,
		})
		total += i.Price
	}

	return items, total
}

func (u *commissionUsecase) GetAll() ([]entity.Commission, error) {
	return u.repo.FindAll()
}

func (u *commissionUsecase) Create(req dto.CommissionRequest, username string) (*entity.Commission, error) {
	deadline, err := time.Parse("2006-01-02", req.Deadline)
	if err != nil {
		return nil, err
	}

	items, total := buildItems(req.Items)

	newCommission := &entity.Commission{
		ClientID:   req.ClientID,
		OrderDate:  time.Now(),
		Deadline:   deadline,
		Notes:      req.Notes,
		Status:     "pending",
		TotalPrice: total,
		Items:      items,
		BaseModel: shared.BaseModel{
			Status:          "Active",
			CreatedUsername: username,
			UpdatedUsername: username,
		},
	}

	if err := u.repo.Create(newCommission); err != nil {
		return nil, err
	}

	return u.repo.FindByID(newCommission.ID)
}

func (u *commissionUsecase) Update(id uint, req dto.CommissionRequest, username string) (*entity.Commission, error) {
	existing, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	deadline, err := time.Parse("2006-01-02", req.Deadline)
	if err != nil {
		return nil, err
	}

	items, total := buildItems(req.Items)

	existing.ClientID = req.ClientID
	existing.Deadline = deadline
	existing.Notes = req.Notes
	existing.TotalPrice = total
	existing.Items = items
	existing.UpdatedUsername = username
	existing.UpdatedAt = time.Now()

	if err := u.repo.Update(existing); err != nil {
		return nil, err
	}

	return u.repo.FindByID(existing.ID)
}

func (u *commissionUsecase) UpdateStatus(id uint, req dto.CommissionStatusRequest, username string) (*entity.Commission, error) {
	existing, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	existing.Status = req.Status
	existing.UpdatedUsername = username
	existing.UpdatedAt = time.Now()

	if err := u.repo.Update(existing); err != nil {
		return nil, err
	}

	return u.repo.FindByID(existing.ID)
}

func (u *commissionUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}