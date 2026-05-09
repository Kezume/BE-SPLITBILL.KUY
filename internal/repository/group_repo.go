package repository

import (
	"context"
	"errors"

	"github.com/Kezume/BE-SPLITBILL.KUY/internal/dto"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/model"
	"github.com/Kezume/BE-SPLITBILL.KUY/pkg/database"
)

type GroupRepository interface {
	CreateGroup(ctx context.Context, group *model.Groups) error
	FetchAllGroup(ctx context.Context, userID string, pagination *dto.PaginationRequest) ([]model.Groups, int64, error)
	GetGroupDetail(ctx context.Context, groupID string) (*model.Groups, error)
	GetExpensesByGroupID(ctx context.Context, groupID string) ([]dto.ExpenseResponse, error)
	GetGroupMembers(ctx context.Context, groupID string) ([]dto.UserPreview, error)
	GetSplitsByExpenseIDs(ctx context.Context, expenseIDs []string) (map[string][]dto.SplitMemberDetail, error)
	GetUserStatsInGroup(ctx context.Context, groupID string, userID string) (yourShare float64, unpaidAmount float64, err error)
	DeleteGroup(ctx context.Context, groupID string, userID string) error
	FindGroupByInviteCode(ctx context.Context, inviteCode string) (*model.Groups, error)
	IsMemberOfGroup(ctx context.Context, groupID string, userID string) (bool, error)
	AddMemberToGroup(ctx context.Context, member *model.GroupMember) error
}

type groupRepository struct {
}

func NewGroupRepository() GroupRepository {
	return &groupRepository{}
}

func (g *groupRepository) CreateGroup(ctx context.Context, group *model.Groups) error {
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.WithContext(ctx).Create(group).Error; err != nil {
		tx.Rollback()
		return err
	}

	member := model.GroupMember{
		GroupID: group.ID,
		UserID:  group.OwnerID,
	}

	if err := tx.WithContext(ctx).Create(&member).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (g *groupRepository) FetchAllGroup(ctx context.Context, userID string, pagination *dto.PaginationRequest) ([]model.Groups, int64, error) {
	var groups []model.Groups

	var total int64

	// Tampilkan grup di mana user adalah owner ATAU member
	condition := "owner_id = ? OR id IN (SELECT group_id FROM group_members WHERE user_id = ?)"

	err := database.DB.WithContext(ctx).Model(model.Groups{}).Where(condition, userID, userID).Count(&total).Error
	if err != nil {
		return nil, 0, errors.New("Data not found")
	}

	offset := (pagination.Page - 1) * pagination.Limit

	if err := database.DB.WithContext(ctx).Where(condition, userID, userID).Offset(offset).Limit(pagination.Limit).Find(&groups).Error; err != nil {
		return nil, 0, errors.New("Data not found")
	}

	if len(groups) == 0 {
		return groups, total, nil
	}

	// untuk menghitung jumlah member
	var counts []int

	for _, group := range groups {
		var count int64
		if err := database.DB.WithContext(ctx).Model(&model.GroupMember{}).Where("group_id = ?", group.ID).Count(&count).Error; err != nil {
			return nil, 0, errors.New("Data not found")
		}
		counts = append(counts, int(count))
	}

	return groups, total, nil
}

func (g *groupRepository) GetGroupDetail(ctx context.Context, groupID string) (*model.Groups, error) {
	var group model.Groups
	if err := database.DB.WithContext(ctx).Where("id = ?", groupID).First(&group).Error; err != nil {
		return nil, errors.New("Data not found")
	}
	return &group, nil
}

func (g *groupRepository) GetExpensesByGroupID(ctx context.Context, groupID string) ([]dto.ExpenseResponse, error) {
	var scanned []dto.ExpenseScanResult
	if err := database.DB.WithContext(ctx).Table("expenses e").
		Select("e.id, DATE_FORMAT(e.date, '%Y-%m-%d') as date, e.description, e.amount, e.group_id, u.id as payer_id, u.username as payer, u.avatar_url as payer_avatar").
		Joins("JOIN users u ON e.paid_by = u.id").
		Where("e.group_id = ?", groupID).
		Order("e.date DESC").
		Scan(&scanned).Error; err != nil {
		return nil, errors.New("Data not found")
	}

	var expenses []dto.ExpenseResponse
	for _, s := range scanned {
		expenses = append(expenses, dto.ExpenseResponse{
			ID:          s.ID,
			Description: s.Description,
			Amount:      s.Amount,
			Date:        s.Date,
			GroupID:     s.GroupID,
			PaidBy: dto.UserPreview{
				ID:        s.PayerID,
				Username:  s.Payer,
				AvatarUrl: s.PayerAvatar,
			},
		})
	}

	return expenses, nil
}

func (g *groupRepository) GetGroupMembers(ctx context.Context, groupID string) ([]dto.UserPreview, error) {
	var members []dto.UserPreview
	err := database.DB.WithContext(ctx).Table("group_members gm").
		Select("u.id, u.username, u.avatar_url").
		Joins("JOIN users u ON u.id = gm.user_id").
		Where("gm.group_id = ?", groupID).
		Scan(&members).Error
	return members, err
}

func (g *groupRepository) GetSplitsByExpenseIDs(ctx context.Context, expenseIDs []string) (map[string][]dto.SplitMemberDetail, error) {
	if len(expenseIDs) == 0 {
		return map[string][]dto.SplitMemberDetail{}, nil
	}

	type scanRow struct {
		ExpenseID string  `json:"expense_id"`
		UserID    string  `json:"user_id"`
		Username  string  `json:"username"`
		AvatarUrl string  `json:"avatar_url"`
		Amount    float64 `json:"amount"`
		IsSettled bool    `json:"is_settled"`
	}

	var rows []scanRow
	err := database.DB.WithContext(ctx).Table("expense_splits es").
		Select("es.expense_id, es.user_id, u.username, u.avatar_url, es.amount, es.is_settled").
		Joins("JOIN users u ON u.id = es.user_id").
		Where("es.expense_id IN ?", expenseIDs).
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}

	result := make(map[string][]dto.SplitMemberDetail)
	for _, r := range rows {
		result[r.ExpenseID] = append(result[r.ExpenseID], dto.SplitMemberDetail{
			User: dto.UserPreview{
				ID:        r.UserID,
				Username:  r.Username,
				AvatarUrl: r.AvatarUrl,
			},
			Amount:    r.Amount,
			IsSettled: r.IsSettled,
		})
	}

	return result, nil
}

