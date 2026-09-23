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
	AddPayment(commissionID uint, req dto.CommissionPaymentRequest, username string) (*entity.Commission, error)
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

	items, subtotal := buildItems(req.Items)

	var discountAmount float64
	if req.DiscountType == "percent" {
		discountAmount = subtotal * (req.DiscountValue / 100)
	} else if req.DiscountType == "fixed" {
		discountAmount = req.DiscountValue
	}
	totalPrice := subtotal - discountAmount
	if totalPrice < 0 {
		totalPrice = 0
	}

	var photos []entity.CommissionPhoto
	for _, url := range req.PhotoUrls {
		photos = append(photos, entity.CommissionPhoto{FileURL: url})
	}

	newCommission := &entity.Commission{
		ClientID:      req.ClientID,
		OrderDate:     time.Now(),
		Deadline:      deadline,
		Notes:         req.Notes,
		Status:        "pending",
		Subtotal:      subtotal,
		DiscountType:  req.DiscountType,
		DiscountValue: req.DiscountValue,
		TotalPrice:    totalPrice,
		PaymentStatus: "unpaid",
		Items:         items,
		Photos:        photos,
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

func (u *commissionUsecase) AddPayment(commissionID uint, req dto.CommissionPaymentRequest, username string) (*entity.Commission, error) {
	commission, err := u.repo.FindByID(commissionID)
	if err != nil {
		return nil, err
	}

	payment := entity.CommissionPayment{
		CommissionID:    commissionID,
		Amount:          req.Amount,
		PaymentType:     req.PaymentType,
		Method:          req.Method,
		PaidAt:          time.Now(),
		Notes:           req.Notes,
		CreatedUsername: username,
	}

	if err := u.repo.CreatePayment(&payment); err != nil {
		return nil, err
	}

	// Hitung total sudah dibayar, update payment_status
	totalPaid := 0.0
	for _, p := range commission.Payments {
		totalPaid += p.Amount
	}
	totalPaid += req.Amount 

	if totalPaid >= commission.TotalPrice {
		commission.PaymentStatus = "paid"
	} else if totalPaid > 0 {
		commission.PaymentStatus = "partial"
	}
	u.repo.Update(commission)

	return u.repo.FindByID(commissionID)
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
