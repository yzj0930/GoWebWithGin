package services

import (
	"fmt"

	"github.com/yzj0930/GoWebWithGin/dao"
	"github.com/yzj0930/GoWebWithGin/dto/request"
	"github.com/yzj0930/GoWebWithGin/dto/response"
	"github.com/yzj0930/GoWebWithGin/repositories"
)

type UserService struct {
}

func (s *UserService) GetUserList() []response.UserResponseDto {
	// 调用 DAO 层获取用户列表
	userList := make([]response.UserResponseDto, 0)
	users, err := repositories.GetUserList()
	if err != nil {
		fmt.Printf("获取用户列表失败: %v\n", err)
		return userList
	}
	for _, user := range users {
		userList = append(userList, response.UserResponseDto{
			ID:          user.ID,
			Name:        user.Name,
			Code:        user.Code,
			CreatedTime: user.CreateTime,
			UpdatedTime: user.UpdateTime,
		})
	}
	return userList
}

func (s *UserService) AddUser(user *request.UserRequest) error {
	// 调用 DAO 层添加用户
	userItem := &dao.User{
		Name: user.Name,
		Code: user.Code,
	}
	return repositories.AddUser(userItem)
}

func (s *UserService) ModifyUser(user *request.UserRequest) error {
	return repositories.UpdateUser(user.Code, user.Name)
}
