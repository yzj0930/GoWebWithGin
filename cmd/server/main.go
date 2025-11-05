package main

import (
	"fmt"
	"runtime/debug"

	"github.com/yzj0930/GoWebWithGin/config"
	"github.com/yzj0930/GoWebWithGin/database"
	"github.com/yzj0930/GoWebWithGin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到 panic: %v\n", r)
			// 这里可以记录日志、清理资源等
			debug.PrintStack()
		}
	}()
	config.LoadYAMLConfig("config/config.yaml")
	database.InitDB()
	port := config.GlobalConfig.Port
	r := gin.Default()
	// 定义路由
	routes.RegisterRoutes(r)
	// 启动服务
	r.Run(fmt.Sprintf(":%d", port)) // 默认监听 8080 端口
}
