package routes

import (
	"babyblog/utils"
	"babyblog/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		// 用户模块接口
		v1.POST("user/addnew", v1.AddUser)
		v1.POST("user/add", v1.AddUser)
		v1.GET("user/list", v1.GetUsers)
		v1.PUT("user/:id", v1.EditUser)
		v1.DELETE("user/:id", v1.DeleteUser)

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
