package service

import (
	"log"
	"net/http"
	"ordersystem/dao"

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

// UserRegister 用户注册
func UserRegister(context *gin.Context) {
	var user dao.UserModel
	if err := context.ShouldBind(&user); err != nil {
		log.Println("err -->", err.Error())
		return
	} else {
		log.Println("email", user.Email, "password", user.Password, "password again", user.PasswordAgain)
		context.Redirect(http.StatusMovedPermanently, "/")
	}
	// email := context.PostForm("email")
	// password := context.DefaultPostForm("password", "123456")
	// passwordAgain := context.DefaultPostForm("password-again", "123456")
	// fmt.Println("email", user.Email, "password", user.Password, "password again", user.PasswordAgain)
}
