package dto

import "time"

type UserStats struct {
	TotalSplit  int `json:"total_split"`
	ActiveGroup int `json:"active_group"`
	TotalFriend int `json:"total_friend"`
	TotalDrama  int `json:"total_drama"`
}

type AuthResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
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
