package types

type CreateUserReq struct {
	Username string `json:"username" binding:"required" copier:"Name"` // 用户名称
	Password string `json:"password" binding:"required" copier:"PasswordHash"`
	Email    string `json:"email" binding:"required"`
}
