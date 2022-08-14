// Package router TODO
package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化路由
func SetupRouter() *gin.Engine {
	// 1.创建路由
	router := gin.Default()
	// 2.绑定路由规则，执行的函数
	router.GET("/", retHelloGinAndMethod)
	router.POST("/", retHelloGinAndMethod)

	// gin.Context，封装了request和response
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	return router
}

// retHelloGinAndMethod 封装了request和response
// gin.Context
func retHelloGinAndMethod(context *gin.Context) {
	str := fmt.Sprintf("hello gin " + strings.ToLower(context.Request.Method) + " method")
	context.String(http.StatusOK, str)
}
