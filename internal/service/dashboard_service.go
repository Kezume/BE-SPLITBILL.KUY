package service

import (
	"time"

	"github.com/Kezume/BE-SPLITBILL.KUY/internal/dto"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/repository"
)

type DashboardService interface {
	GetDashboardData(userId string) (*dto.DashboardResponse, error)
}

type dashboardService struct {
	repo repository.DashboardRepository
}

func NewDashboardService(repo repository.DashboardRepository) DashboardService {
	return &dashboardService{
		repo: repo,
	}
}

func (s *dashboardService) GetDashboardData(userId string) (*dto.DashboardResponse, error) {
	user := dto.UserPreview{
		ID:          userId,
		Username:    "USER_01",
		AvatarColor: "bg-yellow-400",
	}

	summary := dto.SummaryDashboard{
		TotalOwe:  150000,
		TotalOwed: 200000,
	}

	activeGroups := []dto.ActiveGroup{
		{
			ID:          "GROUP_01",
			Name:        "Makan Bakso",
			Icon:        "🍜",
			TotalAmount: 150000,
			MemberCount: 3,
			MembersPreview: []dto.UserPreview{
				{
					ID:          "USER_01",
					Username:    "USER_01",
					AvatarColor: "bg-yellow-400",
				},
				{
					ID:          "USER_02",
					Username:    "USER_02",
					AvatarColor: "bg-red-400",
				},
			},
		},
	}

	recentTransactions := []dto.RecentTransaction{
		{
			ID:          "TXN_01",
			Description: "Makan Bakso",
			Amount:      150000,
			Status:      "paid",
			Icon:        "🍜",
			GroupID:     "GROUP_01",
			GroupName:   "Makan Bakso",
			RelatedUser: dto.UserPreview{
				ID:          "USER_01",
				Username:    "USER_01",
				AvatarColor: "bg-yellow-400",
			},
			CreatedAt: time.Now(),
			SettledAt: time.Now(),
		},
	}

	return &dto.DashboardResponse{
		User:               user,
		Summary:            summary,
		ActiveGroups:       activeGroups,
		RecentTransactions: recentTransactions,
	}, nil
}
