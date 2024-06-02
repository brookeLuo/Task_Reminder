package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple_project/Task_Reminder/handler"
	"simple_project/Task_Reminder/models"
	"simple_project/Task_Reminder/utils"
)

func TaskAddService(c *gin.Context) {
	token := c.GetHeader("token")

	userId, err := utils.ExtractUserIDFromToken(token)
	if err != nil {
		HttpAddTaskFailedResponse(c, err, "ExtractUserIDFromToken Failed")
		return
	}

	task := new(models.TaskRequest)
	if err := c.ShouldBindJSON(&task); err != nil {
		HttpAddTaskFailedResponse(c, err, "json to struct failed")
		return
	}

	err = handler.AddTask(utils.ToPtr(userId), task.TaskInfo)
	if err != nil {
		HttpAddTaskFailedResponse(c, err, "add task handler failed")
		return
	}

	HttpAddTaskSuccessResponse(c, "add task success", nil)
	return
}

func HttpAddTaskFailedResponse(c *gin.Context, err error, msg string) {
	utils.FailOnError(err, "add task request failed")
	c.JSON(http.StatusOK, gin.H{
		"code": 401,
		"msg":  msg,
		"data": nil,
	})
}

func HttpAddTaskSuccessResponse(c *gin.Context, msg string, obj interface{}) {
	utils.LoggerInfo(msg)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
		"data": obj,
	})
}
