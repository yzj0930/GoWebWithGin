package repositories

import (
	"fmt"

	"github.com/yzj0930/GoWebWithGin/dao"
	"github.com/yzj0930/GoWebWithGin/database"
	"gorm.io/gorm"
)

func GetSkuList(cond map[string]interface{}, limit int, offset int) ([]dao.Sku, error) {
    var items []dao.Sku
    query := database.DB.Model(&dao.Sku{})
    for k, v := range cond {
        query = query.Where(fmt.Sprintf("%s = ?", k), v)
    }
    if limit > 0 {
        query = query.Limit(limit)
    }
    if offset > 0 {
        query = query.Offset(offset)
    }
    res := query.Find(&items)
    if res.Error != nil {
        return nil, fmt.Errorf("查询商品列表失败: %w", res.Error)
    }
    return items, nil
}

func GetSkuTotal(cond map[string]interface{}) (int64, error) {
    var total int64
    query := database.DB.Model(&dao.Sku{})
    for k, v := range cond {
        query = query.Where(fmt.Sprintf("%s = ?", k), v)
    }
    res := query.Count(&total)
    if res.Error != nil {
        return 0, fmt.Errorf("查询商品总数失败: %w", res.Error)
    }
    return total, nil
}

func AddSku(s *dao.Sku) error {
    err := database.DB.Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(s).Error; err != nil {
            return err
        }
        return nil
    })
    return err
}

func UpdateSkuByID(id uint, updates map[string]interface{}) error {
    res := database.DB.Model(&dao.Sku{}).Where("sku_id = ?", id).Updates(updates)
    if res.Error != nil {
        return res.Error
    }
    return nil
}
