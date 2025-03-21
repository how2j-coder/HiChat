package types

type CreatePlatReq struct {
	PlatformName string `json:"platform_name" binding:"required"`
	PlatformCode string `json:"platform_code"`
	IsEnable int `json:"is_enable" binding:"required"`
	PlatformUrl string `json:"platform_url" binding:"required"`
}

type UpdatePlatReq struct {
	PlatformName string `json:"platform_name" binding:""`
	IsEnabled int `json:"is_enabled" binding:""`
	PlatformUrl string `json:"platform_url" binding:""`
}

type PlatDetailResp struct {
	ID string `json:"platform_id" copier:"-"`
	PlatformName string `json:"platform_name"`
	PlatformCode string `json:"platform_code" `
	IsEnabled string `json:"is_enabled" copier:"-"`
	PlatformUrl string `json:"platform_url"`
}