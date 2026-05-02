package handler

import (
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/service"
	"github.com/Kezume/BE-SPLITBILL.KUY/pkg/utils/response"
	"github.com/gin-gonic/gin"
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
