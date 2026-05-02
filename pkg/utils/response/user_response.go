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

func ToProfileResponse(user *model.User) dto.GetProfileResponse {
	avatarUrl := "https://api.dicebear.com/9.x/notionists/svg?seed=" + user.Username + "&backgroundColor=ffd700"
	if user.AvatarUrl != nil {
		avatarUrl = *user.AvatarUrl
	}

	return dto.GetProfileResponse{
		ID:          user.ID.String(),
		Username:    user.Username,
		Email:       user.Email,
		Phone:       user.Phone,
		AvatarUrl:   avatarUrl,
		Badge:       user.Badge,
		MemberSince: user.CreatedAt,
		Stats: dto.UserStats{
			TotalSplit:  0,
			ActiveGroup: 0,
			TotalFriend: 0,
			TotalDrama:  0,
		},
	}
}
