package types

type CreateMenuReq struct {
	ParentMenuID uint64 `json:"parent_menu_id" binding:"number"`
	MenuCode string `json:"menu_code" binding:"required"`
	MenuName string `json:"menu_name" binding:"required"`
	MenuPath string `json:"menu_path" binding:"required"`
	MenuSource string `json:"menu_source" binding:"required"`
	IsEnable int8 `json:"is_enable"`
	Type int8 `json:"type"`
	IsRefresh int8 `json:"is_refresh"`
	IsVisible int8 `json:"is_visible"`
}

type UpdateMenuReq struct {
	ParentMenuID uint64 `json:"parent_menu_id" binding:""`
	MenuCode *string `json:"menu_code" binding:"optional_not_empty"`
	MenuName *string `json:"menu_name" binding:"optional_not_empty"`
	MenuPath *string `json:"menu_path" binding:"optional_not_empty"`
	MenuSource *string `json:"menu_source" binding:"optional_not_empty"`
	IsEnable int8 `json:"is_enable"`
	Type int8 `json:"type"`
	IsRefresh int8 `json:"is_refresh"`
	IsVisible int8 `json:"is_visible"`
}

type ListMenuDetail struct {
	ID uint64 `json:"menu_id" binding:""`
	ParentMenuID uint64 `json:"parent_menu_id" binding:""`
	MenuCode string `json:"menu_code" binding:""`
	MenuName string `json:"menu_name" binding:""`
	MenuPath string `json:"menu_path" binding:""`
	MenuSource string `json:"menu_source" binding:""`
	IsEnable int8 `json:"is_enable"`
	Type int8 `json:"type"`
	IsRefresh int8 `json:"is_refresh"`
	IsVisible int8 `json:"is_visible"`
}