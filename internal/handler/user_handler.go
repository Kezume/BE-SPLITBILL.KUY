package handler

import (
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/dto"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/model"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/service"
	"github.com/Kezume/BE-SPLITBILL.KUY/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return UserHandler{
		service: service,
	}
}

func (u *UserHandler) GetProfile(c *gin.Context) {
	userID := c.GetString("user_id")

	user, err := u.service.GetProfile(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, 500, "Gagal ambil data profil lo!")
		return
	}

	response.Success(c, response.ToProfileResponse(user))
}

func (u *UserHandler) UpdateProfile(c *gin.Context) {
	var payload *dto.UpdateProfile
	userID := c.GetString("user_id")

	// Mencegah panic jika rute tidak sengaja terlepas dari middleware
	if userID == "" {
		response.Error(c, 401, "Lo belum login nih, login dulu!")
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		response.Error(c, 500, "Ada yang salah nih, coba lagi ya!")
		return
	}

	user := &model.User{
		ID:       uuid.Must(uuid.Parse(userID)),
		Username: payload.Username,
		Phone:    payload.Phone,
		Email:    payload.Email,
	}

	err := u.service.UpdateProfile(c.Request.Context(), user)
	if err != nil {
		response.Error(c, 500, "Gagal update profil lo!")
		return
	}

	response.Success(c, response.ToProfileResponse(user))
}

func (u *UserHandler) DeleteProfile(c *gin.Context) {
	userID := c.GetString("user_id")

	if userID == "" {
		response.Error(c, 401, "Lo belum login nih, login dulu!")
		return
	}

	err := u.service.DeleteProfile(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, 500, "Gagal hapus akun lo!")
		return
	}

	response.Success(c, "Profile deleted successfully")
}
