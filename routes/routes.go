package routes

import (
	"github.com/gin-gonic/gin"
	"simple_project/Task_Reminder/controllers"
)

func Router() *gin.Engine {

	r := gin.Default()

	// 创建一个带有前缀的路由组 /api/v2
	apiV2 := r.Group("/api/v2")
	{
		// 用户相关的路由组
		userGroup := apiV2.Group("/user")
		{
			userGroup.POST("/login", controllers.UserLoginService)
			userGroup.POST("/register", controllers.UserRegisterService)
			userGroup.GET("/userinfo", controllers.UserInfoService)
		}

		// 任务相关的路由组
		taskGroup := apiV2.Group("/task")
		{
			taskGroup.POST("/list", controllers.TaskListService)
			taskGroup.POST("/add", controllers.TaskAddService)
			taskGroup.POST("/edit", controllers.TaskEditService)
			taskGroup.DELETE("/delete", controllers.TaskDeleteService)
		}
	}

	//ws
	SetupRouter(r)

	// 启动服务器
	r.Run(":8080") // 默认监听并服务于 0.0.0.0:8080

	return r
}
