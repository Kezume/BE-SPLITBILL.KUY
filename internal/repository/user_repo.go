package repository

import (
	"context"

	"github.com/Kezume/BE-SPLITBILL.KUY/internal/model"
	"github.com/Kezume/BE-SPLITBILL.KUY/pkg/database"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	FindUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetByID(ctx context.Context, id string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id string) error
}

type userRepo struct {
}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (u *userRepo) CreateUser(ctx context.Context, user *model.User) error {
	err := database.DB.WithContext(ctx).Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := database.DB.WithContext(ctx).Where("email = ? ", email).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) GetByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := database.DB.WithContext(ctx).Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) UpdateUser(ctx context.Context, user *model.User) error {
	err := database.DB.WithContext(ctx).Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) DeleteUser(ctx context.Context, id string) error {
	var user model.User
	err := database.DB.WithContext(ctx).Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
