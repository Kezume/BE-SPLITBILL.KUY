package service

import (
	"errors"

	"github.com/Kezume/BE-SPLITBILL.KUY/internal/model"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/repository"
)

type UserService interface {
	GetProfile(id string) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{
		repo: repo,
	}
}

func (u *userService) GetProfile(id string) (*model.User, error) {
	user, err := u.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("User Not Found")
	}

	return user, nil
}
