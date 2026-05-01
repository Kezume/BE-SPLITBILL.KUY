package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Groups struct {
	ID         uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Name       string    `gorm:"type:varchar(100)" json:"name"`
	Icon       string    `gorm:"type:varchar(50)" json:"icon"`
	InviteCode string    `gorm:"type:varchar(20)" json:"invite_code"`
	OwnerID    uuid.UUID `gorm:"type:char(36)" json:"owner_id"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Owner User `gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE"`
}

func (g *Groups) BeforeCreate(tx *gorm.DB) (err error) {
	g.ID = uuid.New()
	return
}
