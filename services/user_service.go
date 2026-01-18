package services

import (
	"fmt"
	"time"

	"github.com/yzj0930/GoWebWithGin/dao"
	"github.com/yzj0930/GoWebWithGin/dto/request"
	"github.com/yzj0930/GoWebWithGin/dto/response"
	"github.com/yzj0930/GoWebWithGin/logger"
	"github.com/yzj0930/GoWebWithGin/repositories"
	util "github.com/yzj0930/GoWebWithGin/utils"
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

func (s *UserService) GetUserList(param request.UserListRequest) (response.UserListResponse, error) {
	// 调用 DAO 层获取用户列表
	res := response.UserListResponse{}
	res.List = make([]response.UserResponseDto, 0)
	cond := buildUserListConditions(param)
	if param.Page <= 0 {
		param.Page = 1
	}
	if param.PageSize <= 0 {
		param.PageSize = 10
	}
	limit := param.PageSize
	offset := (param.Page - 1) * param.PageSize
	users, total, err := repositories.GetUserListWithTotal(cond, limit, offset)
	if err != nil {
		return res, fmt.Errorf("获取用户列表失败: %v", err)
	}
	res.Total = total
	for _, user := range users {
		res.List = append(res.List, response.UserResponseDto{
			ID:          user.UserId,
			Name:        user.UserName,
			Code:        user.UserCode,
			CreatedTime: user.CreateTime,
			UpdatedTime: user.UpdateTime,
		})
	}
	logger.Info("获取用户列表成功，数量：", total)
	return res, nil
}

func (s *UserService) AddUser(user *request.UserRequest) error {
	// 调用 DAO 层添加用户
	pwdHash, err := util.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("密码加密失败: %v", err)
	}
	userItem := &dao.Users{
		UserName: user.Name,
		UserCode: user.Code,
		Telephone: user.Telephone,
		Password: pwdHash,
	}
	return repositories.AddUser(userItem)
}

func (s *UserService) ModifyUser(user *request.UserRequest) error {
	// 调用 DAO 层更新用户
	if user.Code == "" {
		return fmt.Errorf("用户编码不能为空")
	}
	pwdHash, err := util.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("密码加密失败: %v", err)
	}
	userItem := &dao.Users{
		UserName: user.Name,
		UserCode: user.Code,
		Telephone: user.Telephone,
		Password: pwdHash,
	}
	return repositories.UpdateUser(userItem)
}

func (s *UserService) UserLogin(userInfo *request.UserLoginRequest) (string, error) {
	// 调用 DAO 层验证用户登录
	user, err := repositories.GetUserByCode(userInfo.Code)
	if err != nil {
		return "", fmt.Errorf("获取用户信息失败: %v", err)
	}
	if user == nil {
		return "", fmt.Errorf("用户不存在")
	}
	if !util.CheckPasswordHash(userInfo.Password, user.Password) {
		return "", fmt.Errorf("密码错误")
	}
	jwtManager := util.NewJWTManager("your-secret-key", 24*7*time.Hour) // 7天有效期
	token, err := jwtManager.GenerateToken(user.UserId, user.UserName, "", "user")
	if err != nil {
		return "", fmt.Errorf("生成Token失败: %v", err)
	}
	return token, nil
}
