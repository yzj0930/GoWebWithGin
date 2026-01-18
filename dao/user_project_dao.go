package dao

import (
	"time"
)

type UserProject struct {
    UserId     uint      `gorm:"column:user_id" json:"user_id"`
    ProjectId  uint      `gorm:"column:project_id" json:"project_id"`
    Status     string    `gorm:"column:status;size:10;default:valid" json:"status"`
    CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

func (UserProject) TableName() string {
    return "user_project"
}
