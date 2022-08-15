// Package router TODO
package router

import (
	"fmt"
	"net/http"
	"ordersystem/service"
	"strings"

	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化路由
func SetupRouter() *gin.Engine {
	router := gin.Default()
	// 1.配置HTML
	initHTMLConfig(router)
	// 2.创建路由
	indexRouter := router.Group("/")
	{
		// 2.1 绑定路由规则，执行的函数
		// index.GET("", retHelloGinAndMethod)
		// index.POST("", retHelloGinAndMethod)
		indexRouter.Any("", service.Index)
	}
	userRouter := router.Group("/user")
	{
		userRouter.GET("/:name", service.UserSave)
		userRouter.GET("/", service.UserSaveNameByQuery)
		userRouter.POST("/register", service.UserRegister)
	}
	return router
}

func initHTMLConfig(router *gin.Engine) {
	// 1.配置网页路径
	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../frontdev/templates/*")
	} else {
		router.LoadHTMLGlob("./frontdev/templates/*")
	}
	// 2.配置css和js路径
	router.Static("/bootstrap", "./frontdev")
	// 2.1 加载图标
	router.StaticFile("/favicon.ico", "./frontdev/images/favicon.ico")
}

// retHelloGinAndMethod 封装了request和response
// gin.Context
func retHelloGinAndMethod(context *gin.Context) {
	str := fmt.Sprintf("hello gin " + strings.ToLower(context.Request.Method) + " method")
	context.String(http.StatusOK, str)
}
