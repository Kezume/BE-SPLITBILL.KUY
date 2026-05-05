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
		response.Error(c, 401, "Lo belum login nih, login dulu!")
		return
	}

	var payload *dto.CreateGroup

	if err := c.ShouldBindJSON(&payload); err != nil {
		response.Error(c, 400, "Ada yang salah nih, coba lagi ya!")
		return
	}

	group := &model.Groups{
		Name: payload.Name,
		Icon: payload.Icon,
	}

	err := g.service.CreateGroup(c.Request.Context(), group, userID)
	if err != nil {
		response.Error(c, 500, "Gagal bikin grup, coba lagi!")
		return
	}

	response.Success(c, response.ToGroupResponse(group))
}

func (g *GroupHandler) GetListGroup(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		response.Error(c, 401, "Lo belum login nih, login dulu!")
		return
	}

	pagination := dto.PaginationRequest{
		Page:  1,
		Limit: 10,
	}

	c.ShouldBindQuery(&pagination)

	if pagination.Limit < 1 {
		pagination.Limit = 10
	}

	if pagination.Page < 1 {
		pagination.Page = 1
	}

	groups, meta, err := g.service.GetListGroup(c.Request.Context(), userID, &pagination)
	if err != nil {
		response.Error(c, 500, "Gagal ambil daftar grup lo!")
		return
	}

	response.SuccessWithMeta(c, groups, meta)
}

func (g *GroupHandler) GetGroupDetail(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		response.Error(c, 401, "Lo belum login nih, login dulu!")
		return
	}

	groupID := c.Param("id")
	if groupID == "" {
		response.Error(c, 400, "Waduh, ada yang salah nih!")
		return
	}

	data, err := g.service.GetGroupDetail(c.Request.Context(), groupID, userID)
	if err != nil {
		response.Error(c, 404, "Grup gak ketemu atau lo gak punya akses!")
		return
	}

	response.Success(c, data)
}

func (g *GroupHandler) DeleteGroup(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		response.Error(c, 401, "Lo belum login nih, login dulu!")
		return
	}

	groupID := c.Param("id")
	if groupID == "" {
		response.Error(c, 400, "Waduh, ada yang salah nih!")
		return
	}

	err := g.service.DeleteGroup(c, groupID, userID)
	if err != nil {
		response.Error(c, 500, "Gagal hapus grup! Cuma owner yang bisa hapus!")
		return
	}

	response.Success(c, "Group deleted successfully")
}

func (g *GroupHandler) JoinGroup(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		response.Error(c, 401, "Lo belum login nih, login dulu!")
		return
	}

	var payload *dto.JoinGroupRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		response.Error(c, 400, "Ada yang salah nih, coba lagi ya!")
		return
	}

	data, err := g.service.JoinGroup(c.Request.Context(), payload.InviteCode, userID)
	if err != nil {
		response.Error(c, 404, "Gagal gabung ke grup!")
		return
	}

	response.Success(c, data)
}
