package service

import (
	"context"
	"errors"

	"github.com/Kezume/BE-SPLITBILL.KUY/internal/model"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/repository"
)

type UserService interface {
	GetProfile(ctx context.Context, id string) (*model.User, error)
	UpdateProfile(ctx context.Context, user *model.User) error
	DeleteProfile(ctx context.Context, id string) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (u *userService) GetProfile(ctx context.Context, id string) (*model.User, error) {
	user, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("User gak ketemu!")
	}

	return user, nil
}

func (u *userService) UpdateProfile(ctx context.Context, inputUser *model.User) error {
	// Ambil data lama dari DB (gunakan nama variabel berbeda)
	existingUser, err := u.repo.GetByID(ctx, inputUser.ID.String())
	if err != nil {
		return errors.New("User gak ketemu!")
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
	err = u.repo.UpdateUser(ctx, existingUser)
	if err != nil {
		return errors.New("Gagal update profil lo!")
	}

	// Timpa pointer asli agar handler mendapatkan data yang sudah di-update
	*inputUser = *existingUser

	return nil
}

func (u *userService) DeleteProfile(ctx context.Context, id string) error {
	err := u.repo.DeleteUser(ctx, id)
	if err != nil {
		return errors.New("Gagal hapus akun lo!")
	}

	return nil
}
