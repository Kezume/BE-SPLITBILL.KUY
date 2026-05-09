package service

import (
	"context"
	"errors"
	"time"

	"github.com/Kezume/BE-SPLITBILL.KUY/internal/dto"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/model"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/repository"
	"github.com/google/uuid"
)

type ExpenseService interface {
	CreateExpense(ctx context.Context, req dto.CreateExpenseRequest) (*dto.ExpenseResponse, error)
	DeleteExpense(ctx context.Context, expenseID string, userID string) error
}

type expenseService struct {
	repo repository.ExpenseRepository
}

func NewExpenseService(repo repository.ExpenseRepository) ExpenseService {
	return &expenseService{
		repo: repo,
	}
}

func (e *expenseService) CreateExpense(ctx context.Context, req dto.CreateExpenseRequest) (*dto.ExpenseResponse, error) {
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, errors.New("Format tanggal salah! Pakai format YYYY-MM-DD")
	}

	expense := &model.Expense{
		GroupID:     uuid.Must(uuid.Parse(req.GroupID)),
		Description: req.Description,
		Amount:      req.Amount,
		PaidBy:      uuid.Must(uuid.Parse(req.PaidBy)),
		Date:        date,
	}

	err = e.repo.CreateExpense(ctx, expense)
	if err != nil {
		return nil, errors.New("Gagal bikin pengeluaran!")
	}

	var splits []model.ExpenseSplit
	var splitDetails []dto.SplitMemberDetail
	var perPerson uint64

	if req.SplitType == "equal" {
		perPerson = req.Amount / uint64(len(req.SplitWith))
		for _, detail := range req.SplitWith {
			isPayer := detail.UserID == req.PaidBy
			splits = append(splits, model.ExpenseSplit{
				ExpenseID: expense.ID,
				UserID:    uuid.Must(uuid.Parse(detail.UserID)),
				Amount:    perPerson,
				IsSettled: isPayer,
			})
			// Ambil user info untuk response
			user, _ := e.repo.GetUserByID(ctx, detail.UserID)
			avatarUrl := ""
			username := "Unknown"
			if user != nil {
				username = user.Username
				if user.AvatarUrl != nil {
					avatarUrl = *user.AvatarUrl
				}
			}
			splitDetails = append(splitDetails, dto.SplitMemberDetail{
				User: dto.UserPreview{
					ID:        detail.UserID,
					Username:  username,
					AvatarUrl: avatarUrl,
				},
				Amount:    float64(perPerson),
				IsSettled: isPayer,
			})
		}
	} else {
		var totalCustom uint64
		for _, detail := range req.SplitWith {
			totalCustom += detail.Amount
		}
		if totalCustom != req.Amount {
			return nil, errors.New("Total pembagian harus sama dengan total pengeluaran!")
		}

		perPerson = req.Amount / uint64(len(req.SplitWith))
		for _, detail := range req.SplitWith {
			isPayer := detail.UserID == req.PaidBy
			splits = append(splits, model.ExpenseSplit{
				ExpenseID: expense.ID,
				UserID:    uuid.Must(uuid.Parse(detail.UserID)),
				Amount:    detail.Amount,
				IsSettled: isPayer,
			})
			user, _ := e.repo.GetUserByID(ctx, detail.UserID)
			avatarUrl := ""
			username := "Unknown"
			if user != nil {
				username = user.Username
				if user.AvatarUrl != nil {
					avatarUrl = *user.AvatarUrl
				}
			}
			splitDetails = append(splitDetails, dto.SplitMemberDetail{
				User: dto.UserPreview{
					ID:        detail.UserID,
					Username:  username,
					AvatarUrl: avatarUrl,
				},
				Amount:    float64(detail.Amount),
				IsSettled: isPayer,
			})
		}
	}

	if err := e.repo.CreateExpenseSplits(ctx, splits); err != nil {
		return nil, errors.New("Gagal simpan pembagian!")
	}

	// Ambil data user yang bayar untuk response
	payer, err := e.repo.GetUserByID(ctx, req.PaidBy)
	var paidByPreview dto.UserPreview
	if err == nil {
		avatarUrl := ""
		if payer.AvatarUrl != nil {
			avatarUrl = *payer.AvatarUrl
		}
		paidByPreview = dto.UserPreview{
			ID:        payer.ID.String(),
			Username:  payer.Username,
			AvatarUrl: avatarUrl,
		}
	}

	return &dto.ExpenseResponse{
		ID:          expense.ID.String(),
		Description: expense.Description,
		Amount:      float64(expense.Amount),
		PaidBy:      paidByPreview,
		SplitWith:   splitDetails,
		PerPerson:   float64(perPerson),
		Date:        req.Date,
		GroupID:     req.GroupID,
		CreatedAt:   expense.CreatedAt,
	}, nil
}

func (e *expenseService) DeleteExpense(ctx context.Context, expenseID string, userID string) error {
	err := e.repo.DeleteExpense(ctx, expenseID, userID)
	if err != nil {
		return err
	}
	return nil
}
