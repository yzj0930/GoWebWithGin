package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yzj0930/GoWebWithGin/dto/request"
	"github.com/yzj0930/GoWebWithGin/services"
)

type PingController struct {
	BaseController
	PingService services.PingService
}

func (ctrl *PingController) Ping(c *gin.Context) {
	message := ctrl.PingService.GetPingMessage()
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func (ctrl *PingController) MockPostJson(c *gin.Context) {
	var obj request.PostJsonRequest

	c.ShouldBindJSON(&obj)
	responseObj := ctrl.PingService.GetPostJson(obj)
	c.JSON(http.StatusOK, responseObj)
}

func (ctrl *PingController) MockPostForm(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	c.JSON(http.StatusOK, gin.H{"username": username, "password": password})
}

func (ctrl *PingController) MockGetQuery(c *gin.Context) {
	param := c.Query("param")
	c.JSON(http.StatusOK, gin.H{"param": param})
}

func (ctrl *PingController) MockGetPath(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (ctrl *PingController) MockPut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "PUT request received"})
}

func NewPingController() *PingController {
	controller := &PingController{
		BaseController: NewBaseController(),
	}
	controller.RequestInfo = append(controller.RequestInfo, RouteConfig{
		Url:      "/ping",
		Method:   "GET",
		Function: controller.Ping,
	})
	controller.RequestInfo = append(controller.RequestInfo, RouteConfig{
		Url:      "/mock_post_json",
		Method:   "POST",
		Function: controller.MockPostJson,
	})
	controller.RequestInfo = append(controller.RequestInfo, RouteConfig{
		Url:      "/mock_post_form",
		Method:   "POST",
		Function: controller.MockPostForm,
	})
	controller.RequestInfo = append(controller.RequestInfo, RouteConfig{
		Url:      "/mock_get_query",
		Method:   "GET",
		Function: controller.MockGetQuery,
	})
	controller.RequestInfo = append(controller.RequestInfo, RouteConfig{
		Url:      "/mock_get_path/:id",
		Method:   "GET",
		Function: controller.MockGetPath,
	})
	controller.RequestInfo = append(controller.RequestInfo, RouteConfig{
		Url:      "/mock_put",
		Method:   "PUT",
		Function: controller.MockPut,
	})

	return controller
}

func init() {
	Register(NewPingController())
}
