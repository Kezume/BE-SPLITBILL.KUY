package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Username  string    `gorm:"type:varchar(30)" json:"username"`
	Email     string    `gorm:"type:varchar(100)" json:"email"`
	Phone     string    `gorm:"type:varchar(20)" json:"phone"`
	Password  string    `gorm:"type:varchar(255)" json:"password"`
	AvatarUrl *string   `gorm:"type:varchar(255)" json:"avatar_url"`
	Badge     string    `gorm:"type:varchar(20)" json:"badge"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
