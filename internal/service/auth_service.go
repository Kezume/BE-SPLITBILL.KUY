package service

import (
	"errors"

	"github.com/Kezume/BE-SPLITBILL.KUY/internal/dto"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/model"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/repository"
	"github.com/Kezume/BE-SPLITBILL.KUY/pkg/utils"
)

type AuthService interface {
	Register(req dto.RegisterUser) (*model.User, error)
	Login(req dto.LoginUser) (*model.User, string, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *authService {
	return &authService{
		repo: repo,
	}
}

func (u *authService) Register(req dto.RegisterUser) (*model.User, error) {
	hashPassword, _ := utils.HashPassword(req.Password)

	user := model.User{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: hashPassword,
	}

	err := u.repo.CreateUser(&user)

	return &user, err
}

func (u *authService) Login(req dto.LoginUser) (*model.User, string, error) {
	user, err := u.repo.FindUserByEmail(req.Email)
	if err != nil {
		return nil, "", errors.New("User Not Found")
	}

	if err := utils.VerifyPassword(req.Password, user.Password); err != nil {
		return nil, "", errors.New("Invalid Credentials")
	}

	token, err := utils.GenerateToken(user.ID.String(), user.Email)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
