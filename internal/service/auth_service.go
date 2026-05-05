package service

import (
	"context"
	"errors"

	"github.com/Kezume/BE-SPLITBILL.KUY/internal/dto"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/model"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/repository"
	"github.com/Kezume/BE-SPLITBILL.KUY/pkg/utils"
)

type AuthService interface {
	Register(ctx context.Context, req dto.RegisterUser) (*model.User, error)
	Login(ctx context.Context, req dto.LoginUser) (*model.User, string, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (u *authService) Register(ctx context.Context, req dto.RegisterUser) (*model.User, error) {
	hashPassword, _ := utils.HashPassword(req.Password)

	user := model.User{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: hashPassword,
	}

	err := u.repo.CreateUser(ctx, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *authService) Login(ctx context.Context, req dto.LoginUser) (*model.User, string, error) {
	user, err := u.repo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, "", errors.New("User gak ketemu!")
	}

	if err := utils.VerifyPassword(req.Password, user.Password); err != nil {
		return nil, "", errors.New("Email atau password lo salah!")
	}

	token, err := utils.GenerateToken(user.ID.String(), user.Email)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
