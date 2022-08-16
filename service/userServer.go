package service

import (
	"log"
	"net/http"
	"ordersystem/dao"
	"strconv"

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
		context.String(http.StatusBadRequest, "输入的数据不合法")
		log.Println("err -->", err.Error())
	}
	passwordAgain := context.PostForm("password-again")
	if passwordAgain != user.Password {
		context.String(http.StatusBadRequest, "密码校验无效，两次密码不一致")
		log.Panicln("密码校验无效，两次密码不一致")
	}
	id := user.Save()
	log.Println("id is ", id)
	context.Redirect(http.StatusMovedPermanently, "/")
}

// UserLogin 用户登陆
func UserLogin(context *gin.Context) {
	var user dao.UserModel
	if err := context.ShouldBind(&user); err != nil {
		log.Panicln("login 绑定错误", err.Error())
	}
	u := user.QueryByEmail()
	log.Println("u.password ** user.password", u.Password, user.Password)
	if user.Password == u.Password {
		log.Println("登陆成功！")
		// context.HTML(http.StatusOK, "index.tmpl", gin.H{
		//	"email": u.Email,
		// })
		context.Redirect(http.StatusMovedPermanently, "/")
	}
}

// UserProfile 获取用户信息
func UserProfile(context *gin.Context) {
	id := context.Query("id")
	var user dao.UserModel
	i, err := strconv.Atoi(id)
	u, e := user.QueryById(i)
	if e != nil || err != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
	} else {
		context.HTML(http.StatusOK, "user_profile.tmpl", gin.H{
			"user": u,
		})
	}
}
