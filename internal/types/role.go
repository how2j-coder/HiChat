package types

type CreateRoleReq struct {
	ParentRoleID string `json:"parent_role_id" binding:""`
	RoleName string `json:"role_name" binding:"required"`
	Remark string `json:"remark" binding:""`
}

type UpdateRoleReq struct {
	RoleName string `json:"role_name" binding:""`
	Remark string `json:"remark" binding:""`
}

type GetRoleListReq struct {
	RoleName string `json:"role_name" form:"role_name" binding:""`
}

type ListRoleDetail struct {
	ID uint64 `json:"role_id" binding:""`
	RoleName string `json:"role_name" binding:""`
	Remark string `json:"remark" binding:""`
}

type SetUserRoleReq struct {
	UserID string `json:"user_id" binding:""`
	RoleIDs []string `json:"role_ids" binding:""`
}

type SetMenuRoleReq struct {
	RoleIDs string `json:"role_id" binding:""`
	MenuIDs []string `json:"menu_ids" binding:""`
}