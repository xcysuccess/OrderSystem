package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserSave 保存User
func UserSave(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "用户"+username+"已经保存")
}

// UserSaveNameByQuery 通过query方法进行后去参数
func UserSaveNameByQuery(context *gin.Context) {
	username := context.Query("name")
	age := context.DefaultQuery("age", "20")
	context.String(http.StatusOK, "用户"+username+",年龄"+age+"已经保存")
}
