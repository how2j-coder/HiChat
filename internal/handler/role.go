package handler

import (
	"com/chat/service/internal/dao"
	"com/chat/service/internal/database"
	"com/chat/service/internal/ecode"
	"com/chat/service/internal/handler/common"
	"com/chat/service/internal/model"
	"com/chat/service/internal/types"
	"com/chat/service/pkg/datastore/query"
	"com/chat/service/pkg/gin/middleware"
	"com/chat/service/pkg/gin/response"
	"com/chat/service/pkg/logger"
	"com/chat/service/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

var _ RoleHandler = (*roleHandler)(nil)

type roleHandler struct {
	iDao dao.RoleDao
}

type RoleHandler interface {
	Create(c *gin.Context)
	UpdateByID(c *gin.Context)
	DeleteByID(c *gin.Context)
	GetColumn(c *gin.Context)
	SetUserRole(c *gin.Context)
	SetMenuRole(c *gin.Context)
}

func NewRoleHandler() RoleHandler {
	return &roleHandler{
		iDao: dao.NewRoleDao(database.GetDB()),
	}
}

func (h *roleHandler) Create(c *gin.Context) {
	form := &types.CreateRoleReq{}

	err := c.ShouldBindJSON(form)

	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	findRole, err := h.iDao.GetByName(ctx, form.RoleName)

	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	if findRole != nil {
		response.Error(c, ecode.ErrCreateRole.RewriteMsg("角色已存在"))
		return
	}

	role := &model.Role{}
	err = copier.Copy(role, form)
	if err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	err = h.iDao.Create(ctx, role)

	if err != nil {
		logger.Warn("Create error: ", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c, gin.H{
		"message": "success",
	})
}
func (h *roleHandler) UpdateByID(c *gin.Context) {
	_, id, isAbort := common.GetIDFromPath(c)

	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	form := &types.UpdateRoleReq{}

	// 获取需要更新的数据
	update, err := common.GetTransmitFields(c, form)
	if err != nil {
		logger.Warn("ShouldBindBodyWithJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)

	findRepeatRole, err := h.iDao.GetByNameExcID(ctx, form.RoleName, id)

	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	if findRepeatRole != nil {
		response.Error(c, ecode.ErrCreateRole.RewriteMsg("角色名称已存在"))
		return
	}

	role := &model.Role{}
	role.ID = id

	err = h.iDao.UpdateByID(ctx, role, update)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}
	response.Success(c)

}
func (h *roleHandler) DeleteByID(c *gin.Context) {
	_, id, isAbort := common.GetIDFromPath(c)

	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}
	ctx := middleware.WrapCtx(c)
	err := h.iDao.DeleteByID(ctx, id)

	if err != nil {
		logger.Warn("DeleteByID error: ", logger.Err(err), logger.Any("id", id), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}
	response.Success(c)
}
func (h *roleHandler) GetColumn(c *gin.Context) {
	form := &types.GetRoleListReq{}

	err := c.ShouldBindQuery(form)
	if err != nil {
		logger.Warn("ShouldBindQuery error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	params := query.Params{
		RowColumns: *form,
		Columns: []query.Column{
			{
				Name:  "role_name",
				Value: "Li",
				Exp:   query.Like,
			},
		},
	}

	ctx := middleware.WrapCtx(c)
	findData, total, err := h.iDao.GetByColumns(ctx, &params)

	data := make([]types.ListRoleDetail, 0)
	err = copier.Copy(&data, &findData)

	response.Success(c, gin.H{
		"list":  data,
		"total": total,
	})
}
func (h *roleHandler) SetUserRole(c *gin.Context)  {
	form := &types.SetUserRoleReq{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	roleIds := utils.StrListToUint64(form.RoleIDs)
	userId := utils.StrToUint64(form.UserID)
	ctx := middleware.WrapCtx(c)

	err = h.iDao.SetUserRole(ctx, roleIds, userId)
	if err != nil {
		logger.Warn("SetUserRole error: ", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

func (h *roleHandler) SetMenuRole(c *gin.Context)  {
	form := &types.SetMenuRoleReq{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	menuIds := utils.StrListToUint64(form.MenuIDs)
	roleId := utils.StrToUint64(form.RoleIDs)
	ctx := middleware.WrapCtx(c)

	err = h.iDao.SetMenuRole(ctx,roleId, menuIds)
	if err != nil {
		logger.Warn("SetMenuRole error: ", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}