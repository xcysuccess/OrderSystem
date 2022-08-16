// Package router TODO
package router

import (
	"net/http"
	"ordersystem/common"
	"ordersystem/service"
	"path/filepath"

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
	// /Users/tomxiang/github/ordersystem/frontdev/images/favicon.ico
	router.StaticFile("/favicon.ico", common.GetCurrentAbPath()+"/frontdev/images/favicon.ico")
	// avatar
	router.StaticFS("/avatar", http.Dir(filepath.Join(common.RootPath(), "avatar")))

	// 2.创建路由index
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
		userRouter.POST("/update", service.UpdateUserProfile)
		userRouter.POST("/register", service.UserRegister)
		userRouter.POST("/login", service.UserLogin)
	}
	return router
}
