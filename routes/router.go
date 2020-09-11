package routes

import (
	v1 "ginblog_backend/api/v1"
	"ginblog_backend/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter()  {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	Rv1 := r.Group("api/Rv1")
	{
	 	// User 模块路由接口
		Rv1.POST("user/add", v1.AddUser)
		Rv1.GET("users", v1.GetUsers)
		Rv1.PUT("user/:id", v1.EditUser)
		Rv1.DELETE("user/:id", v1.DeleteUser)
		// Article 模块路由接口

		// Category 模块路由接口
	}
	r.Run(utils.HttpPort)
}