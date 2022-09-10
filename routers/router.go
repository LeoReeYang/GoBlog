package routers

// "GoBLog/utils"

import (
	v1 "github.com/LeoReeYang/GoBlog/api/v1"
	"github.com/LeoReeYang/GoBlog/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	router := r.Group("api/v1")
	{
		// AddUser
		router.POST("user/add", v1.AddUser)
		// DeleteUser
		router.DELETE("user/delete", v1.DeleteUser)
		// GetUsers
		router.GET("users", v1.GetUsers)
		// EditUser
		router.PUT("user/:id", v1.EditUser)
		// GetUser
		router.GET("user", v1.GetUser)

		// AddCategory
		router.POST("category/add", v1.AddCategory)
		// DeleteCategory
		router.DELETE("category/delete", v1.DeleteCategory)
		// GetCategorys
		router.GET("categorys", v1.GetCategoryList)
		// GetCategory
		router.GET("category", v1.GetCategory)
		// EditCategory
		router.POST("category/edit", v1.EditCategory)

		// AddArticle
		router.POST("article/add", v1.AddArticle)
		// DeleteArticle
		router.DELETE("article/delete", v1.DeleteArticle)
		// GetArticles
		router.GET("articles", v1.GetSameCategoryArticleList)
		// GetArticle
		router.GET("article", v1.GetArticle)
		// EditAticle
		router.PUT("article/edit", v1.EditArticle)
	}

	// routerUser := r.Group("api/v1/user")
	// {
	// 	// AddUser
	// 	routerUser.POST("add", v1.AddUser)
	// 	// DeleteUser
	// 	routerUser.DELETE("delete", v1.DeleteUser)
	// 	// GetUsers
	// 	routerUser.GET("users", v1.GetUsers)
	// 	// EditUser
	// 	routerUser.PUT(":id", v1.EditUser)
	// 	// GetUser
	// 	routerUser.GET("user", v1.GetUser)
	// }

	r.Run(utils.HttpPort)
}
