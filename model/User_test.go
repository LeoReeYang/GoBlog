package model

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	// _ "github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/assert"
)

func TestScryptPassword(t *testing.T) {
	hashpass := ScryptPassword("123456")

	t.Log(hashpass)
	if hashpass == "" {
		t.Error()
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("api/v1/user/:id", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func TestGetUser(t *testing.T) {
	r := setupRouter()
	r.Run(":3030")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "api/v1/user/:id", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
	// assert.Equal(t, "pong", w.Body.String())

}
