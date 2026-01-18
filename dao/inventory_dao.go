package dao

import (
	"time"
)

type Inventory struct {
    Id           uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
    ProjectId    uint      `gorm:"column:project_id;not null;uniqueIndex:uk_project_category" json:"project_id"`
    CategoryId   uint      `gorm:"column:category_id;not null;uniqueIndex:uk_project_category" json:"category_id"`
    TotalNum     int64     `gorm:"column:total_num;not null;default:0" json:"total_num"`
    AvailableNum int64     `gorm:"column:available_num;not null;default:0" json:"available_num"`
    SoldNum      int64     `gorm:"column:sold_num;not null;default:0" json:"sold_num"`
    CreateTime   time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
    UpdateTime   time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

func (Inventory) TableName() string {
    return "inventory"
}
