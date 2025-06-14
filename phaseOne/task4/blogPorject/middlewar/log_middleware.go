package middleware

import (
	"log"
	"path"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

func LogRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		pc, file, line, ok := runtime.Caller(2)
		fileName := "unknown"
		funcName := "unknown"
		if ok {
			fileName = path.Base(file)
			funcName = runtime.FuncForPC(pc).Name()
		}
		log.Printf("[GIN] %s %s | %d | %v | %s:%d | %s",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			latency,
			fileName,
			line,
			funcName,
		)
	}
}
