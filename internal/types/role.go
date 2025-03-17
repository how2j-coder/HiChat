package types

type CreateRoleReq struct {
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