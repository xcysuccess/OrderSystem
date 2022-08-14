// Package router TODO
package router

import (
	"fmt"
	"net/http"
	"ordersystem/handler"
	"strings"

	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化路由
func SetupRouter() *gin.Engine {
	router := gin.Default()
	// 1.配置HTML
	// initHTMLConfig(router)
	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}
	// 2.创建路由
	indexRouter := router.Group("/")
	// 2.绑定路由规则，执行的函数
	// index.GET("", retHelloGinAndMethod)
	// index.POST("", retHelloGinAndMethod)
	indexRouter.Any("", retHelloGinAndMethod)

	userRouter := router.Group("/user")
	userRouter.GET("/:name", handler.UserSave)
	userRouter.GET("/", handler.UserSaveNameByQuery)

	return router
}

// func initHTMLConfig(router *gin.Engine) {
//	if mode := gin.Mode(); mode == gin.TestMode {
//		router.LoadHTMLGlob("./../templates/*")
//	} else {
//		router.LoadHTMLGlob("templates/*")
//	}
// }

// retHelloGinAndMethod 封装了request和response
// gin.Context
func retHelloGinAndMethod(context *gin.Context) {
	str := fmt.Sprintf("hello gin " + strings.ToLower(context.Request.Method) + " method")
	context.String(http.StatusOK, str)
}
