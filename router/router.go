// Package router TODO
package router

import (
	"ordersystem/common"
	"ordersystem/service"

	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化路由
func SetupRouter() *gin.Engine {
	router := gin.Default()
	// 1.配置网页
	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../frontdev/templates/*")
	} else {
		router.LoadHTMLGlob(common.GetCurrentAbPath() + "/frontdev/templates/*")
	}
	router.Static("/bootstrap", "./frontdev/bootstrap")
	router.StaticFile("/favicon.ico", common.GetCurrentAbPath()+"/frontdev/images/favicon.ico")
	// 2.创建路由
	indexRouter := router.Group("/")
	{
		indexRouter.Any("", service.Index)
	}
	// 3. User
	userRouter := router.Group("/user")
	{
		userRouter.GET("/:name", service.UserSave)
		userRouter.GET("/", service.UserSaveNameByQuery)
		userRouter.GET("/profile/", service.UserProfile)
		userRouter.POST("/register", service.UserRegister)
		userRouter.POST("/login", service.UserLogin)
	}
	return router
}
