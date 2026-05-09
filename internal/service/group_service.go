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
		return errors.New("Gagal bikin grup!")
	}

	return nil
}

func (g *groupService) GetListGroup(ctx context.Context, userID string, pagination *dto.PaginationRequest) ([]dto.ListGroupResponse, *dto.Meta, error) {
	groups, total, err := g.repo.FetchAllGroup(ctx, userID, pagination)
	if err != nil {
		return nil, nil, errors.New("Gagal ambil daftar grup!")
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

	// Ambil split details untuk semua expense sekaligus
	var expenseIDs []string
	for _, exp := range expenses {
		expenseIDs = append(expenseIDs, exp.ID)
	}

	splitsMap, _ := g.repo.GetSplitsByExpenseIDs(ctx, expenseIDs)
	for i, exp := range expenses {
		if splits, ok := splitsMap[exp.ID]; ok {
			expenses[i].SplitWith = splits
			if len(splits) > 0 {
				expenses[i].PerPerson = splits[0].Amount
			}
		} else {
			expenses[i].SplitWith = []dto.SplitMemberDetail{}
		}
	}

	// Ambil members dari repo
	members, err := g.repo.GetGroupMembers(ctx, groupID)
	if err != nil || members == nil {
		members = []dto.UserPreview{}
	}

	// Hitung stats dari data asli
	yourShare, unpaidAmount, _ := g.repo.GetUserStatsInGroup(ctx, groupID, userID)

	var data dto.GroupDetailResponse
	data.ID = group.ID.String()
	data.Name = group.Name
	data.Icon = group.Icon
	data.InviteCode = group.InviteCode
	data.TotalAmount = totalSpent
	data.MemberCount = len(members)
	data.IsOwner = group.OwnerID.String() == userID
	data.CreatedAt = group.CreatedAt

	data.Stats = dto.GroupDetailStats{
		TotalSpent:   totalSpent,
		YourShare:    yourShare,
		UnpaidAmount: unpaidAmount,
		IsSettled:    unpaidAmount == 0,
	}

	data.Members = members
	data.Expenses = expenses

	return &data, nil
}

func (g *groupService) DeleteGroup(ctx context.Context, groupID string, userID string) error {
	err := g.repo.DeleteGroup(ctx, groupID, userID)
	if err != nil {
		return errors.New("Gagal hapus grup!")
	}

	return nil
}

func (g *groupService) JoinGroup(ctx context.Context, inviteCode string, userID string) (*dto.CreateGroupResponse, error) {
	group, err := g.repo.FindGroupByInviteCode(ctx, inviteCode)
	if err != nil {
		return nil, errors.New("Invite code gak valid nih!")
	}

	isMember, err := g.repo.IsMemberOfGroup(ctx, group.ID.String(), userID)
	if err != nil {
		return nil, errors.New("Gagal cek keanggotaan lo!")
	}

	if isMember {
		return nil, errors.New("Lo udah gabung di grup ini, bro!")
	}

	if group.OwnerID.String() == userID {
		return nil, errors.New("Lo udah jadi owner grup ini!")
	}

	member := model.GroupMember{
		GroupID: group.ID,
		UserID:  uuid.Must(uuid.Parse(userID)),
	}

	if err := g.repo.AddMemberToGroup(ctx, &member); err != nil {
		return nil, errors.New("Gagal gabung ke grup!")
	}

	return &dto.CreateGroupResponse{
		ID:         group.ID.String(),
		Name:       group.Name,
		Icon:       group.Icon,
		InviteCode: group.InviteCode,
		CreatedAt:  group.CreatedAt,
	}, nil
}
