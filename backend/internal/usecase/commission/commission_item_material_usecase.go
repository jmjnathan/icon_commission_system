package commission

import (
	"errors"
	"time"

	dto "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/commission"
	shared "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/admin"
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/commission"
	commissionRepo "github.com/nathanchristiawan02/icon-commission-system-backend/internal/repository/commission"
	masterRepo "github.com/nathanchristiawan02/icon-commission-system-backend/internal/repository/master"
)

type CommissionItemMaterialUsecase interface {
	GetAll(commissionItemID uint) ([]entity.CommissionItemMaterial, error)

	Create(
		commissionItemID uint,
		req dto.CommissionItemMaterialRequest,
		username string,
	) (*entity.CommissionItemMaterial, error)

	Update(
		id uint,
		req dto.CommissionItemMaterialRequest,
		username string,
	) (*entity.CommissionItemMaterial, error)

	Delete(id uint) error
}
type MaterialComponentVariantUsecase interface {
	GetAll(commissionItemID uint) ([]entity.CommissionItemMaterial, error)

	Create(
		commissionItemID uint,
		req dto.CommissionItemMaterialRequest,
		username string,
	) (*entity.CommissionItemMaterial, error)

	Update(
		id uint,
		req dto.CommissionItemMaterialRequest,
		username string,
	) (*entity.CommissionItemMaterial, error)

	Delete(id uint) error
}

type commissionItemMaterialUsecase struct {
	repo        commissionRepo.CommissionItemMaterialRepository
	variantRepo masterRepo.MaterialComponentVariantRepository
}

func NewCommissionItemMaterialUsecase(
	repo commissionRepo.CommissionItemMaterialRepository,
	variantRepo masterRepo.MaterialComponentVariantRepository,
) CommissionItemMaterialUsecase {
	return &commissionItemMaterialUsecase{
		repo:        repo,
		variantRepo: variantRepo,
	}
}

func (u *commissionItemMaterialUsecase) GetAll(
	commissionItemID uint,
) ([]entity.CommissionItemMaterial, error) {

	return u.repo.FindAllByCommissionItemID(
		commissionItemID,
	)
}

func (u *commissionItemMaterialUsecase) Create(
	commissionItemID uint,
	req dto.CommissionItemMaterialRequest,
	username string,
) (*entity.CommissionItemMaterial, error) {

	// 1. Pastikan variant memang ada
	variant, err := u.variantRepo.FindByID(
		req.MaterialComponentVariantID,
	)

	if err != nil {
		return nil, errors.New(
			"material component variant tidak ditemukan",
		)
	}

	// 2. Ambil harga dari master variant
	unitPrice := variant.UnitPrice

	// 3. Hitung total cost di backend
	totalCost := req.Quantity * unitPrice

	// 4. Simpan snapshot harga saat commission dibuat
	material := &entity.CommissionItemMaterial{
		CommissionItemID:           commissionItemID,
		MaterialComponentVariantID: req.MaterialComponentVariantID,
		Quantity:                   req.Quantity,
		UnitPrice:                  unitPrice,
		TotalCost:                  totalCost,
		Remark:                     req.Remark,

		BaseModel: shared.BaseModel{
			Status:          "active",
			CreatedUsername: username,
			UpdatedUsername: username,
		},
	}

	if err := u.repo.Create(material); err != nil {
		return nil, err
	}

	return u.repo.FindByID(material.ID)
}

func (u *commissionItemMaterialUsecase) Update(
	id uint,
	req dto.CommissionItemMaterialRequest,
	username string,
) (*entity.CommissionItemMaterial, error) {

	// Ambil record existing
	existing, err := u.repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	// Cari variant baru
	variant, err := u.variantRepo.FindByID(
		req.MaterialComponentVariantID,
	)

	if err != nil {
		return nil, errors.New(
			"material component variant tidak ditemukan",
		)
	}
	unitPrice := variant.UnitPrice
	totalCost := req.Quantity * unitPrice

	existing.MaterialComponentVariantID =
		req.MaterialComponentVariantID

	existing.Quantity = req.Quantity
	existing.UnitPrice = unitPrice
	existing.TotalCost = totalCost
	existing.Remark = req.Remark

	existing.UpdatedUsername = username
	existing.UpdatedAt = time.Now()

	if err := u.repo.Update(existing); err != nil {
		return nil, err
	}

	return u.repo.FindByID(existing.ID)
}

func (u *commissionItemMaterialUsecase) Delete(
	id uint,
) error {
	return u.repo.Delete(id)
}
