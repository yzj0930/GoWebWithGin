package dao

import (
	"time"
)

type UserAuth struct {
    UserId      uint      `gorm:"column:user_id" json:"user_id"`
    PrivilegeId uint      `gorm:"column:privilege_id" json:"privilege_id"`
    Status      string    `gorm:"column:status;size:10;default:valid" json:"status"`
    CreateTime  time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
    UpdateTime  time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

func (UserAuth) TableName() string {
    return "user_auth"
}
