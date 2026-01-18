package dao

import (
	"time"

	"gorm.io/datatypes"
)

type Category struct {
    CategoryId   uint           `gorm:"column:category_id;primaryKey;autoIncrement" json:"category_id"`
    ProjectId    uint           `gorm:"column:project_id;not null;uniqueIndex:uk_project_code" json:"project_id"`
    CategoryName string         `gorm:"column:category_name;size:100;not null" json:"category_name"`
    CategoryCode string         `gorm:"column:category_code;size:50;not null;uniqueIndex:uk_project_code" json:"category_code"`
    IsValid      bool           `gorm:"column:is_valid;not null;default:true" json:"is_valid"`
    Remark       string         `gorm:"column:remark;size:500;not null" json:"remark"`
    ExtraInfo    datatypes.JSON `gorm:"column:extra_info;type:json" json:"extra_info"`
    CreateTime   time.Time      `gorm:"column:create_time;autoCreateTime" json:"create_time"`
    UpdateTime   time.Time      `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

func (Category) TableName() string {
    return "category"
}
