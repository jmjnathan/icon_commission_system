package master

import (
	"time"

	dto "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/master"
	adminsistrator "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/admin"
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/master"
	repository "github.com/nathanchristiawan02/icon-commission-system-backend/internal/repository/master"
)

type MasterMaterialUsecase interface {
	GetAll() ([]entity.MasterMaterial, error)
	Create(req dto.MasterMaterialRequest, username string) (*entity.MasterMaterial, error)
	Update(id uint, req dto.MasterMaterialRequest, username string) (*entity.MasterMaterial, error)
	Delete(id uint) error
}

type masterMaterialUsecase struct {
	repo repository.MasterMaterialRepository
}

func NewMasterMaterialUsecase(repo repository.MasterMaterialRepository) MasterMaterialUsecase {
	return &masterMaterialUsecase{repo: repo}
}

func (u *masterMaterialUsecase) GetAll() ([]entity.MasterMaterial, error) {
	return u.repo.FindAll()
}

func (u *masterMaterialUsecase) Create(req dto.MasterMaterialRequest, username string) (*entity.MasterMaterial, error) {
	material := &entity.MasterMaterial{
		Name:   req.Name,
		Remark: req.Remark,
		BaseModel: adminsistrator.BaseModel{
			Status:          "active",
			CreatedUsername: username,
			UpdatedUsername: username,
		},
	}
	err := u.repo.Create(material)
	return material, err
}

func (u *masterMaterialUsecase) Update(id uint, req dto.MasterMaterialRequest, username string) (*entity.MasterMaterial, error) {
	material, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	material.Name = req.Name
	material.Remark = req.Remark
	material.Status = req.Status
	material.UpdatedUsername = username
	material.UpdatedAt = time.Now()

	err = u.repo.Update(material)
	return material, err
}

func (u *masterMaterialUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
