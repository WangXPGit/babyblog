package routes

import (
	v1 "babyblog/api/v1"
	"babyblog/middleware"
	"babyblog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	router := gin.New()
	router.Use(middleware.Logger())
	router.Use(gin.Recovery())

	// 鉴权组
	auth := router.Group("/api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块 路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		// 分类模块 路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)

		// 文章模块 路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)

		// 上传
		auth.POST("upload", v1.Upload)
	}

	// 无需鉴权组
	r := router.Group("/api/v1")
	{
		// 用户模块 路由接口
		r.POST("user/add", v1.AddUser)
		r.GET("user/list", v1.GetUsers)

		// 分类模块 路由接口
		r.GET("category/list", v1.GetCategories)

		// 文章模块 路由接口
		r.GET("article/list", v1.GetArticles)
		r.GET("article/info/:id", v1.GetCategory)
		r.GET("categoryarticle/list", v1.GetCategoryArticle)
		r.POST("login", v1.Login)
	}

	router.Run(utils.HttpPort)
}
