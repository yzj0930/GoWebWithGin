package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/yzj0930/GoWebWithGin/logger"
)

func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 获取堆栈信息
				stack := debug.Stack()

				// 记录错误日志
				logger.Error(fmt.Sprintf("Panic recovered: %v\n%s", err, stack))

				// 根据环境决定返回信息
				if gin.Mode() == gin.ReleaseMode {
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": "内部服务器错误",
					})
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": fmt.Sprintf("%v", err),
						"stack": string(stack),
					})
				}

				c.Abort()
			}
		}()

		c.Next()
	}
}