func (g *groupRepository) GetUserStatsInGroup(ctx context.Context, groupID string, userID string) (float64, float64, error) {
	// your_share: total yang harus user bayar (belum lunas)
	var yourShare float64
	database.DB.WithContext(ctx).Table("expense_splits es").
		Select("COALESCE(SUM(es.amount), 0)").
		Joins("JOIN expenses e ON e.id = es.expense_id").
		Where("e.group_id = ? AND es.user_id = ? AND es.is_settled = false", groupID, userID).
		Scan(&yourShare)

	// unpaid_amount: total semua yang belum lunas di grup
	var unpaidAmount float64
	database.DB.WithContext(ctx).Table("expense_splits es").
		Select("COALESCE(SUM(es.amount), 0)").
		Joins("JOIN expenses e ON e.id = es.expense_id").
		Where("e.group_id = ? AND es.is_settled = false", groupID).
		Scan(&unpaidAmount)

	return yourShare, unpaidAmount, nil
}

func (g *groupRepository) DeleteGroup(ctx context.Context, groupID string, userID string) error {
	result := database.DB.WithContext(ctx).Where("id = ? AND owner_id = ?", groupID, userID).Delete(model.Groups{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Group not found or you don't have permission to delete")
	}

	return nil
}

func (g *groupRepository) FindGroupByInviteCode(ctx context.Context, inviteCode string) (*model.Groups, error) {
	var group model.Groups
	if err := database.DB.WithContext(ctx).Where("invite_code = ?", inviteCode).First(&group).Error; err != nil {
		return nil, errors.New("Group not found")
	}
	return &group, nil
}

func (g *groupRepository) IsMemberOfGroup(ctx context.Context, groupID string, userID string) (bool, error) {
	var count int64
	if err := database.DB.WithContext(ctx).Model(model.GroupMember{}).Where("group_id = ? AND user_id = ?", groupID, userID).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (g *groupRepository) AddMemberToGroup(ctx context.Context, member *model.GroupMember) error {
	err := database.DB.WithContext(ctx).Create(member).Error
	if err != nil {
		return err
	}

	return nil
}
