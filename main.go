package main

import (
	"ginblog_backend/model"
	"ginblog_backend/routes"
)

func main()  {
	// 引用数据库
	model.InitDb()
	routes.InitRouter()
}