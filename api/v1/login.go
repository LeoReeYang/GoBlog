package v1

import (
	"net/http"

	"github.com/LeoReeYang/GoBlog/middleware"
	"github.com/LeoReeYang/GoBlog/model"
	"github.com/LeoReeYang/GoBlog/utils/errormsg"
	"github.com/gin-gonic/gin"
)

func UserLogin(ctx *gin.Context) {
	var user model.User
	ctx.ShouldBindJSON(&user)

	var token string
	var code int

	errcode := model.UserLogin(user.Name, user.Password)
	if errcode == errormsg.SUCCESS {
		token, code = middleware.SetToken(user.Name)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"user":   user.Name,
		"token":  token,
	})
}

func AdminLogin(ctx *gin.Context) {
	var user model.User
	ctx.ShouldBindJSON(&user)

	var token string
	var code int

	user, errcode := model.AdminLogin(user.Name, user.Password)
	if errcode == errormsg.SUCCESS {
		token, code = middleware.SetToken(user.Name)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"user":   user,
		"token":  token,
	})
}
