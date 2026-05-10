package repository

import (
	"context"
	"errors"

	"github.com/Kezume/BE-SPLITBILL.KUY/internal/model"
	"github.com/Kezume/BE-SPLITBILL.KUY/pkg/database"
)

type ExpenseRepository interface {
	CreateExpense(ctx context.Context, expense *model.Expense) error
	CreateExpenseSplits(ctx context.Context, splits []model.ExpenseSplit) error
	DeleteExpense(ctx context.Context, expenseID string, userId string) error
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
}

type expenseRepository struct {
}

func NewExpenseRepository() ExpenseRepository {
	return &expenseRepository{}
}

func (e *expenseRepository) CreateExpense(ctx context.Context, expense *model.Expense) error {
	err := database.DB.WithContext(ctx).Create(&expense).Error
	if err != nil {
		return err
	}
	return nil
}

func (e *expenseRepository) CreateExpenseSplits(ctx context.Context, splits []model.ExpenseSplit) error {
	err := database.DB.WithContext(ctx).Create(&splits).Error
	if err != nil {
		return err
	}
	return nil
}

func (e *expenseRepository) DeleteExpense(ctx context.Context, expenseID string, userID string) error {
	result := database.DB.WithContext(ctx).Where("id = ? AND paid_by = ?", expenseID, userID).Delete(&model.Expense{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Expense gak ketemu atau lo bukan yang bayar!")
	}

	return nil
}

func (e *expenseRepository) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	var user model.User
	if err := database.DB.WithContext(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
