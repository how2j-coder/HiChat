package types

type CreateUserReq struct {
	Account string `json:"account" binding:""`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type UpdateUserReq struct {
	ID       uint64 `json:"id" binding:"-"`
	Username string `json:"username" binding:""`
	Gender   string `json:"gender" binding:""`
	AvatarURL string `json:"avatar_url" binding:""`
}


type UserReq interface {
	GetPassword() string
}

type UserEmailLoginReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
func (u *UserEmailLoginReq) GetPassword() string {
	return u.Password
}

type UserAccountLogoutReq struct {
	Account string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}
func (u *UserAccountLogoutReq) GetPassword() string {
	return u.Password
}
