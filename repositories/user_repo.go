package repositories

import (
	"fmt"

	"github.com/yzj0930/GoWebWithGin/dao"
	"github.com/yzj0930/GoWebWithGin/database"
	"gorm.io/gorm"
)

func GetUserList() ([]dao.User, error) {
	var users []dao.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		panic("查询用户列表失败: " + result.Error.Error())
	}
	return users, nil
}

func AddUser(user *dao.User) error {
	result := database.DB.Create(user)
	return result.Error
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
