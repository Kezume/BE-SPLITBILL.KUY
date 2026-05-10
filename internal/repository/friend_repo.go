package repository

import (
	"context"

	"github.com/Kezume/BE-SPLITBILL.KUY/internal/dto"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/model"
	"github.com/Kezume/BE-SPLITBILL.KUY/pkg/database"
	"github.com/google/uuid"
)

type FriendRepository interface {
	AddFriend(ctx context.Context, userID string, friendID string) error
	GetFriends(ctx context.Context, userID string, search string, status string) ([]dto.FriendResponse, error)
	RemoveFriend(ctx context.Context, userID string, friendID string) error
	CheckFriendship(ctx context.Context, userID string, friendID string) (bool, error)
}

type friendRepository struct {
}

func NewFriendRepository() FriendRepository {
	return &friendRepository{}
}

func (f *friendRepository) AddFriend(ctx context.Context, userID string, friendID string) error {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return err
	}
	friendUUID, err := uuid.Parse(friendID)
	if err != nil {
		return err
	}

	friend := model.Friendship{
		UserID:   userUUID,
		FriendID: friendUUID,
	}

	return database.DB.WithContext(ctx).Create(&friend).Error
}

type rawFriendData struct {
	ID           string
	Username     string
	AvatarColor  string
	MutualGroups int
	Balance      float64
}

func (f *friendRepository) GetFriends(ctx context.Context, userID string, search string, status string) ([]dto.FriendResponse, error) {
	query := `
	WITH FriendIDs AS (
		SELECT CASE WHEN user_id = ? THEN friend_id ELSE user_id END as f_id
		FROM friendships
		WHERE user_id = ? OR friend_id = ?
	)
	SELECT 
		u.id, 
		u.username, 
		u.avatar_color,
		(
			SELECT COUNT(*) 
			FROM group_members gm1 
			JOIN group_members gm2 ON gm1.group_id = gm2.group_id 
			WHERE gm1.user_id = ? AND gm2.user_id = u.id
		) as mutual_groups,
		(
			COALESCE((
				SELECT SUM(es.amount) 
				FROM expense_splits es 
				JOIN expenses e ON es.expense_id = e.id 
				WHERE e.paid_by = ? AND es.user_id = u.id AND es.is_settled = false
			), 0)
			-
			COALESCE((
				SELECT SUM(es.amount) 
				FROM expense_splits es 
				JOIN expenses e ON es.expense_id = e.id 
				WHERE e.paid_by = u.id AND es.user_id = ? AND es.is_settled = false
			), 0)
		) as balance
	FROM users u
	JOIN FriendIDs f ON u.id = f.f_id
	`
	
	if search != "" {
		query += " WHERE u.username LIKE '%" + search + "%'"
	}

	var rawData []rawFriendData
	err := database.DB.WithContext(ctx).Raw(query, userID, userID, userID, userID, userID, userID).Scan(&rawData).Error
	if err != nil {
		return nil, err
	}

	var friends []dto.FriendResponse
	for _, raw := range rawData {
		// Filter by status if provided
		if status == "settled" && raw.Balance != 0 {
			continue
		}
		if status == "unsettled" && raw.Balance == 0 {
			continue
		}

		statusStr := "LUNAS"
		if raw.Balance != 0 {
			statusStr = "HUTANG"
		}

		friends = append(friends, dto.FriendResponse{
			ID:           raw.ID,
			Username:     raw.Username,
			AvatarColor:  raw.AvatarColor,
			MutualGroups: raw.MutualGroups,
			Balance:      raw.Balance,
			Status:       statusStr,
		})
	}

	return friends, nil
}

func (f *friendRepository) RemoveFriend(ctx context.Context, userID string, friendID string) error {
	return database.DB.WithContext(ctx).
		Where("(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)", userID, friendID, friendID, userID).
		Delete(&model.Friendship{}).Error
}

func (f *friendRepository) CheckFriendship(ctx context.Context, userID string, friendID string) (bool, error) {
	var count int64
	err := database.DB.WithContext(ctx).Model(&model.Friendship{}).
		Where("(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)", userID, friendID, friendID, userID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
