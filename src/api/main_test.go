package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {

	route := MapRoutes()
	res := PerformRequest("GET", "/ping", "", route)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, `{"message":"pong"}`, res.Body.String())

}

//PerformRequest ...
func PerformRequest(method, target, body string, engine *gin.Engine) *httptest.ResponseRecorder {
	payload := strings.NewReader(body)
	req := httptest.NewRequest(method, target, payload)
	res := httptest.NewRecorder()
	engine.ServeHTTP(res, req)
	return res
}
