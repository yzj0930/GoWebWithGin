package repositories

import (
	"fmt"

	"github.com/yzj0930/GoWebWithGin/dao"
	"github.com/yzj0930/GoWebWithGin/database"
)

func GetUserAuthList(cond map[string]interface{}, limit int, offset int) ([]dao.UserAuth, error) {
    var items []dao.UserAuth
    query := database.DB.Model(&dao.UserAuth{})
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
        return nil, fmt.Errorf("查询用户权限列表失败: %w", res.Error)
    }
    return items, nil
}

func GetUserAuthTotal(cond map[string]interface{}) (int64, error) {
    var total int64
    query := database.DB.Model(&dao.UserAuth{})
    for k, v := range cond {
        query = query.Where(fmt.Sprintf("%s = ?", k), v)
    }
    res := query.Count(&total)
    if res.Error != nil {
        return 0, fmt.Errorf("查询用户权限总数失败: %w", res.Error)
    }
    return total, nil
}

func AddUserAuth(a *dao.UserAuth) error {
    if err := database.DB.Create(a).Error; err != nil {
        return err
    }
    return nil
}
