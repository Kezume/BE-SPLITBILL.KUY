package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Expense struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	GroupID     uuid.UUID `gorm:"type:char(36)" json:"group_id"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	Amount      uint64    `gorm:"type:bigint" json:"amount"`
	PaidBy      uuid.UUID `gorm:"type:char(36)" json:"paid_by"`
	Date        time.Time `gorm:"type:date" json:"date"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (e *Expense) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New()
	return
}
