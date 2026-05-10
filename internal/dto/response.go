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

type ListGroupResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Icon        string    `json:"icon"`
	InviteCode  string    `json:"invite_code"`
	TotalAmount float64   `json:"total_amount"`
	MemberCount int       `json:"member_count"`
	CreatedAt   time.Time `json:"created_at"`
}

type Meta struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

type GroupDetailStats struct {
	TotalSpent   float64 `json:"total_spent"`
	YourShare    float64 `json:"your_share"`
	UnpaidAmount float64 `json:"unpaid_amount"`
	IsSettled    bool    `json:"is_settled"`
}

// Struct flat untuk GORM scan (tidak bisa pakai nested struct)
type ExpenseScanResult struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Payer       string  `json:"payer"`
	PayerID     string  `json:"payer_id"`
	PayerAvatar string  `json:"payer_avatar"`
	Date        string  `json:"date"`
	GroupID     string  `json:"group_id"`
}

// Detail split per member
type SplitMemberDetail struct {
	User      UserPreview `json:"user"`
	Amount    float64     `json:"amount"`
	IsSettled bool        `json:"is_settled"`
}

// Flat scan untuk expense_splits
type SplitScanResult struct {
	UserID    string  `json:"user_id"`
	Username  string  `json:"username"`
	AvatarUrl string  `json:"avatar_url"`
	Amount    float64 `json:"amount"`
	IsSettled bool    `json:"is_settled"`
}

type ExpenseResponse struct {
	ID          string              `json:"id"`
	Description string              `json:"description"`
	Amount      float64             `json:"amount"`
	PaidBy      UserPreview         `json:"paid_by"`
	SplitWith   []SplitMemberDetail `json:"split_with"`
	PerPerson   float64             `json:"per_person"`
	Date        string              `json:"date"`
	GroupID     string              `json:"group_id"`
	CreatedAt   time.Time           `json:"created_at"`
}

type GroupDetailResponse struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Icon        string            `json:"icon"`
	InviteCode  string            `json:"invite_code"`
	TotalAmount float64           `json:"total_amount"`
	MemberCount int               `json:"member_count"`
	IsOwner     bool              `json:"is_owner"`
	CreatedAt   time.Time         `json:"created_at"`
	Stats       GroupDetailStats  `json:"stats"`
	Members     []UserPreview     `json:"members"`
	Expenses    []ExpenseResponse `json:"expenses"`
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
	ID        string `json:"id"`
	Username  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
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

type FriendResponse struct {
	ID           string  `json:"id"`
	Username     string  `json:"username"`
	AvatarColor  string  `json:"avatar_color"`
	MutualGroups int     `json:"mutual_groups"`
	Balance      float64 `json:"balance"`
	Status       string  `json:"status"`
}

type FriendStats struct {
	Total     int `json:"total"`
	Settled   int `json:"settled"`
	Unsettled int `json:"unsettled"`
}

type FriendListResponse struct {
	Friends []FriendResponse `json:"friends"`
	Stats   FriendStats      `json:"stats"`
}
