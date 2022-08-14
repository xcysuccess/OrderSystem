package test

import (
	"net/http"
	"net/http/httptest"
	"ordersystem/router"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var g_router *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	g_router = router.SetupRouter()
}

// http://localhost:8000
func TestIndexGetRouter(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	g_router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin get method", w.Body.String())
}

// http://localhost:8000
func TestIndexPostRouter(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	g_router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin post method", w.Body.String())
}
