package handler

import (
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/dto"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/model"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/service"
	"github.com/Kezume/BE-SPLITBILL.KUY/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type GroupHandler struct {
	service service.GroupService
}

func NewGroupHandler(service service.GroupService) GroupHandler {
	return GroupHandler{
		service: service,
	}
}

func (g *GroupHandler) CreateGroup(c *gin.Context) {
	userID := c.GetString("user_id")

	if userID == "" {
		response.Error(c, 401, "Unauthorized: Missing Token")
		return
	}

	var payload *dto.CreateGroup

	if err := c.ShouldBindJSON(&payload); err != nil {
		response.Error(c, 400, "Failed to Bind JSON")
		return
	}

	group := &model.Groups{
		Name: payload.Name,
		Icon: payload.Icon,
	}

	err := g.service.CreateGroup(c.Request.Context(), group, userID)
	if err != nil {
		response.Error(c, 500, "Failed to Create Group")
		return
	}

	response.Success(c, response.ToGroupResponse(group))
}
