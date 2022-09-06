package v1

import (
	"net/http"
	"strconv"

	"github.com/LeoReeYang/GoBlog/model"
	"github.com/LeoReeYang/GoBlog/utils/error"
	"github.com/gin-gonic/gin"
)

// UserExit
func UserExit(ctx *gin.Context) {

}

// EditUser
func EditUser(ctx *gin.Context) {

}

// AddUser
func AddUser(ctx *gin.Context) {
	var user model.User

	ctx.ShouldBindJSON(&user)

	code := model.AddUser(&user)

	ctx.JSON(http.StatusOK, gin.H{
		"status": error.ErrMsg(code),
		"user":   user,
	})

}

// DeleteUser
func DeleteUser(ctx *gin.Context) {

}

func GetUsers(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))

	strconv.Atoi(ctx.Query("pageSize"))

	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	data, total := model.GetUsers(pageSize, pageNum)

	ctx.JSON(http.StatusOK, gin.H{
		"status": error.SUCCESS,
		"data":   data,
		"total":  total,
	})
}
