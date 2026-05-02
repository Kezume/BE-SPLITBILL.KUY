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

	user, err := u.service.GetProfile(userID)
	if err != nil {
		response.Error(c, 500, "Failed to Get Profile")
		return
	}

	response.Success(c, response.ToProfileResponse(user))
}

func (u *UserHandler) UpdateProfile(c *gin.Context) {
	var payload *dto.UpdateProfile
	userID := c.GetString("user_id")

	// Mencegah panic jika rute tidak sengaja terlepas dari middleware
	if userID == "" {
		response.Error(c, 401, "Unauthorized: Missing Token")
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		response.Error(c, 500, "Failed to Bind JSON")
		return
	}

	user := &model.User{
		ID:       uuid.Must(uuid.Parse(userID)),
		Username: payload.Username,
		Phone:    payload.Phone,
		Email:    payload.Email,
	}

	err := u.service.UpdateProfile(user)
	if err != nil {
		response.Error(c, 500, "Failed to Update user")
		return
	}

	response.Success(c, response.ToProfileResponse(user))
}

func (u *UserHandler) DeleteProfile(c *gin.Context) {
	userID := c.GetString("user_id")

	if userID == "" {
		response.Error(c, 401, "Unauthorized: Missing Token")
		return
	}

	err := u.service.DeleteProfile(userID)
	if err != nil {
		response.Error(c, 500, "Failed to Delete Profile")
		return
	}

	response.Success(c, "Profile deleted successfully")
}
