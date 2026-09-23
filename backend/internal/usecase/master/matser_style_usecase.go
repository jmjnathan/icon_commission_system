package master

import (
	"time"

	dto "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/master"
	adminsistrator "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/admin"
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/master"
	repository "github.com/nathanchristiawan02/icon-commission-system-backend/internal/repository/master"
)

type MasterStyleUsecase interface {
	GetAll() ([]entity.MasterStyle, error)
	Create(req dto.MasterStyleRequest, username string) (*entity.MasterStyle, error)
	Update(id uint, req dto.MasterStyleRequest, username string) (*entity.MasterStyle, error)
	Delete(id uint) error
}

type masterStyleUsecase struct {
	repo repository.MasterStyleRepository
}

func NewMasterStyleUsecase(repo repository.MasterStyleRepository) MasterStyleUsecase {
	return &masterStyleUsecase{repo: repo}
}

func (u *masterStyleUsecase) GetAll() ([]entity.MasterStyle, error) {
	return u.repo.FindAll()
}

func (u *masterStyleUsecase) Create(req dto.MasterStyleRequest, username string) (*entity.MasterStyle, error) {
	material := &entity.MasterStyle{
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

func (u *masterStyleUsecase) Update(id uint, req dto.MasterStyleRequest, username string) (*entity.MasterStyle, error) {
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

func (u *masterStyleUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
