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

func buildUserListConditions(param request.UserListRequest) map[string]interface{} {
	cond := make(map[string]interface{})
	if param.Name != "" {
		cond["user_name"] = param.Name
	}
	if param.Code != "" {
		cond["user_code"] = param.Code
	}
	return cond
}

func (s *UserService) GetUserList(param request.UserListRequest) ([]response.UserResponseDto, error) {
	// 调用 DAO 层获取用户列表
	userList := make([]response.UserResponseDto, 0)
	cond := buildUserListConditions(param)
	if param.Page <= 0 {
		param.Page = 1
	}
	if param.PageSize <= 0 {
		param.PageSize = 10
	}
	limit := param.PageSize
	offset := (param.Page - 1) * param.PageSize
	users, err := repositories.GetUserList(cond, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("获取用户列表失败: %v", err)
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
	return userList, nil
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
	// 调用 DAO 层更新用户
	return repositories.UpdateUser(user.Code, user.Name)
}
