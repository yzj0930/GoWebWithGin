package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yzj0930/GoWebWithGin/dto/request"
	"github.com/yzj0930/GoWebWithGin/dto/response"
	"github.com/yzj0930/GoWebWithGin/services"
)

type UserController struct {
	BaseController
	UserService services.UserService
}

func (ctrl *UserController) GetUserList(c *gin.Context) {
	// Implementation for getting user
	results := response.ResponseDto{
		Status:  0,
		Message: "Success",
		Data:    ctrl.UserService.GetUserList(),
	}
	c.JSON(http.StatusOK, results)
}

func (ctrl *UserController) AddUser(c *gin.Context) {
	// Implementation for adding user
	var userRequest request.UserRequest

	c.ShouldBindJSON(&userRequest)
	ctrl.UserService.AddUser(&userRequest)

	c.JSON(http.StatusOK, gin.H{"message": "User added successfully"})
}

func NewUserController() Controller {
	controller := &UserController{
		BaseController: BaseController{RequestInfo: make([]RouteConfig, 0)},
	}
	controller.RequestInfo = append(controller.RequestInfo, RouteConfig{
		Url:      "/userlist",
		Method:   "GET",
		Function: controller.GetUserList,
	})
	controller.RequestInfo = append(controller.RequestInfo, RouteConfig{
		Url:      "/adduser",
		Method:   "POST",
		Function: controller.AddUser,
	})
	return controller
}

func init() {
	Register(NewUserController())
}
