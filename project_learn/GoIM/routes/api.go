package routes

import (
	"GoIM/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
		api.GET("/messages/user_his", controllers.GetMessageHistory)
		api.GET("/messages/group_his", controllers.GetGroupMessageHistory)
		// Route to send a message to a specific user
		api.POST("/messages/send_to_user", controllers.SendMessageToUser)
		// Route to send a message to a group
		api.POST("/messages/send_to_group", controllers.SendMessageToGroup)
	}
}
