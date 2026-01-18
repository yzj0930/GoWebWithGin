package dao

import (
	"time"

	"gorm.io/datatypes"
)

type Projects struct {
    ProjectId   uint            `gorm:"column:project_id;primaryKey;autoIncrement" json:"project_id"`
    ProjectName string          `gorm:"column:project_name;size:100;not null" json:"project_name"`
    ProjectCode string          `gorm:"column:project_code;size:50;not null;uniqueIndex:project_code" json:"project_code"`
    Remark      string          `gorm:"column:remark;size:500;not null" json:"remark"`
    ExtraInfo   datatypes.JSON  `gorm:"column:extra_info;type:json" json:"extra_info"`
    CreateTime  time.Time       `gorm:"column:create_time;autoCreateTime" json:"create_time"`
    UpdateTime  time.Time       `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
    ExpireTime  time.Time       `gorm:"column:expire_time" json:"expire_time"`
}

func (Projects) TableName() string {
    return "projects"
}
