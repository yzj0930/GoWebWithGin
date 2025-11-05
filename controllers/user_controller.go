package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yzj0930/GoWebWithGin/services"
)

type UserController struct {
	BaseController
	UserService services.UserService
}

func (ctrl *UserController) GetUserList(c *gin.Context) {
	// Implementation for getting user
	ctrl.UserService.GetUserList()
	c.JSON(200, gin.H{"message": "User list"})
}

func NewUserController() Controller {
	controller := &UserController{
		BaseController: BaseController{RequestInfo: make([]RouteConfig, 0)},
	}
	controller.RequestInfo = append(controller.RequestInfo, RouteConfig{
		Url:      "/user",
		Method:   "GET",
		Function: controller.GetUserList,
	})
	return controller
}

func init() {
	Register(NewUserController())
}
