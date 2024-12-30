package service

import (
	. "HiChat/common"
	"HiChat/dao"
	"HiChat/global"
	"HiChat/models"
	"HiChat/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type tempPlatform struct {
	PlatformName string `json:"name"`
	PlatformCode string `json:"code"`
	PlatformUrl  string `json:"url" binding:"http_url" httpUrlMsg:"错误的URL地址"`
	Version      string `json:"version"`
	IsEnable     int    `json:"is_enable"`
}

// CratePlatform 创建
func CratePlatform(ctx *gin.Context) {
	temp := tempPlatform{}
	if err := ctx.ShouldBind(&temp); err != nil {
		errText := utils.GetErrorMsg(err, temp)
		ctx.JSON(http.StatusBadRequest, ParamsNilError.WithMsg(errText))
		return
	}
	findPlatform, err := dao.FindNameToPlatform(temp.PlatformName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ParamsNilError.WithMsg(err.Error()))
		return
	}
	if findPlatform != nil {
		ctx.JSON(http.StatusOK, Error.WithMsg("平台名称已存在"))
		return
	}
	platCode := utils.GenerateRandomString(24)
	platform := models.Platform{
		PlatformName: temp.PlatformName,
		PlatformCode: "plat_" + platCode,
		PlatformUrl:  temp.PlatformUrl,
		Version:      temp.Version,
		IsEnable:     temp.IsEnable,
	}
	_, err = dao.CratePlatform(platform)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ParamsNilError.WithMsg(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, Success.WithMsg("创建成功"))
}

// FindPlatformList 获取平台列表
func FindPlatformList(ctx *gin.Context) {
	list, err := dao.FindPlatformList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ParamsNilError.WithMsg(err.Error()))
		global.Logger.Error("find platform list error", zap.Error(err))
	} else {
		if list == nil {
			list = make([]*models.Platform, 0)
		}
		ctx.JSON(http.StatusOK, Success.WithData(list))
	}
}

// UpdatePlatform 修改
func UpdatePlatform(ctx *gin.Context) {
	type tempUpdatePlatform struct {
		ID           string `json:"id" binding:"required" requiredMsg:"platform id 不能为空"`
		PlatformName string `json:"name"`
		PlatformUrl  string `json:"url,omitempty"`
		Version      string `json:"version,omitempty"`
		IsEnable     int    `json:"is_enable,omitempty"`
	}
	// TODO: chose
	temp := tempUpdatePlatform{}
	reqJson, err := utils.GetJsonAndExistField(ctx, &temp)

	//更新数据
	fmt.Println(reqJson)
	fmt.Println(temp)

	if err != nil {
		errText := utils.GetErrorMsg(err, temp)

		ctx.JSON(http.StatusBadRequest, ParamsNilError.WithMsg(errText))
		return
	}


	//查询id数据
	findPlatform, err := dao.FindIdToPlatform(temp.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ParamsNilError.WithMsg(err.Error()))
		global.Logger.Error("find platform error", zap.Error(err))
		return
	}
	if findPlatform == nil {
		ctx.JSON(http.StatusOK, Success.WithMsg("未查询到数据"))
		return
	}

	//判断名称是否重复
	findPlatform, err = dao.FindNameToPlatform(temp.PlatformName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ParamsNilError.WithMsg(err.Error()))
		global.Logger.Error("find platform error", zap.Error(err))
		return
	}
	if findPlatform != nil {
		ctx.JSON(http.StatusOK, Success.WithMsg("平台名称已存在"))
		return
	}


	//newPlatform, err := dao.UpdatePlatform(temp.ID, reqJson)
 	ctx.JSON(http.StatusOK, reqJson)
}
