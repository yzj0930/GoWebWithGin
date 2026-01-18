package repositories

import (
	"fmt"

	"github.com/yzj0930/GoWebWithGin/dao"
	"github.com/yzj0930/GoWebWithGin/database"
)

func GetUserProjectList(cond map[string]interface{}, limit int, offset int) ([]dao.UserProject, error) {
    var items []dao.UserProject
    query := database.DB.Model(&dao.UserProject{})
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
        return nil, fmt.Errorf("查询用户-项目关联列表失败: %w", res.Error)
    }
    return items, nil
}

func GetUserProjectTotal(cond map[string]interface{}) (int64, error) {
    var total int64
    query := database.DB.Model(&dao.UserProject{})
    for k, v := range cond {
        query = query.Where(fmt.Sprintf("%s = ?", k), v)
    }
    res := query.Count(&total)
    if res.Error != nil {
        return 0, fmt.Errorf("查询用户-项目关联总数失败: %w", res.Error)
    }
    return total, nil
}

func AddUserProject(up *dao.UserProject) error {
    if err := database.DB.Create(up).Error; err != nil {
        return err
    }
    return nil
}
