package response

import (
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/dto"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/model"
)

func ToUserResponse(user *model.User, token string) dto.AuthResponse {
	return dto.AuthResponse{
		ID:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}
}
