package v1

import (
	"net/http"
	"strconv"

	"github.com/LeoReeYang/GoBlog/model"
	"github.com/LeoReeYang/GoBlog/utils/errormsg"
	"github.com/gin-gonic/gin"
)

func AddCategory(ctx *gin.Context) {
	var category model.Category

	ctx.ShouldBindJSON(&category)

	code := model.AddCategory(&category)
	ctx.JSON(http.StatusOK, gin.H{
		"status":   errormsg.ErrMsg(code),
		"catogory": category,
	})
}

func DeleteCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))

	code := model.DeleteCategory(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
	})
}

func GetCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))

	category, code := model.GetCategory(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errormsg.ErrMsg(code),
		"data":   category,
	})
}

func GetCategoryList(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))

	list, total := model.GetCategoryList(pageSize, pageNum)

	ctx.JSON(http.StatusOK, gin.H{
		"data":  list,
		"total": total,
	})
}

func EditCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))

	var cate model.Category

	ctx.ShouldBindJSON(&cate)

	cate.ID = uint(id)
	code := model.EditCategory(&cate)

	ctx.JSON(http.StatusOK, gin.H{
		"data":   cate,
		"status": code,
	})
}
