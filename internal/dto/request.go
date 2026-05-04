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
