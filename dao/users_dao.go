package dao

import (
	"time"
)

type Users struct {
	UserId     uint      `gorm:"column:user_id;primaryKey;autoIncrement" json:"user_id"`
	UserName   string    `gorm:"column:user_name;size:100;not null" json:"user_name"`
	UserCode   string    `gorm:"column:user_code;size:50;not null;uniqueIndex:uk_user_code" json:"user_code"`
	Telephone  string	 `gorm:"column:telephone;size:100;not null" json:"telephone"`
	Password   string	 `gorm:"column:password;size:20;not null" json:"password"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

func (Users) TableName() string {
	return "users"
}
