package repositories

import (
	"fmt"

	"github.com/yzj0930/GoWebWithGin/dao"
	"github.com/yzj0930/GoWebWithGin/database"
	"gorm.io/gorm"
)

func GetProjectsList(cond map[string]interface{}, limit int, offset int) ([]dao.Projects, error) {
    var items []dao.Projects
    query := database.DB.Model(&dao.Projects{})
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
        return nil, fmt.Errorf("查询项目列表失败: %w", res.Error)
    }
    return items, nil
}

func GetProjectsTotal(cond map[string]interface{}) (int64, error) {
    var total int64
    query := database.DB.Model(&dao.Projects{})
    for k, v := range cond {
        query = query.Where(fmt.Sprintf("%s = ?", k), v)
    }
    res := query.Count(&total)
    if res.Error != nil {
        return 0, fmt.Errorf("查询项目总数失败: %w", res.Error)
    }
    return total, nil
}

func AddProject(p *dao.Projects) error {
    err := database.DB.Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(p).Error; err != nil {
            return err
        }
        return nil
    })
    return err
}

func UpdateProjectByID(id uint, updates map[string]interface{}) error {
    res := database.DB.Model(&dao.Projects{}).Where("project_id = ?", id).Updates(updates)
    if res.Error != nil {
        return res.Error
    }
    return nil
}
