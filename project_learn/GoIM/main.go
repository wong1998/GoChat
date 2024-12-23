package main

import (
	"GoIM/config"
	"GoIM/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	config.InitDB()
	gin.SetMode(gin.DebugMode)
	// 创建Gin引擎
	r := gin.Default()
	// 配置 CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // 替换为你的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// 加载路由
	routes.RegisterRoutes(r)

	// 启动服务
	r.Run(":8080")
}
