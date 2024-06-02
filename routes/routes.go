package routes

import (
	"github.com/gin-gonic/gin"
	"simple_project/Task_Reminder/controllers"
)

func Router() *gin.Engine {

	r := gin.Default()

	// 用户相关的路由组
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", controllers.UserLoginService)
		userGroup.POST("/register", controllers.UserRegisterService)
	}

	// 任务相关的路由组
	taskGroup := r.Group("/task")
	{
		taskGroup.GET("/list", controllers.TaskList)
		taskGroup.POST("/add", controllers.TaskAdd)
		taskGroup.POST("/edit", controllers.TaskEdit)
		taskGroup.DELETE("/delete", controllers.TaskDelete)
	}

	return r
}
