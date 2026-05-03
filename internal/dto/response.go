package dto

import "time"

type AuthResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type CreateGroupResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Icon        string    `json:"icon"`
	InviteCode  string    `json:"invite_code"`
	TotalAmount float64   `json:"total_amount"`
	MemberCount int       `json:"member_count"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserStats struct {
	TotalSplit  int `json:"total_split"`
	ActiveGroup int `json:"active_group"`
	TotalFriend int `json:"total_friend"`
	TotalDrama  int `json:"total_drama"`
}

type GetProfileResponse struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	AvatarUrl   string    `json:"avatar_url"`
	Badge       string    `json:"badge"`
	MemberSince time.Time `json:"member_since"`
	Stats       UserStats `json:"stats"`
}

// === digunakan untuk dashboard ===
type SummaryDashboard struct {
	TotalOwe  float64 `json:"total_owe"`
	TotalOwed float64 `json:"total_owed"`
}

type ActiveGroup struct {
	ID             string        `json:"id"`
	Name           string        `json:"name"`
	Icon           string        `json:"icon"`
	TotalAmount    float64       `json:"total_amount"`
	MemberCount    int           `json:"member_count"`
	MembersPreview []UserPreview `json:"members_preview"`
}

type UserPreview struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	AvatarColor string `json:"avatar_color"`
}

type RecentTransaction struct {
	ID          string      `json:"id"`
	Description string      `json:"description"`
	Amount      float64     `json:"amount"`
	Status      string      `json:"status"`
	Icon        string      `json:"icon"`
	GroupID     string      `json:"group_id"`
	GroupName   string      `json:"group_name"`
	RelatedUser UserPreview `json:"related_user"`
	CreatedAt   time.Time   `json:"created_at"`
	SettledAt   time.Time   `json:"settled_at"`
}

type DashboardResponse struct {
	User               UserPreview         `json:"user"`
	Summary            SummaryDashboard    `json:"summary"`
	ActiveGroups       []ActiveGroup       `json:"active_groups"`
	RecentTransactions []RecentTransaction `json:"recent_transactions"`
}
