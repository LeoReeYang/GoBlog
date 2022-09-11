package v1

import (
	"net/http"
	"strconv"

	"github.com/LeoReeYang/GoBlog/model"
	"github.com/LeoReeYang/GoBlog/utils/errormsg"
	"github.com/gin-gonic/gin"
)

// EditUser
func EditUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	var user model.User
	ctx.ShouldBindJSON(&user)

	code := model.EditUser(id, &user)
	ctx.JSON(http.StatusOK, gin.H{
		"status": errormsg.ErrMsg(code),
		"user":   user,
	})
}

// AddUser
func AddUser(ctx *gin.Context) {
	var user model.User
	ctx.ShouldBindJSON(&user)

	code := model.AddUser(&user)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errormsg.ErrMsg(code),
		"user":   user,
	})
}

// DeleteUser
func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	model.DeleteUser(id)
}

func GetUsers(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))

	data, total := model.GetUsers(pageSize, pageNum)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errormsg.SUCCESS,
		"data":   data,
		"total":  total,
	})
}

func GetUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))

	user, code := model.GetUser(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errormsg.ErrMsg(code),
		"data":   user,
	})
}

func EditPassword(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	var user model.User
	ctx.ShouldBindJSON(&user)

	code := model.EditPassword(id, &user)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   user,
	})
}
