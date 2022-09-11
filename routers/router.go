package routers

// "GoBLog/utils"

import (
	v1 "github.com/LeoReeYang/GoBlog/api/v1"
	"github.com/LeoReeYang/GoBlog/middleware"
	"github.com/LeoReeYang/GoBlog/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	withAuthorization := r.Group("api/v1")
	withAuthorization.Use(middleware.JwtToken())
	{
		// DeleteUser
		withAuthorization.DELETE("user/delete", v1.DeleteUser)
		// EditUser
		withAuthorization.PUT("user/:id", v1.EditUser)
		// EditPassword
		withAuthorization.POST("user/editpassword", v1.EditPassword)

		// AddCategory
		withAuthorization.POST("category/add", v1.AddCategory)
		// DeleteCategory
		withAuthorization.DELETE("category/delete", v1.DeleteCategory)
		// EditCategory
		withAuthorization.POST("category/edit", v1.EditCategory)

		// AddArticle
		withAuthorization.POST("article/add", v1.AddArticle)
		// DeleteArticle
		withAuthorization.DELETE("article/delete", v1.DeleteArticle)
		// EditAticle
		withAuthorization.PUT("article/edit", v1.EditArticle)
	}
	router := r.Group("api/v1")
	{
		// GetUser
		router.GET("user", v1.GetUser)
		// AddUser
		router.POST("user/add", v1.AddUser)
		// GetUsers
		router.GET("users", v1.GetUsers)

		// Login
		router.POST("login", v1.UserLogin)

		// GetCategory
		router.GET("category", v1.GetCategory)
		// GetCategorys
		router.GET("categorys", v1.GetCategoryList)

		// GetArticle
		router.GET("article", v1.GetArticle)
		// GetArticles
		router.GET("articles", v1.GetSameCategoryArticleList)

	}
	r.Run(utils.HttpPort)
}
