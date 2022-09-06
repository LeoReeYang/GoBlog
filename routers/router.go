package routers

// "GoBLog/utils"

import (
	"net/http"

	v1 "github.com/LeoReeYang/GoBlog/api/v1"
	"github.com/LeoReeYang/GoBlog/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	router := r.Group("api/v1")
	{
		// hello test
		router.GET("hello", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "hello,world",
			})
		})

		// AddUser
		router.POST("user/add", v1.AddUser)
		// DeleteUser
		router.DELETE("user/:id", v1.DeleteUser)
		// ExsitUser
		router.GET("users", v1.EditUser)
		//EditUser
		router.PUT("user/:id", v1.EditUser)
	}

	r.Run(utils.HttpPort)
}
