package service

import (
	"database/sql"
	"log"
	"net/http"
	"ordersystem/common"
	"ordersystem/dao"
	"os"
	"path/filepath"
	"strconv"
	"time"

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
	if user.Password == u.Password {
		log.Println("登录成功", u.Email)
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"email": u.Email,
			"id":    u.Id,
		})
	} else {
		log.Println("登录失败", u.Email)
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
	}
	context.HTML(http.StatusOK, "user_profile.tmpl", gin.H{
		"user": u,
	})

}

// UpdateUserProfile 更新用户信息
func UpdateUserProfile(context *gin.Context) {

	var user dao.UserModel
	if err := context.ShouldBind(&user); err != nil {
		log.Panicln("UpdateUserProfile 绑定错误", err.Error())
	}

	// 1.获取上传表单
	file, err := context.FormFile("avatar-file")
	if err != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err,
		})
		log.Panicln("文件上传错误", err.Error())
	}

	// 2.创建文件夹
	path := common.GetCurrentAbPath() + "/avatar"
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err,
		})
		log.Panicln("无法创建文件夹", err.Error())
	}

	// 3.保存文件
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	log.Println("tomxiang:保存的图片", fileName)
	err = context.SaveUploadedFile(file, filepath.Join(path, fileName))
	if err != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err,
		})
		log.Panicln("无法保存文件", err.Error())
	}

	// 4.更新数据库avatalurl字段
	avatarUrl := "/avatar/" + fileName
	user.Avatar = sql.NullString{String: avatarUrl}
	err = user.Update(user.Id)
	if err != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err,
		})
		log.Panicln("数据无法更新", err.Error())
	}
	context.Redirect(http.StatusMovedPermanently, "/user/profile?id="+strconv.Itoa(user.Id))

}
