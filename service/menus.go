package service

import (
	"github.com/gin-gonic/gin"
)

type tempMenus struct {
	PlatformID string `json:"platform_id" binding:"required"`
	ParentMenuID int `json:"parent_menu_id" binding:"required"`
	MenuName string `json:"menu_name"`
	MenuCode string `json:"menu_code"`
	MenuType int `json:"menu_type"`
	MenuFilePath string  `json:"menu_file_path"`
	IsVisible int `json:"is_visible"`
	IsEnabled int `json:"is_enabled"`
	IsRefresh int `json:"is_refresh"`
	SortOrder int `json:"sort_order"`
}

func CreateMus(ctx *gin.Context)  {

}
