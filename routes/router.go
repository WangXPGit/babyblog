package routes

import (
	v1 "babyblog/api/v1"
	"babyblog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	router := gin.Default()

	r := router.Group("/api/v1")
	{
		// 用户模块 路由接口
		r.POST("user/add", v1.AddUser)
		r.GET("user/list", v1.GetUsers)
		r.PUT("user/:id", v1.EditUser)
		r.DELETE("user/:id", v1.DeleteUser)

		// 分类模块 路由接口
		r.POST("category/add", v1.AddCategory)
		r.GET("category/list", v1.GetCategories)
		r.PUT("category/:id", v1.EditCategory)
		r.DELETE("category/:id", v1.DeleteCategory)

		// 文章模块 路由接口
		r.POST("article/add", v1.AddArticle)
		r.GET("article/list", v1.GetArticles)
		r.GET("article/info/:id", v1.GetCategory)
		r.GET("categoryarticle/list", v1.GetCategoryArticle)
		r.PUT("article/:id", v1.EditArticle)
		r.DELETE("article/:id", v1.DeleteArticle)

		// example
		//v1.GET("hello", func(c *gin.Context) {
		//	c.JSON(http.StatusOK, gin.H{
		//		"msg": "ok",
		//	})
		//})
	}

	router.Run(utils.HttpPort)
}
