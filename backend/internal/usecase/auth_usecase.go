package usecase

import (
	"errors"

	"github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto"
	"github.com/nathanchristiawan02/icon-commission-system-backend/internal/repository"
	"github.com/nathanchristiawan02/icon-commission-system-backend/pkg"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Login(req dto.LoginRequest) (*dto.LoginResponse, error)
}

type authUsecase struct {
	adminRepo repository.AdminRepository
}

func NewAuthUsecase(adminRepo repository.AdminRepository) AuthUsecase {
	return &authUsecase{adminRepo: adminRepo}
}

func (u *authUsecase) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	admin, err := u.adminRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, errors.New("username atau password salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("username atau password salah")
	}

	token, err := pkg.GenerateToken(admin.ID, admin.Username)
	if err != nil {
		return nil, errors.New("gagal generate token")
	}

	return &dto.LoginResponse{
		Token:    token,
		Username: admin.Username,
	}, nil
}