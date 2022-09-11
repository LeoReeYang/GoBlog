package v1

import (
	"net/http"
	"strconv"

	"github.com/LeoReeYang/GoBlog/model"
	"github.com/LeoReeYang/GoBlog/utils/errormsg"
	"github.com/gin-gonic/gin"
)

func AddArticle(ctx *gin.Context) {
	var article model.Article
	ctx.ShouldBindJSON(&article)

	code := model.AddArticle(&article)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    article,
		"message": errormsg.ErrMsg(code),
	})
}

func DeleteArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	code := model.DeleteArticle(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.ErrMsg(code),
	})
}

func GetArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	article, code := model.GetArticle(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    article,
		"message": errormsg.ErrMsg(code),
	})
}

func GetSameCategoryArticleList(ctx *gin.Context) {

	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))
	cid, _ := strconv.Atoi(ctx.Query("cid"))

	list, total, code := model.GetSameCategoryArticleList(cid, pageSize, pageNum)

	ctx.JSON(http.StatusOK, gin.H{
		"data":  list,
		"total": total,
		"code":  code,
	})
}

func EditArticle(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Query("id"))

	var article model.Article
	ctx.ShouldBindJSON(&article)

	code := model.EditArticle(id, &article)

	ctx.JSON(http.StatusOK, gin.H{
		"data": article,
		"code": code,
	})
}
