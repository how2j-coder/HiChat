package types

type CreateMenuReq struct {
	ParentMenuID string `json:"parent_menu_id" binding:""`
	MenuCode string `json:"menu_code" binding:"required"`
	MenuName string `json:"menu_name" binding:"required"`
	MenuPath string `json:"menu_path" binding:"required"`
	MenuSource string `json:"menu_source" binding:"required"`
	IsEnable int8 `json:"is_enable"`
	Type int8 `json:"type"`
	IsRefresh int8 `json:"is_refresh"`
	IsVisible int8 `json:"is_visible"`
}

