package types

type CreateMenuReq struct {
	ParentMenuID string `json:"parent_menu_id" binding:"" copier:"-"`
	PlatformID string `json:"platform_id" binding:"number" copier:"-"`
	MenuCode     string `json:"menu_code" binding:"required"`
	MenuName     string `json:"menu_name" binding:"required"`
	MenuPath     string `json:"menu_path" binding:"required"`
	MenuSource   string `json:"menu_source" binding:"required"`
	IsEnable     int8   `json:"is_enable"`
	Type         int8   `json:"type"`
	IsRefresh    int8   `json:"is_refresh"`
	IsVisible    int8   `json:"is_visible"`
}

type UpdateMenuReq struct {
	ParentMenuID uint64  `json:"parent_menu_id" binding:""`
	MenuCode     *string `json:"menu_code" binding:"optional_not_empty"`
	MenuName     *string `json:"menu_name" binding:"optional_not_empty"`
	MenuPath     *string `json:"menu_path" binding:"optional_not_empty"`
	MenuSource   *string `json:"menu_source" binding:"optional_not_empty"`
	MenuIcon     string `json:"menu_icon"`
	Type         int8    `json:"type"`
	IsRefresh    int8    `json:"is_refresh"`
	IsVisible    int8    `json:"is_visible"`
	IsEnable     int8    `json:"is_enable"`
	IsSingle     int8   `json:"is_single"`
}

type ListMenuDetail struct {
	ID           string `json:"menu_id" copier:"-"`
	PlatformID   string `json:"platform_id" copier:"-"`
	ParentMenuID string `json:"parent_menu_id" copier:"-"`
	MenuCode     string `json:"menu_code" `
	MenuName     string `json:"menu_name" `
	MenuPath     string `json:"menu_path" `
	MenuIcon     string `json:"menu_icon" `
	MenuSource   string `json:"menu_source" `
	Type         int8   `json:"type"`
	Sort         int8   `json:"sort"`
	IsRefresh    int8   `json:"is_refresh"`
	IsEnable     int8   `json:"is_enable"`
	IsVisible    int8   `json:"is_visible"`
	IsSingle     int8   `json:"is_single"`
}

type GetMenuListReq struct {
	PlatformID string `json:"platform_id" form:"platform_id" binding:""`
	MenuID     string `json:"menu_id" form:"menu_id" binding:""`
}
