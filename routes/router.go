package routes

import (
	"babyblog/api/v1"
	"babyblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	router := gin.Default()

	r := router.Group("/api/v1")
	{
		// 用户模块接口
		r.POST("user/add", v1.AddUser)
		r.GET("user/list", v1.GetUsers)
		r.PUT("user/:id", v1.EditUser)
		r.DELETE("user/:id", v1.DeleteUser)

		// 文章模块接口
		//

		// example
		//v1.GET("hello", func(c *gin.Context) {
		//	c.JSON(http.StatusOK, gin.H{
		//		"msg": "ok",
		//	})
		//})
	}

	router.Run(utils.HttpPort)
}
