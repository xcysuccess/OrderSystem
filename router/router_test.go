package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/gin/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello World!", w.Body.String())
}
