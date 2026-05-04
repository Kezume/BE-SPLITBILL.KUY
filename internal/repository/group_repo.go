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
	FetchAllGroup(ctx context.Context, userID string, pagination *dto.PaginationRequest) ([]model.Groups, error)
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

func (g *groupRepository) FetchAllGroup(ctx context.Context, userID string, pagination *dto.PaginationRequest) ([]model.Groups, error) {
	var groups []model.Groups

	offset := (pagination.Page - 1) * pagination.Limit

	if err := database.DB.WithContext(ctx).Where("owner_id = ?", userID).Offset(offset).Limit(pagination.Limit).Find(&groups).Error; err != nil {
		return nil, errors.New("Data not found")
	}

	if len(groups) == 0 {
		return nil, errors.New("Data not found")
	}

	// untuk menghitung jumlah member
	var counts []int

	for _, group := range groups {
		var count int64
		if err := database.DB.WithContext(ctx).Model(&model.GroupMember{}).Where("group_id = ?", group.ID).Count(&count).Error; err != nil {
			return nil, errors.New("Data not found")
		}
		counts = append(counts, int(count))
	}

	return groups, nil
}
