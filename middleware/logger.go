package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yzj0930/GoWebWithGin/logger"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()
		latency := endTime.Sub(startTime)

		logger.Info(fmt.Sprintf("%s %s %d %s", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), latency))
	}
}
