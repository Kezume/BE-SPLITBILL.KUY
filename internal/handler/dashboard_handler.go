package handler

import (
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/service"
	"github.com/Kezume/BE-SPLITBILL.KUY/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	service service.DashboardService
}

func NewDashboardHandler(service service.DashboardService) DashboardHandler {
	return DashboardHandler{
		service: service,
	}
}

func (h *DashboardHandler) GetDashboard(c *gin.Context) {
	userID := c.GetString("user_id")

	if userID == "" {
		response.Error(c, 401, "Lo belum login nih, login dulu!")
		return
	}

	data, err := h.service.GetDashboardData(userID)
	if err != nil {
		response.Error(c, 500, "Gagal muat data dashboard lo!")
		return
	}

	response.Success(c, data)
}
