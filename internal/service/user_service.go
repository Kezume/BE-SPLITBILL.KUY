package service

import (
	"errors"

	"github.com/Kezume/BE-SPLITBILL.KUY/internal/model"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/repository"
)

type UserService interface {
	GetProfile(id string) (*model.User, error)
	UpdateProfile(user *model.User) error
	DeleteProfile(id string) error
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

func (u *userService) UpdateProfile(inputUser *model.User) error {
	// Ambil data lama dari DB (gunakan nama variabel berbeda)
	existingUser, err := u.repo.GetByID(inputUser.ID.String())
	if err != nil {
		return errors.New("User not found!")
	}

	// Lakukan pengecekan partial update berdasarkan inputUser
	if inputUser.Email != "" {
		existingUser.Email = inputUser.Email
	}
	if inputUser.Phone != "" {
		existingUser.Phone = inputUser.Phone
	}
	if inputUser.Username != "" {
		existingUser.Username = inputUser.Username
	}

	// Simpan perubahan ke DB
	err = u.repo.UpdateUser(existingUser)
	if err != nil {
		return errors.New("Failed to update user!")
	}

	// Timpa pointer asli agar handler mendapatkan data yang sudah di-update
	*inputUser = *existingUser

	return nil
}

func (u *userService) DeleteProfile(id string) error {
	err := u.repo.DeleteUser(id)
	if err != nil {
		return errors.New("Failed to delete user!")
	}

	return nil
}
