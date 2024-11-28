package models

import (
	"time"
)

type User struct {
	CommonBaseModel `gorm:"embedded"`
	Name            string     `gorm:"comment:用户名" json:"name"`
	PassWord        string     `gorm:"comment:密码" json:"password,omitempty"`
	Avatar          string     `gorm:"comment:头像" json:"avatar"`
	Gender          string     `gorm:"column:gender;default:male;type:varchar(6);comment:male表示男，female表示女" json:"gender"` //gorm为数据库字段约束
	Phone           string     `valid:"matches(^1[3-9]{1}\\d{9}$)" json:"phone"`                                           //valid为条件约束
	Email           string     `valid:"email" json:"email"`
	Identity        string     `gorm:"comment:用户身份" json:"identity,omitempty"`
	ClientIp        string     `valid:"ipv4" gorm:"comment:设备IP" json:"client_ip,omitempty"`
	ClientPort      string     `gorm:"comment:设备端口" json:"client_port,omitempty"`
	Salt            string     `gorm:"comment:用户密码MD5盐值" json:"salt,omitempty"`
	LoginTime       *time.Time `gorm:"column:login_time" json:"login_time,omitempty"`
	HeartBeatTime   *time.Time `gorm:"column:heart_beat_time" json:"heart_beat_time"`
	LoginOutTime    *time.Time `gorm:"column:login_out_time" json:"login_out_time,omitempty"`
	IsLoginOut      bool       `json:"is_login_out"`
	DeviceInfo      string     `json:"device_info"` //登录设备
}

func (table *User) UserTableName() string {
	return "users"
}
