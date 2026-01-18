package repositories

import (
	"fmt"

	"github.com/yzj0930/GoWebWithGin/dao"
	"github.com/yzj0930/GoWebWithGin/database"
	"gorm.io/gorm"
)

func GetUserList(cond map[string]interface{}, limit int, offset int) ([]dao.Users, error) {
	var users []dao.Users
	query := database.DB.Model(&dao.Users{})

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
	query := database.DB.Model(&dao.Users{})

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

func GetUserListWithTotal(cond map[string]interface{}, limit int, offset int) ([]*dao.Users, int, error) {
	var users []*dao.Users
	total := int64(0)
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		
		query := database.DB.Model(&dao.Users{})

		// 应用查询条件
		for key, value := range cond {
			query = query.Where(fmt.Sprintf("%s = ?", key), value)
		}
		err := query.Count(&total)
		if err.Error != nil {
			return fmt.Errorf("查询用户数量失败: %w", err.Error)
		}
		if limit > 0 {
			query = query.Limit(limit)
		}
		if offset > 0 {
			query = query.Offset(offset)
		}

		err = query.Find(&users)
		if err.Error != nil {
			return fmt.Errorf("查询用户列表失败: %w", err.Error)
		}
		return nil
	})
	if err != nil {
		return nil, int(0), err
	}
	return users, int(total), nil
}

func GetUserByCode(userCode string) (*dao.Users, error) {
	var user dao.Users
	err := database.DB.Where("user_code = ?", userCode).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("根据用户编码查询用户失败: %v", err)
	}
	return &user, nil
}

func AddUser(user *dao.Users) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// 检查用户是否已存在
		var existingUser dao.Users
		if err := tx.Where("user_code = ?", user.UserCode).First(&existingUser).Error; err == nil {
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

func UpdateUser(modifyUser *dao.Users) error {
	var user dao.Users

	// 开始事务
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// 先检查用户是否存在
		if err := database.DB.Where("user_code = ?", modifyUser.UserCode).First(&user).Error; err != nil {
			return fmt.Errorf("用户不存在: %v", err)
		}

		// 更新用户名
		updateMap := make(map[string]interface{})
		if modifyUser.UserName != "" {
			updateMap["user_name"] = modifyUser.UserName
		}
		// 更新电话
		if modifyUser.Telephone != "" {
			updateMap["telephone"] = modifyUser.Telephone
		}
		// 更新密码
		if modifyUser.Password != "" {
			updateMap["password"] = modifyUser.Password
		}
		result := database.DB.Model(&user).Updates(updateMap)
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
