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
	"com/chat/service/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type MenuHandler interface {
	Create(c *gin.Context)
	UpdateByID(c *gin.Context)
	DeleteByID(c *gin.Context)
	GetByPlatformIDToColumn(c *gin.Context)
	GetColumn(c *gin.Context)
	GetDetailByID(c *gin.Context)
}

type menuHandler struct {
	iDao dao.MenuDao
}

var _ MenuHandler = (*menuHandler)(nil)

func NewMenuHandler() MenuHandler {
	return &menuHandler{
		iDao: dao.NewMenuDao(database.GetDB()),
	}
}

// Create 创建菜单
func (m *menuHandler) Create(c *gin.Context) {
	form := &types.CreateMenuReq{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	findMenu, _ := m.iDao.GetByMenuCode(ctx, form.MenuCode)
	if findMenu != nil {
		response.Error(c, ecode.ErrCreateMenu.RewriteMsg("菜单Code重复"))
		return
	}

	menu := &model.Menu{}
	err = copier.Copy(menu, form)
	menu.ParentMenuID = utils.StrToUint64(form.ParentMenuID)
	menu.PlatformID = utils.StrToUint64(form.PlatformID)
	if err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	err = m.iDao.Create(ctx, menu)

	if err != nil {
		logger.Warn("Create error: ", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c, gin.H{"id": menu.ID})
}

// UpdateByID 更新菜单数据
func (m *menuHandler) UpdateByID(c *gin.Context) {
	_, id, isAbort := common.GetIDFromPath(c)

	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	form := &types.UpdateMenuReq{}

	// 获取需要更新的数据
	update, err := common.GetTransmitFields(c, form)
	if err != nil {
		logger.Warn("ShouldBindBodyWithJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	menu := &model.Menu{}
	menu.ID = id

	ctx := middleware.WrapCtx(c)
	err = m.iDao.UpdateByID(ctx, menu, update)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}
	response.Success(c)
}

// DeleteByID 删除菜单
func (m *menuHandler) DeleteByID(c *gin.Context) {
	_, id, isAbort := common.GetIDFromPath(c)

	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}
	menu := &model.Menu{}
	menu.ID = id

	ctx := middleware.WrapCtx(c)
	err := m.iDao.DeleteByID(ctx, menu)
	if err != nil {
		logger.Error("DeleteByID error", logger.Err(err), logger.Any("menu", menu), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}
	response.Success(c)
}


// GetByPlatformIDToColumn 根据平台获取子级菜单
func (m *menuHandler) GetByPlatformIDToColumn(c *gin.Context)  {
	form := &types.GetMenuListReq{}
	err := c.ShouldBindQuery(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	menuID := utils.StrToUint64(form.MenuID)
	platformID := utils.StrToUint64(form.PlatformID)

	ctx := middleware.WrapCtx(c)
	menus, err := m.iDao.GetByPlatParentID(ctx, menuID, platformID)

	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	menuList, err := m.convert(menus)

	if err != nil {
		logger.Error("GetColumn Copy error", logger.Err(err), logger.Any("menus", menus), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c, menuList)
}


func (m *menuHandler) GetColumn(c *gin.Context) {
	// TODO: 这里根据角色权限去获取
	ctx := middleware.WrapCtx(c)
	menus, err := m.iDao.GetColumn(ctx)
	if err != nil {
		logger.Error("GetColumn error", logger.Err(err), logger.Any("menus", menus), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	menuList, err := m.convert(menus)

	if err != nil {
		logger.Error("GetColumn Copy error", logger.Err(err), logger.Any("menus", menus), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c,  menuList)
}

func (m *menuHandler) GetDetailByID(c *gin.Context) {
	_, id, isAbort := common.GetIDFromPath(c)

	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	findMenu, err := m.iDao.GetDetailByID(ctx, id)
	if err != nil {
		logger.Warn("GetByID error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	menu := types.ListMenuDetail{}
	menu.ID = utils.Uint64ToStr(findMenu.ID)
	menu.ParentMenuID = utils.Uint64ToStr(findMenu.ParentMenuID)
	menu.PlatformID = utils.Uint64ToStr(findMenu.PlatformID)

	err = copier.Copy(&menu, findMenu)

	if err != nil {
		logger.Warn("GetDetailByID error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c, menu)
}

func (m *menuHandler) convert(menus []*model.Menu) ([]types.ListMenuDetail, error) {
	menuList := make([]types.ListMenuDetail, 0)
	for _, menu := range menus {
		data := types.ListMenuDetail{}
		data.ID = utils.Uint64ToStr(menu.ID)
		data.ParentMenuID = utils.Uint64ToStr(menu.ParentMenuID)
		data.PlatformID = utils.Uint64ToStr(menu.PlatformID)
		err := copier.Copy(&data, menu)
		if err != nil {
			return nil, err
		}
		menuList = append(menuList, data)
	}
	return menuList, nil
}