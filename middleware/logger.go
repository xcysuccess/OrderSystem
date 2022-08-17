package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger 日志
func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		host := context.Request.Host
		url := context.Request.URL
		method := context.Request.Method
		fmt.Printf("%s::%s \t %s \t %s ", time.Now().Format("2006-01-02 15:04:05"), host, url, method)
		context.Next()
		fmt.Println(context.Writer.Status())
	}
}

// Auth 是否授权
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		println("已经授权")
		context.Next()
	}
}
