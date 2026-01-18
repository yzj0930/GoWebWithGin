package repositories

import (
	"fmt"

	"github.com/yzj0930/GoWebWithGin/dao"
	"github.com/yzj0930/GoWebWithGin/database"
	"gorm.io/gorm"
)

func GetInventoryList(cond map[string]interface{}, limit int, offset int) ([]dao.Inventory, error) {
    var items []dao.Inventory
    query := database.DB.Model(&dao.Inventory{})
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
        return nil, fmt.Errorf("查询库存列表失败: %w", res.Error)
    }
    return items, nil
}

func GetInventoryTotal(cond map[string]interface{}) (int64, error) {
    var total int64
    query := database.DB.Model(&dao.Inventory{})
    for k, v := range cond {
        query = query.Where(fmt.Sprintf("%s = ?", k), v)
    }
    res := query.Count(&total)
    if res.Error != nil {
        return 0, fmt.Errorf("查询库存总数失败: %w", res.Error)
    }
    return total, nil
}

func AddInventory(it *dao.Inventory) error {
    err := database.DB.Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(it).Error; err != nil {
            return err
        }
        return nil
    })
    return err
}

func UpdateInventoryByID(id uint, updates map[string]interface{}) error {
    res := database.DB.Model(&dao.Inventory{}).Where("id = ?", id).Updates(updates)
    if res.Error != nil {
        return res.Error
    }
    return nil
}
