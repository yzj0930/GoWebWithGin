package repositories

import (
	"fmt"

	"github.com/yzj0930/GoWebWithGin/dao"
	"github.com/yzj0930/GoWebWithGin/database"
	"gorm.io/gorm"
)

func GetUserList(cond map[string]interface{}, limit int, offset int) ([]dao.User, error) {
	var users []dao.User
	query := database.DB.Model(&dao.User{})

	// 应用查询条件
	for key, value := range cond {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	err := query.Find(&users)
	if err.Error != nil {
		return nil, fmt.Errorf("查询用户列表失败: %w", err.Error)
	}
	return users, nil
}

func GetUserTotal(cond map[string]interface{}) (int64, error) {
	var total int64
	query := database.DB.Model(&dao.User{})

	// 应用查询条件
	for key, value := range cond {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	err := query.Count(&total)
	if err.Error != nil {
		return 0, fmt.Errorf("查询用户列表失败: %w", err.Error)
	}
	return total, nil
}

func AddUser(user *dao.User) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// 检查用户是否已存在
		var existingUser dao.User
		if err := tx.Where("user_code = ?", user.Code).First(&existingUser).Error; err == nil {
			return fmt.Errorf("用户已存在，无法添加")
		}
		// 添加新用户
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(userCode, userName string) error {
	var user dao.User

	// 开始事务
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// 先检查用户是否存在
		if err := database.DB.Where("user_code = ?", userCode).First(&user).Error; err != nil {
			return fmt.Errorf("用户不存在: %v", err)
		}

		// 更新用户名
		result := database.DB.Model(&user).Update("user_name", userName)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
