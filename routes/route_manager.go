package routes

import (
	"strings"

	"github.com/yzj0930/GoWebWithGin/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	for _, controller := range controllers.GetControllers() {
		for _, route := range controller.GetRequestInfo() {
			fn, ok := route.Function.(func(*gin.Context))
			if !ok {
				// todo: 补充日志文件
				continue
			}
			switch strings.ToUpper(route.Method) {
			case "GET":
				r.GET(route.Url, fn)
			case "POST":
				r.POST(route.Url, fn)
			case "PUT":
				r.PUT(route.Url, fn)
			case "DELETE":
				r.DELETE(route.Url, fn)
			case "PATCH":
				r.PATCH(route.Url, fn)
			case "OPTIONS":
				r.OPTIONS(route.Url, fn)
			case "HEAD":
				r.HEAD(route.Url, fn)
			default:
				// todo: 补充日志文件
				// return errors("unsupported HTTP method: %s", route.Method)
			}
		}
	}
}
