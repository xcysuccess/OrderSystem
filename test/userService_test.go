package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"ordersystem/router"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

var userRouter *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	userRouter = router.SetupRouter()
}

// http://localhost:8000/user/name=sally
func TestUserSave(t *testing.T) {
	username := "lisi"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/user/"+username, nil)
	userRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户"+username+"已经保存", w.Body.String())
}

// http://localhost:8000/user/?name=tom&age=33
func TestUserSaveNameAge(t *testing.T) {
	username := "lisi"
	age := 18
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/user/?name="+username+"&age="+strconv.Itoa(age), nil)
	userRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户"+username+",年龄"+strconv.Itoa(age)+"已经保存", w.Body.String())
}

func TestUserPost(t *testing.T) {
	value := url.Values{}
	value.Add("email", "xcysuccess@qq.com")
	value.Add("password", "12345678_XXXX")
	value.Add("password-again", "12345678_XXXX")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/register", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	userRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserPostFormEmailErrorAndPasswordError(t *testing.T) {
	value := url.Values{}
	value.Add("email", "333")
	value.Add("password", "1234")
	value.Add("password-again", "qwer")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/register", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	userRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
