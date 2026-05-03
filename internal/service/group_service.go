package service

import (
	"context"
	"errors"

	"github.com/Kezume/BE-SPLITBILL.KUY/internal/model"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/repository"
	"github.com/Kezume/BE-SPLITBILL.KUY/pkg/utils"
	"github.com/google/uuid"
)

type GroupService interface {
	CreateGroup(ctx context.Context, group *model.Groups, userID string) error
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
