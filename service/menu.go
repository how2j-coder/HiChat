package service

import (
	. "HiChat/common"
	"HiChat/dao"
	"HiChat/global"
	"HiChat/models"
	"HiChat/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateMenu 创建菜单
func CreateMenu(ctx *gin.Context) {
	temp := models.Menu{}
	if err := ctx.ShouldBind(&temp); err != nil {
		errText := utils.GetErrorMsg(err, temp)
		ctx.JSON(http.StatusBadRequest, ParamsNilError.WithMsg(errText))
		return
	}

	findPlatform, err := dao.FindIdToPlatform(temp.PlatformID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ParamsNilError.WithMsg(err.Error()))
		return
	}

	if findPlatform == nil {
		ctx.JSON(http.StatusOK, Error.WithMsg("未查询到平台信息"))
		return
	}

	findMenu, err := dao.FindMenuCodeToMenu(temp.MenuCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ParamsNilError.WithMsg(err.Error()))
		return
	}

	if findMenu != nil {
		ctx.JSON(http.StatusOK, Error.WithMsg("菜单Code不能重复"))
		return
	}

	if temp.ParentMenuID != "" {
		findMenu, _ = dao.FindIdToMenu(temp.ParentMenuID)
		if findMenu == nil {
			ctx.JSON(http.StatusOK, ParamsNilError.WithMsg("未查询到父级菜单"))
			return
		}
	}

	_, err = dao.CreateMenu(temp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ParamsNilError.WithMsg(err.Error()))
		global.Logger.Error(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, Success.WithMsg("success"))
}

// UpdateMenu 更新菜单数据
func UpdateMenu(ctx *gin.Context) {
	temp := struct {
		ID string `json:"id" binding:"required" requiredMsg:"菜单ID不能为空"`
		PlatformID string `json:"platform_id" has_required:"请选择平台"`
		MenuName string `json:"menu_name" has_required:"菜单名称不能为空"`
		MenuCode string `json:"menu_code" has_required:"菜单Code不能为空"`
		MenuPath string `json:"menu_path" has_required:"菜单路由地址不能为空"`
		MenuFilePath string `json:"menu_file_path" has_required:"模板路径不能为空"`
	}{}

	reqJson, err := utils.GetJsonAndExistField(ctx, &temp)
	if err != nil {
		errText := utils.GetErrorMsg(err, temp)
		ctx.JSON(http.StatusInternalServerError, ParamsNilError.WithMsg(errText))
		return
	}

	if _, exists := reqJson["platform_id"]; exists {
		if reqJson["platform_id"] == "" {
			ctx.JSON(http.StatusOK, ParamsNilError.WithMsg("请选择平台"))
			return
		}
	}

	findMenu, err := dao.FindIdToMenu(temp.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ParamsNilError.WithMsg(err.Error()))
		return
	}

	if findMenu == nil {
		ctx.JSON(http.StatusOK, Error.WithMsg("未查询到菜单数据"))
		return
	}

	findMenu, err = dao.FindMenuCodeToMenu(temp.MenuCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ParamsNilError.WithMsg(err.Error()))
		return
	}

	if findMenu != nil {
		ctx.JSON(http.StatusOK, Error.WithMsg("菜单Code不能重复"))
		return
	}
	fmt.Println(reqJson)

	_, err = dao.UpdateMenuIdToMenu(temp.ID, reqJson)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ParamsNilError.WithMsg(err.Error()))
		global.Logger.Error(err.Error())
		return
	}
	ctx.JSON(http.StatusOK, Success.WithMsg("success"))
}

// GetMenuTree 获取菜单树
func GetMenuTree(ctx *gin.Context) {
	allPlatform, err := dao.FindPlatformList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ParamsNilError.WithMsg(err.Error()))
		return
	}

	var menuTree [][]*models.Menu
	//查找所有平台
	for i := 0; i < len(allPlatform); i++ {
		//获取对应平台的所有菜单
		platformToMenus, _ := dao.FindPlatformToMenus(allPlatform[i].ID)
		treeMenu := buildMenuTree(platformToMenus, "")
		menuTree = append(menuTree, treeMenu)
	}

	ctx.JSON(http.StatusOK, Success.WithData(menuTree))
}

func buildMenuTree(menus []*models.Menu, parentId string) []*models.Menu {
	menuTree := make([]*models.Menu, 0)
	for _, menu := range menus {
		if menu.ParentMenuID == parentId {
			children := buildMenuTree(menus, menu.ID)
			menu.Children = children
			menuTree  = append(menuTree, menu)
		}
	}
	return menuTree
}
