package master

import (
	"time"

	dto "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/master"
	admin "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/admin"
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/master"
	repository "github.com/nathanchristiawan02/icon-commission-system-backend/internal/repository/master"
)

type MaterialComponentVariantUsecase interface {
	GetAll() ([]entity.MaterialComponentVariant, error)

	GetByMaterialComponentID(
		materialComponentID uint,
	) ([]entity.MaterialComponentVariant, error)

	Create(
		materialComponentID uint,
		req dto.MaterialComponentVariantRequest,
		username string,
	) (*entity.MaterialComponentVariant, error)

	Update(
		id uint,
		req dto.MaterialComponentVariantRequest,
		username string,
	) (*entity.MaterialComponentVariant, error)

	Delete(id uint) error
}

type materialComponentVariantUsecase struct {
	repo repository.MaterialComponentVariantRepository
}

func NewMaterialComponentVariantUsecase(
	repo repository.MaterialComponentVariantRepository,
) MaterialComponentVariantUsecase {
	return &materialComponentVariantUsecase{
		repo: repo,
	}
}

func (u *materialComponentVariantUsecase) GetAll() (
	[]entity.MaterialComponentVariant,
	error,
) {
	return u.repo.FindAll()
}

func (u *materialComponentVariantUsecase) GetByMaterialComponentID(
	materialComponentID uint,
) ([]entity.MaterialComponentVariant, error) {
	return u.repo.FindByMaterialComponentID(
		materialComponentID,
	)
}

func (u *materialComponentVariantUsecase) Create(
	materialComponentID uint,
	req dto.MaterialComponentVariantRequest,
	username string,
) (*entity.MaterialComponentVariant, error) {

	variant := &entity.MaterialComponentVariant{
		MaterialComponentID: materialComponentID,
		Specification:       req.Specification,
		Length:              req.Length,
		Width:               req.Width,
		Thickness:           req.Thickness,
		Unit:                req.Unit,
		UnitPrice:           req.UnitPrice,
		Remark:              req.Remark,

		BaseModel: admin.BaseModel{
			Status:          "Active",
			CreatedUsername: username,
			UpdatedUsername: username,
		},
	}

	err := u.repo.Create(variant)
	if err != nil {
		return nil, err
	}

	return u.repo.FindByID(variant.ID)
}

func (u *materialComponentVariantUsecase) Update(
	id uint,
	req dto.MaterialComponentVariantRequest,
	username string,
) (*entity.MaterialComponentVariant, error) {

	variant, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	variant.Specification = req.Specification
	variant.Length = req.Length
	variant.Width = req.Width
	variant.Thickness = req.Thickness
	variant.Unit = req.Unit
	variant.UnitPrice = req.UnitPrice
	variant.Remark = req.Remark
	variant.Status = req.Status

	variant.UpdatedUsername = username
	variant.UpdatedAt = time.Now()

	err = u.repo.Update(variant)
	if err != nil {
		return nil, err
	}

	return u.repo.FindByID(variant.ID)
}

func (u *materialComponentVariantUsecase) Delete(
	id uint,
) error {
	return u.repo.Delete(id)
}
