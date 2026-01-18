package dao

import (
	"time"
)

type Privileges struct {
	PrivilegeId     uint      `gorm:"column:privilege_id;primaryKey;autoIncrement" json:"privilege_id"`
	PrivilegeName   string    `gorm:"column:privilege_name;size:100;not null" json:"privilege_name"`
	PrivilegeCode   string    `gorm:"column:privilege_code;size:100;not null;uniqueIndex:uk_privilege_code" json:"privilege_code"`
	Description  string	 `gorm:"column:description;size:20;not null" json:"description"`
	Remark   string	 `gorm:"column:remark;size:500;not null" json:"remark"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

func (Privileges) TableName() string {
	return "privileges"
}
