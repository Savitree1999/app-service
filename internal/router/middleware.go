package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GinLoggerMiddleware(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		logger.Infow("Request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"latency", duration.String(),
			"clientIP", c.ClientIP(),
		)
	}
}
