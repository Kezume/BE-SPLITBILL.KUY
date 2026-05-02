package handler

import (
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/dto"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/service"
	"github.com/Kezume/BE-SPLITBILL.KUY/pkg/utils"
	"github.com/Kezume/BE-SPLITBILL.KUY/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) AuthHandler {
	return AuthHandler{
		service: service,
	}
}

func (a *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterUser

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid Register Request, try again")
		return
	}

	user, err := a.service.Register(req)
	if err != nil {
		response.Error(c, 500, "Failed to Register, Try again")
		return
	}

	token, _ := utils.GenerateToken(user.ID.String(), user.Email)

	response.Success(c, response.ToUserResponse(user, token))
}

func (a *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginUser

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid Login Request, Try again")
		return
	}

	user, token, err := a.service.Login(req)
	if err != nil {
		response.Error(c, 500, "Failed to Login, Try Again")
		return
	}

	response.Success(c, response.ToUserResponse(user, token))
}

func (a *AuthHandler) Logout(c *gin.Context) {
	response.Success(c, gin.H{
		"message": "Logout Successfully",
	})
}
