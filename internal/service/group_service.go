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
	GetListGroup(ctx context.Context, userID string, pagination *dto.PaginationRequest) ([]dto.ListGroupResponse, *dto.Meta, error)
	GetGroupDetail(ctx context.Context, groupID string, userID string) (*dto.GroupDetailResponse, error)
	DeleteGroup(ctx context.Context, groupID string, userID string) error
	JoinGroup(ctx context.Context, inviteCode string, userID string) (*dto.CreateGroupResponse, error)
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

func (g *groupService) GetListGroup(ctx context.Context, userID string, pagination *dto.PaginationRequest) ([]dto.ListGroupResponse, *dto.Meta, error) {
	groups, total, err := g.repo.FetchAllGroup(ctx, userID, pagination)
	if err != nil {
		return nil, nil, errors.New("Failed to fetch list group!")
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

	meta := &dto.Meta{
		Page:  pagination.Page,
		Limit: pagination.Limit,
		Total: int(total),
	}

	return data, meta, nil
}

func (g *groupService) GetGroupDetail(ctx context.Context, groupID string, userID string) (*dto.GroupDetailResponse, error) {
	group, err := g.repo.GetGroupDetail(ctx, groupID)
	if err != nil {
		return nil, err
	}

	// Ambil expenses
	expenses, err := g.repo.GetExpensesByGroupID(ctx, groupID)
	if err != nil || expenses == nil {
		expenses = []dto.ExpenseResponse{}
	}

	var totalSpent float64
	for _, exp := range expenses {
		totalSpent += exp.Amount
	}

	// Ambil members dari repo
	members, err := g.repo.GetGroupMembers(ctx, groupID)
	if err != nil || members == nil {
		members = []dto.UserPreview{}
	}

	var data dto.GroupDetailResponse
	data.ID = group.ID.String()
	data.Name = group.Name
	data.Icon = group.Icon
	data.InviteCode = group.InviteCode
	data.TotalAmount = totalSpent
	data.MemberCount = len(members)
	data.CreatedAt = group.CreatedAt

	data.Stats = dto.GroupDetailStats{
		TotalSpent:   totalSpent,
		YourShare:    0,
		UnpaidAmount: 0,
		IsSettled:    true,
	}

	data.Members = members
	data.Expenses = expenses

	return &data, nil
}

func (g *groupService) DeleteGroup(ctx context.Context, groupID string, userID string) error {
	err := g.repo.DeleteGroup(ctx, groupID, userID)
	if err != nil {
		return errors.New("Failed to delete group!")
	}

	return nil
}

func (g *groupService) JoinGroup(ctx context.Context, inviteCode string, userID string) (*dto.CreateGroupResponse, error) {
	group, err := g.repo.FindGroupByInviteCode(ctx, inviteCode)
	if err != nil {
		return nil, errors.New("Invite Code is not valid!")
	}

	isMember, err := g.repo.IsMemberOfGroup(ctx, group.ID.String(), userID)
	if err != nil {
		return nil, errors.New("Failed to check if user is member of group!")
	}

	if isMember {
		return nil, errors.New("You are already a member of this group!")
	}

	if group.OwnerID.String() == userID {
		return nil, errors.New("You are already the owner of this group!")
	}

	member := model.GroupMember{
		GroupID: group.ID,
		UserID:  uuid.Must(uuid.Parse(userID)),
	}

	if err := g.repo.AddMemberToGroup(ctx, &member); err != nil {
		return nil, errors.New("Failed to join group!")
	}

	return &dto.CreateGroupResponse{
		ID:         group.ID.String(),
		Name:       group.Name,
		Icon:       group.Icon,
		InviteCode: group.InviteCode,
		CreatedAt:  group.CreatedAt,
	}, nil
}
