package test

import (
	"net/http"
	"net/http/httptest"
	"ordersystem/router"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var gRouter *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	gRouter = router.SetupRouter()
}

// http://localhost:8000
func TestIndexRouter(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	gRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "hello gin get method", "返回的HTML页面中应该包含 hello gin get method")
}
