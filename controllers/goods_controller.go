package controllers

import (
	"github.com/yzj0930/GoWebWithGin/services"
)

type GoodsController struct {
	BaseController
	GoodsService services.GoodsService
}



func NewGoodsController() Controller {
	controller := &GoodsController{
		BaseController: BaseController{RequestInfo: make([]RouteConfig, 0)},
	}
	
	return controller
}

func init() {
	Register(NewGoodsController())
}
