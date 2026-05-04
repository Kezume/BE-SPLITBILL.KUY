package response

import (
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/dto"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/model"
)

func ToGroupResponse(group *model.Groups) dto.CreateGroupResponse {
	return dto.CreateGroupResponse{
		ID:          group.ID.String(),
		Name:        group.Name,
		Icon:        group.Icon,
		InviteCode:  group.InviteCode,
		TotalAmount: 0,
		MemberCount: 1,
		CreatedAt:   group.CreatedAt,
	}
}

// func ToGroupListResponse(groups []model.Groups) []dto.ListGroupResponse {
// 	var data []dto.ListGroupResponse

// 	for _, group := range groups {
// 		data = append(data, dto.ListGroupResponse{
// 			ID:          group.ID.String(),
// 			Name:        group.Name,
// 			Icon:        group.Icon,
// 			InviteCode:  group.InviteCode,
// 			TotalAmount: 0,
// 			MemberCount: 1,
// 			CreatedAt:   group.CreatedAt,
// 		})
// 	}

// 	return data
// }
