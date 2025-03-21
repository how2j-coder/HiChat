package handler

import (
	"com/chat/service/internal/dao"
	"com/chat/service/internal/database"
	"com/chat/service/internal/ecode"
	"com/chat/service/internal/handler/common"
	"com/chat/service/internal/model"
	"com/chat/service/internal/types"
	"com/chat/service/pkg/gin/middleware"
	"com/chat/service/pkg/gin/response"
	"com/chat/service/pkg/logger"
	"com/chat/service/pkg/srand"
	"com/chat/service/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type PlatformHandler interface {
	Create(c *gin.Context)
	UpdateById(c *gin.Context)
	GetColumn(c *gin.Context)
	DeleteById(c *gin.Context)
}

type platformHandler struct {
	iDao dao.PlatformDao
}

func NewPlatformHandler() PlatformHandler {
	return &platformHandler{
		iDao: dao.NewPlatformDao(database.GetDB()),
	}
}

func (h *platformHandler) Create(c *gin.Context) {
	form := &types.CreatePlatReq{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	platform := &model.Platform{}
	err = copier.Copy(platform, form)

	if err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	platform.PlatformCode = srand.String(10, 10)

	ctx := middleware.WrapCtx(c)
	err = h.iDao.Create(ctx, platform)
	if err != nil {
		logger.Warn("Create error: ", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

func (h *platformHandler) GetColumn(c *gin.Context) {
	ctx := middleware.WrapCtx(c)
	platforms, err := h.iDao.GetColumn(ctx)
	if err != nil {
		logger.Warn("GetPlatformColumn error: ", logger.Err(err), logger.Any("form", platforms), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	platRes, err := h.convert(platforms)

	if err != nil {
		logger.Error("GetPlatformColumn error", logger.Err(err), logger.Any("Platform", platRes), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}
	response.Success(c, platRes)
}

func (h *platformHandler) UpdateById(c *gin.Context) {
	_, id, isAbort := common.GetIDFromPath(c)

	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	form := &types.UpdatePlatReq{}

	update, err := common.GetTransmitFields(c, form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	platform := &model.Platform{}
	platform.ID = id

	ctx := middleware.WrapCtx(c)
	err = h.iDao.UpdateByID(ctx, platform, update)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}
	response.Success(c)
}

func (h *platformHandler) DeleteById(c *gin.Context) {
	_, id, isAbort := common.GetIDFromPath(c)

	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	platform := model.Platform{}
	platform.ID = id

	ctx := middleware.WrapCtx(c)
	err := h.iDao.DeleteByID(ctx, &platform)

	if err != nil {
		logger.Warn("DeleteById error: ", logger.Err(err), logger.Any("form", platform), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}
	response.Success(c)
}

func (h *platformHandler) convert(platforms []*model.Platform) ([]types.PlatDetailResp, error)  {
	var resPlat []types.PlatDetailResp
	for _, platform := range platforms {
		data := types.PlatDetailResp{}
		data.ID = utils.Uint64ToStr(platform.ID)
		data.IsEnabled = utils.IntToStr(int(platform.IsEnabled))
		err := copier.Copy(&data, platform)
		if err != nil {
			return nil, err
		}
		resPlat = append(resPlat, data)
	}
	return resPlat, nil
}


