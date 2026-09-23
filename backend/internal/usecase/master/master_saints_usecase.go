package master

import (
	"time"

	dto "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/master"
	adminsistrator "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/admin"
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/master"
	repository "github.com/nathanchristiawan02/icon-commission-system-backend/internal/repository/master"
)

type MasterSaintsUsecase interface {
	GetAll() ([]entity.MasterSaints, error)
	Create(req dto.MasterSaintsRequest, username string) (*entity.MasterSaints, error)
	Update(id uint, req dto.MasterSaintsRequest, username string) (*entity.MasterSaints, error)
	Delete(id uint) error
}

type masterSaintsUsecase struct {
	repo repository.MasterSaintsRepository
}

func NewMasterSaintsUsecase(repo repository.MasterSaintsRepository) MasterSaintsUsecase {
	return &masterSaintsUsecase{repo: repo}
}

func (u *masterSaintsUsecase) GetAll() ([]entity.MasterSaints, error) {
	return u.repo.FindAll()
}

func (u *masterSaintsUsecase) Create(req dto.MasterSaintsRequest, username string) (*entity.MasterSaints, error) {
	saint := &entity.MasterSaints{
		Name:   req.Name,
		Remark: req.Remark,
		BaseModel: adminsistrator.BaseModel{
			Status:          "active",
			CreatedUsername: username,
			UpdatedUsername: username,
		},
	}
	err := u.repo.Create(saint)
	return saint, err
}

func (u *masterSaintsUsecase) Update(id uint, req dto.MasterSaintsRequest, username string) (*entity.MasterSaints, error) {
	saint, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	saint.Name = req.Name
	saint.Remark = req.Remark
	saint.Status = req.Status
	saint.UpdatedUsername = username
	saint.UpdatedAt = time.Now()

	err = u.repo.Update(saint)
	return saint, err
}

func (u *masterSaintsUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
