package routes

import (
	"github.com/gin-gonic/gin"
	"simple_project/Task_Reminder/controllers"
)

func SetupRouter(r *gin.Engine) {
	// 定义 WebSocket 路由
	r.GET("/ws", controllers.WsHandler)
}
