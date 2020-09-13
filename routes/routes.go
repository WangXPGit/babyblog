package routes

import (
	"babyblog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		// 用户模块接口
		router.POST("user/add", router.AddUser)
		router.GET("user/list", router.GetUsers)
		router.PUT("user/:id", router.EditUser)
		router.DELETE("user/:id", router.DeleteUser)

		// 文章模块接口
		//

		// example
		// router.GET("hello", func(c *gin.Context) {
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"msg": "ok",
		// 	})
		// })
	}

	r.Run(utils.HttpPort)
}
