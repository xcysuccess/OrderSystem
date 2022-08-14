package test

import (
	"net/http"
	"net/http/httptest"
	"ordersystem/router"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

var user_router *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	user_router = router.SetupRouter()
}

// http://localhost:8000/user/name=sally
func TestUserSave(t *testing.T) {
	username := "lisi"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/user/"+username, nil)
	user_router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户"+username+"已经保存", w.Body.String())
}

// http://localhost:8000/user/?name=tom&age=33
func TestUserSaveNameAge(t *testing.T) {
	username := "lisi"
	age := 18
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/user/?name="+username+"&age="+strconv.Itoa(age), nil)
	user_router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户"+username+",年龄"+strconv.Itoa(age)+"已经保存", w.Body.String())
}
