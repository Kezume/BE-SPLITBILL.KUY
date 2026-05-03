package repository

import (
	"context"

	"github.com/Kezume/BE-SPLITBILL.KUY/internal/model"
	"github.com/Kezume/BE-SPLITBILL.KUY/pkg/database"
)

type GroupRepository interface {
	CreateGroup(ctx context.Context, group *model.Groups) error
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
