package handler

import (
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/dto"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/service"
	"github.com/Kezume/BE-SPLITBILL.KUY/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
	service service.ExpenseService
}

func NewExpenseHandler(service service.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{
		service: service,
	}
}

func (e *ExpenseHandler) CreateExpense(c *gin.Context) {
	userID := c.GetString("user_id")

	if userID == "" {
		response.Error(c, 401, "Lo belum login nih, login dulu!")
		return
	}

	var req *dto.CreateExpenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Ada yang salah nih, coba lagi ya!")
		return
	}

	data, err := e.service.CreateExpense(c, *req)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, data)
}

func (e *ExpenseHandler) DeleteExpense(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		response.Error(c, 401, "Lo belum login nih, login dulu!")
		return
	}

	expenseID := c.Param("id")
	if expenseID == "" {
		response.Error(c, 400, "Waduh, ada yang salah nih!")
		return
	}

	err := e.service.DeleteExpense(c, expenseID, userID)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, "Pengeluaran berhasil dihapus!")
}
