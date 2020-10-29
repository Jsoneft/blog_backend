package routes

import (
	v1 "ginblog_backend/api/v1"
	"ginblog_backend/middleware"
	"ginblog_backend/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	auth := r.Group("api/Rv1")
	auth.Use(middleware.JwtToken())
	{
		// User 模块路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		// Article 模块路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
	}
	// Category 模块路由接口

	pub := r.Group("api/Rv1")
	{
		pub.POST("user/add", v1.AddUser)
		pub.GET("users", v1.GetUsers)
		pub.GET("articles", v1.GetArticles)
		pub.GET("article/list/:id", v1.GetCateArt)
		pub.GET("article/info/:id", v1.GetArtInfo)
		pub.GET("categories", v1.GetCategories)
		pub.POST("login",v1.Login)
	}
	_ = r.Run(utils.HttpPort)
}
