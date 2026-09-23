package client

import (
	"time"

	dto "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/client"
	adminsistrator "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/admin"
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/client"
	repository "github.com/nathanchristiawan02/icon-commission-system-backend/internal/repository/client"
)

type ClientUsecase interface {
	GetAll() ([]entity.Client, error)
	Create(req dto.ClientRequest, username string) (*entity.Client, error)
	Update(id uint, req dto.ClientRequest, username string) (*entity.Client, error)
	Delete(id uint) error
}

type clientUsecase struct {
	repo repository.ClientRepository
}

func NewClientUsecase(repo repository.ClientRepository) ClientUsecase {
	return &clientUsecase{repo: repo}
}

func (u *clientUsecase) GetAll() ([]entity.Client, error) {
	return u.repo.FindAll()
}

func (u *clientUsecase) Create(req dto.ClientRequest, username string) (*entity.Client, error) {
	newClient := &entity.Client{
		Name:                   req.Name,
		Nickname:               req.Nickname,
		Phone:                  req.Phone,
		AlternatePhone:         req.AlternatePhone,
		Email:                  req.Email,
		InstagramHandle:        req.InstagramHandle,
		Address:                req.Address,
		City:                   req.City,
		Province:               req.Province,
		PostalCode:             req.PostalCode,
		Country:                req.Country,
		Denomination:           req.Denomination,
		PatronSaintPreference:  req.PatronSaintPreference,
		ParishName:             req.ParishName,
		PreferredPaymentMethod: req.PreferredPaymentMethod,
		Notes:                  req.Notes,
		BaseModel: adminsistrator.BaseModel{
			Status:          "Active",
			CreatedUsername: username,
			UpdatedUsername: username,
		},
	}
	err := u.repo.Create(newClient)
	return newClient, err
}

func (u *clientUsecase) Update(id uint, req dto.ClientRequest, username string) (*entity.Client, error) {
	existingClient, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	existingClient.Name = req.Name
	existingClient.Nickname = req.Nickname
	existingClient.Phone = req.Phone
	existingClient.AlternatePhone = req.AlternatePhone
	existingClient.Email = req.Email
	existingClient.InstagramHandle = req.InstagramHandle
	existingClient.Address = req.Address
	existingClient.City = req.City
	existingClient.Province = req.Province
	existingClient.PostalCode = req.PostalCode
	existingClient.Country = req.Country
	existingClient.Denomination = req.Denomination
	existingClient.PatronSaintPreference = req.PatronSaintPreference
	existingClient.ParishName = req.ParishName
	existingClient.PreferredPaymentMethod = req.PreferredPaymentMethod
	existingClient.Notes = req.Notes
	existingClient.UpdatedUsername = username
	existingClient.UpdatedAt = time.Now()

	err = u.repo.Update(existingClient)
	return existingClient, err
}

func (u *clientUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
