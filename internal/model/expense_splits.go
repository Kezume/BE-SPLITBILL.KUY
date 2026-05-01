package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ExpenseSplit struct {
	ID        uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id"`
	ExpenseID uuid.UUID  `gorm:"type:char(36)" json:"expense_id"`
	UserID    uuid.UUID  `gorm:"type:char(36)" json:"user_id"`
	Amount    uint64     `gorm:"type:bigint" json:"amount"`
	IsSettled bool       `gorm:"type:boolean;default:false" json:"is_settled"`
	SettledAt *time.Time `gorm:"type:timestamp" json:"settled_at"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`

	Expense Expense `gorm:"foreignKey:ExpenseID;constraint:OnDelete:CASCADE"`
	User    User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (es *ExpenseSplit) BeforeCreate(tx *gorm.DB) (err error) {
	es.ID = uuid.New()
	return
}
