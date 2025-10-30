package controllers

type Controller interface {
	GetRequestInfo() []RouteConfig
}

type RouteConfig struct {
	Url      string
	Method   string
	Function interface{}
}

type BaseController struct {
	RequestInfo []RouteConfig
}

func NewBaseController() BaseController {
	return BaseController{RequestInfo: make([]RouteConfig, 0)}
}

func (ctrl *BaseController) GetRequestInfo() []RouteConfig {
	return ctrl.RequestInfo
}
