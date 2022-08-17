// Package router TODO
package router

import (
	"net/http"
	"ordersystem/common"
	"ordersystem/middleware"
	"ordersystem/service"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化路由
func SetupRouter() *gin.Engine {
	// router := gin.Default()
	router := gin.New()
	router.Use(middleware.Logger(), gin.Recovery())
	// 1.配置网页
	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob(common.GetCurrentAbPath() + "/templates/*")
	}
	router.Static("/statics", common.GetCurrentAbPath()+"/statics")
	// /Users/tomxiang/github/ordersystem/favicon.ico
	router.StaticFile("/favicon.ico", common.GetCurrentAbPath()+"/favicon.ico")

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
		userRouter.GET("/profile/", middleware.Auth(), service.UserProfile)
		userRouter.POST("/update", middleware.Auth(), service.UpdateUserProfile)
		userRouter.POST("/register", service.UserRegister)
		userRouter.POST("/login", service.UserLogin)
	}
	return router
}
