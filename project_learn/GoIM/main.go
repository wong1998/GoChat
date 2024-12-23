package main

import (
	"GoIM/config"
	"GoIM/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	config.InitDB()

	// 创建Gin引擎
	r := gin.Default()

	// 加载路由
	routes.RegisterRoutes(r)

	// 启动服务
	r.Run(":8080")
}
