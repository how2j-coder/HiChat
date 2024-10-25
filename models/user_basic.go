package models

import (
	"HiChat/utils"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type CommonModel struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserBasic struct {
	CommonModel
	Name          string `gorm:"comment '用户名'"`
	PassWord      string `gorm:"comment '密码'"`
	Avatar        string `gorm:"comment '头像'"`
	Gender        string `gorm:"column:gender;default:male;type:varchar(6) comment 'male表示男， female表示女'"` //gorm为数据库字段约束
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`                                             //valid为条件约束
	Email         string `valid:"email"`
	Identity      string
	ClientIp      string `valid:"ipv4"`
	ClientPort    string
	Salt          string     //盐值
	LoginTime     *time.Time `gorm:"column:login_time"`
	HeartBeatTime *time.Time `gorm:"column:heart_beat_time"`
	LoginOutTime  *time.Time `gorm:"column:login_out_time"`
	IsLoginOut    bool
	DeviceInfo    string //登录设备
}

var Snowflake = utils.NewSnowflake(int64(52))

func (table *UserBasic) UserTableName() string {
	return "user_basic"
}
func (com *CommonModel) BeforeCreate(_ *gorm.DB) (err error) {
	com.ID = strconv.FormatInt(Snowflake.GenerateID(), 10)
	return
}
