package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yzj0930/GoWebWithGin/dto/request"
	"github.com/yzj0930/GoWebWithGin/services"
	util "github.com/yzj0930/GoWebWithGin/utils"
)

type UserController struct {
	BaseController
	UserService services.UserService
}

func (ctrl *UserController) GetUserList(c *gin.Context) {
	var userRequest request.UserListRequest

	c.ShouldBindJSON(&userRequest)
	result, err := ctrl.UserService.GetUserList(userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ReturnError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.ReturnSuccess(result))
}

func (ctrl *UserController) AddUser(c *gin.Context) {
	// Implementation for adding user
	var userRequest request.UserRequest

	c.ShouldBindJSON(&userRequest)
	err := ctrl.UserService.AddUser(&userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ReturnErrorE(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User added successfully"})
}

func (ctrl *UserController) ModifyUser(c *gin.Context) {
	// Implementation for updating user
	var userRequest request.UserRequest
	c.ShouldBindJSON(&userRequest)
	err := ctrl.UserService.ModifyUser(&userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ReturnError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.ReturnSuccess("User updated successfully"))
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
	controller.RequestInfo = append(controller.RequestInfo, RouteConfig{
		Url:      "/modifyuser",
		Method:   "POST",
		Function: controller.ModifyUser,
	})
	return controller
}

func init() {
	Register(NewUserController())
}
