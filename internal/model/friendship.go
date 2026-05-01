package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Friendship struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:char(36)" json:"user_id"`
	FriendID  uuid.UUID `gorm:"type:char(36)" json:"friend_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	User   User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Friend User `gorm:"foreignKey:FriendID;constraint:OnDelete:CASCADE"`
}

func (f *Friendship) BeforeCreate(tx *gorm.DB) (err error) {
	f.ID = uuid.New()
	return
}
