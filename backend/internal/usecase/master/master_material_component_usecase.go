package master

import (
	"time"

	dto "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/master"
	shared "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/admin"
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/master"
	repository "github.com/nathanchristiawan02/icon-commission-system-backend/internal/repository/master"
)

type MaterialComponentUsecase interface {
	GetAll() ([]entity.MaterialComponent, error)
	Create(req dto.MaterialComponentRequest, username string) (*entity.MaterialComponent, error)
	Update(id uint, req dto.MaterialComponentRequest, username string) (*entity.MaterialComponent, error)
	Delete(id uint) error
}

type materialComponentUsecase struct {
	repo repository.MaterialComponentRepository
}

func NewMaterialComponentUsecase(
	repo repository.MaterialComponentRepository,
) MaterialComponentUsecase {
	return &materialComponentUsecase{repo: repo}
}

func (u *materialComponentUsecase) GetAll() ([]entity.MaterialComponent, error) {
	return u.repo.FindAll()
}

func (u *materialComponentUsecase) Create(
	req dto.MaterialComponentRequest,
	username string,
) (*entity.MaterialComponent, error) {

	component := &entity.MaterialComponent{
		Name:     req.Name,
		Category: req.Category,
		Remark:   req.Remark,
		BaseModel: shared.BaseModel{
			Status:          "Active",
			CreatedUsername: username,
			UpdatedUsername: username,
		},
	}

	err := u.repo.Create(component)

	return component, err
}

func (u *materialComponentUsecase) Update(
	id uint,
	req dto.MaterialComponentRequest,
	username string,
) (*entity.MaterialComponent, error) {

	component, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	component.Name = req.Name
	component.Category = req.Category
	component.Remark = req.Remark
	component.Status = req.Status
	component.UpdatedUsername = username
	component.UpdatedAt = time.Now()

	err = u.repo.Update(component)

	return component, err
}

func (u *materialComponentUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
