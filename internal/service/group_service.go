package service

import (
	"context"
	"errors"

	"github.com/Kezume/BE-SPLITBILL.KUY/internal/dto"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/model"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/repository"
	"github.com/Kezume/BE-SPLITBILL.KUY/pkg/database"
	"github.com/Kezume/BE-SPLITBILL.KUY/pkg/utils"
	"github.com/google/uuid"
)

type GroupService interface {
	CreateGroup(ctx context.Context, group *model.Groups, userID string) error
	GetListGroup(ctx context.Context, userID string, pagination *dto.PaginationRequest) ([]dto.ListGroupResponse, error)
}

type groupService struct {
	repo repository.GroupRepository
}

func NewGroupService(repo repository.GroupRepository) GroupService {
	return &groupService{
		repo: repo,
	}
}

func (g *groupService) CreateGroup(ctx context.Context, group *model.Groups, userID string) error {
	group.OwnerID = uuid.Must(uuid.Parse(userID))
	group.InviteCode = utils.GenerateRandomInviteCode()
	err := g.repo.CreateGroup(ctx, group)
	if err != nil {
		return errors.New("Failed to create group!")
	}

	return nil
}

func (g *groupService) GetListGroup(ctx context.Context, userID string, pagination *dto.PaginationRequest) ([]dto.ListGroupResponse, error) {
	groups, err := g.repo.FetchAllGroup(ctx, userID, pagination)
	if err != nil {
		return nil, errors.New("Failed to fetch list group!")
	}

	var data []dto.ListGroupResponse

	for _, group := range groups {
		var memberCount int64
		var totalAmount float64

		// 1. Menghitung jumlah anggota (Member Count) menggunakan package database
		database.DB.WithContext(ctx).Table("group_members").Where("group_id = ?", group.ID).Count(&memberCount)

		// 2. Menghitung total tagihan (Total Amount) menggunakan package database
		database.DB.WithContext(ctx).Table("expenses").Where("group_id = ?", group.ID).Select("COALESCE(SUM(amount), 0)").Scan(&totalAmount)

		data = append(data, dto.ListGroupResponse{
			ID:          group.ID.String(),
			Name:        group.Name,
			Icon:        group.Icon,
			InviteCode:  group.InviteCode,
			TotalAmount: totalAmount,
			MemberCount: int(memberCount),
			CreatedAt:   group.CreatedAt,
		})
	}

	return data, nil
}
