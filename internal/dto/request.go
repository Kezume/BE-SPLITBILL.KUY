package dto

type RegisterUser struct {
	Username string `json:"username" validate:"required,min=5"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,max=13"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UpdateProfile struct {
	Username string `json:"username" validate:"min=5"`
	Phone    string `json:"phone" validate:"max=13"`
	Email    string `json:"email" validate:"email"`
}

type CreateGroup struct {
	Name string `json:"name" validate:"required"`
	Icon string `json:"icon"`
}

type PaginationRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type JoinGroupRequest struct {
	InviteCode string `json:"invite_code" validate:"required"`
}

type CreateExpenseRequest struct {
	GroupID     string        `json:"group_id" validate:"required"`
	Description string        `json:"description" validate:"required"`
	Amount      uint64        `json:"amount" validate:"required"`
	PaidBy      string        `json:"paid_by" validate:"required"`
	SplitType   string        `json:"split_type" validate:"required"`
	SplitWith   []SplitDetail `json:"split_with" validate:"required"`
	Date        string        `json:"date" validate:"required"`
}

type SplitDetail struct {
	UserID string `json:"user_id" validate:"required"`
	Amount uint64 `json:"amount"`
}
