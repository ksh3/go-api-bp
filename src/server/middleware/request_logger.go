package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ksh3/go-api/src/core"
)

func RequestLogger(logger *core.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)

		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path

		logger.InfoLog(
			fmt.Sprintf("[%d] %s %s (%s)", status, method, path, latency))
	}
}
