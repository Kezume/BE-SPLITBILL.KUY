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
		response.Error(c, 400, "Data registrasi lo salah nih, coba lagi!")
		return
	}

	user, err := a.service.Register(c.Request.Context(), req)
	if err != nil {
		response.Error(c, 500, "Gagal daftar nih, coba lagi ya!")
		return
	}

	token, _ := utils.GenerateToken(user.ID.String(), user.Email)

	response.Success(c, response.ToUserResponse(user, token))
}

func (a *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginUser

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Data login lo salah, coba lagi!")
		return
	}

	user, token, err := a.service.Login(c.Request.Context(), req)
	if err != nil {
		response.Error(c, 500, "Gagal login nih, coba lagi!")
		return
	}

	response.Success(c, response.ToUserResponse(user, token))
}

func (a *AuthHandler) Logout(c *gin.Context) {
	response.Success(c, gin.H{
		"message": "Logout Successfully",
	})
}
