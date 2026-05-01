package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GroupMember struct {
	ID       uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	GroupID  uuid.UUID `gorm:"type:char(36)" json:"group_id"`
	UserID   uuid.UUID `gorm:"type:char(36)" json:"user_id"`
	JoinedAt time.Time `gorm:"autoCreateTime" json:"joined_at"`

	Group Groups `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE"`
	User  User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (gm *GroupMember) BeforeCreate(tx *gorm.DB) (err error) {
	gm.ID = uuid.New()
	return
}
