package repository

import (
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/model"
	"github.com/Kezume/BE-SPLITBILL.KUY/pkg/database"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	FindUserByEmail(email string) (*model.User, error)
	GetByID(id string) (*model.User, error)
}

type userRepo struct {
}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (u *userRepo) CreateUser(user *model.User) error {
	return database.DB.Create(user).Error
}

func (u *userRepo) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := database.DB.Where("email = ? ", email).Find(&user).Error

	return &user, err
}

func (u *userRepo) GetByID(id string) (*model.User, error) {
	var user model.User
	err := database.DB.Where("id = ?", id).Find(&user).Error

	return &user, err
}
