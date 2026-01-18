package repositories

import (
	"fmt"

	"github.com/yzj0930/GoWebWithGin/dao"
	"github.com/yzj0930/GoWebWithGin/database"
	"gorm.io/gorm"
)

func GetCategoryList(cond map[string]interface{}, limit int, offset int) ([]dao.Category, error) {
    var items []dao.Category
    query := database.DB.Model(&dao.Category{})
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
        return nil, fmt.Errorf("查询货品列表失败: %w", res.Error)
    }
    return items, nil
}

func GetCategoryTotal(cond map[string]interface{}) (int64, error) {
    var total int64
    query := database.DB.Model(&dao.Category{})
    for k, v := range cond {
        query = query.Where(fmt.Sprintf("%s = ?", k), v)
    }
    res := query.Count(&total)
    if res.Error != nil {
        return 0, fmt.Errorf("查询货品总数失败: %w", res.Error)
    }
    return total, nil
}

func AddCategory(c *dao.Category) error {
    err := database.DB.Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(c).Error; err != nil {
            return err
        }
        return nil
    })
    return err
}

func UpdateCategoryByID(id uint, updates map[string]interface{}) error {
    res := database.DB.Model(&dao.Category{}).Where("category_id = ?", id).Updates(updates)
    if res.Error != nil {
        return res.Error
    }
    return nil
}
