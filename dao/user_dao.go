package dao

import (
	"fmt"
	"time"

	"github.com/yzj0930/GoWebWithGin/database"
)

type User struct {
	ID         uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name       string    `gorm:"column:user_name;size:100;not null" json:"user_name"`
	Code       string    `gorm:"column:user_code;size:50;not null;uniqueIndex:uk_user_code" json:"user_code"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

func (User) TableName() string {
	return "users"
}

func GetUserList() ([]User, error) {
	var users []User
	result := database.DB.Find(&users)
	for _, user := range users {
		fmt.Printf("查询到用户: %+v\n", user)
	}
	if result.Error != nil {
		panic("查询用户列表失败: " + result.Error.Error())
	}
	return users, nil
}

func AddUser(user *User) error {
	result := database.DB.Create(user)
	return result.Error
}
