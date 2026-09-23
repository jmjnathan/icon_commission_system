package master

import (
	"time"

	dto "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/master"
	adminsistrator "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/admin"
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/master"
	repository "github.com/nathanchristiawan02/icon-commission-system-backend/internal/repository/master"
)

type MasterSizeUsecase interface {
	GetAll() ([]entity.MasterSize, error)
	Create(req dto.MasterSizeRequest, username string) (*entity.MasterSize, error)
	Update(id uint, req dto.MasterSizeRequest, username string) (*entity.MasterSize, error)
	Delete(id uint) error
}

type masterSizeUsecase struct {
	repo repository.MasterSizeRepository
}

func NewMasterSizeUsecase(repo repository.MasterSizeRepository) MasterSizeUsecase {
	return &masterSizeUsecase{repo: repo}
}

func (u *masterSizeUsecase) GetAll() ([]entity.MasterSize, error) {
	return u.repo.FindAll()
}

func (u *masterSizeUsecase) Create(req dto.MasterSizeRequest, username string) (*entity.MasterSize, error) {
	size := &entity.MasterSize{
		Name: req.Name,
		Size: req.Size,
		BaseModel: adminsistrator.BaseModel{
			Status:          "active",
			CreatedUsername: username,
			UpdatedUsername: username,
		},
	}
	err := u.repo.Create(size)
	return size, err
}

func (u *masterSizeUsecase) Update(id uint, req dto.MasterSizeRequest, username string) (*entity.MasterSize, error) {
	size, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	size.Name = req.Name
	size.Size = req.Size
	size.Status = req.Status
	size.UpdatedUsername = username
	size.UpdatedAt = time.Now()

	err = u.repo.Update(size)
	return size, err
}

func (u *masterSizeUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
