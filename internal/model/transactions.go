package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID          uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id"`
	Description string     `gorm:"type:varchar(255)" json:"description"`
	Status      string     `gorm:"type:varchar(20)" json:"status"`
	Icon        string     `gorm:"type:varchar(50)" json:"icon"`
	FromUser    uuid.UUID  `gorm:"type:char(36)" json:"from_user"`
	ToUser      uuid.UUID  `gorm:"type:char(36)" json:"to_user"`
	ExpenseID   *uuid.UUID `gorm:"type:char(36)" json:"expense_id"`
	GroupID     *uuid.UUID `gorm:"type:char(36)" json:"group_id"`
	Amount      uint64     `gorm:"type:bigint" json:"amount"`
	SettledAt   *time.Time `gorm:"type:timestamp" json:"settled_at"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`

	Sender   User     `gorm:"foreignKey:FromUser;constraint:OnDelete:CASCADE"`
	Receiver User     `gorm:"foreignKey:ToUser;constraint:OnDelete:CASCADE"`
	Expense  *Expense `gorm:"foreignKey:ExpenseID;constraint:OnDelete:SET NULL"`
	Group    *Groups  `gorm:"foreignKey:GroupID;constraint:OnDelete:SET NULL"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}
