package dao

import (
	"time"
)

type Sku struct {
    SkuId     uint      `gorm:"column:sku_id;primaryKey;autoIncrement" json:"sku_id"`
    ProjectId uint      `gorm:"column:project_id;not null;uniqueIndex:uk_project_sku" json:"project_id"`
    SkuName   string    `gorm:"column:sku_name;size:100;not null" json:"sku_name"`
    SkuCode   string    `gorm:"column:sku_code;size:50;not null;uniqueIndex:uk_project_sku" json:"sku_code"`
    Price     float64   `gorm:"column:price;not null;default:0" json:"price"`
    ExpPrice  float64   `gorm:"column:exp_price;not null;default:0" json:"exp_price"`
    ActPrice  float64   `gorm:"column:act_price;not null;default:0" json:"act_price"`
    IsSold    bool      `gorm:"column:is_sold;not null;default:false" json:"is_sold"`
    CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

func (Sku) TableName() string {
    return "sku"
}
