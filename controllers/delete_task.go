package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"simple_project/Task_Reminder/handler"
	"simple_project/Task_Reminder/models"
	"simple_project/Task_Reminder/utils"
)

func TaskDeleteService(c *gin.Context) {
	token := c.GetHeader("token")

	userId, err := utils.ExtractUserIDFromToken(token)
	if err != nil {
		httpDeleteTaskFailedResponse(c, err, "ExtractUserIDFromToken Failed")
		return
	}

	task := new(models.DeleteTaskRequest)
	if err := c.ShouldBindJSON(&task); err != nil {
		httpDeleteTaskFailedResponse(c, err, "json to struct failed")
		return
	}

	err = handler.DeleteTask(utils.ToPtr(userId), &models.TaskInfo{
		Model: gorm.Model{
			ID: task.TaskInfo.ID,
		},
	})
	if err != nil {
		httpDeleteTaskFailedResponse(c, err, err.Error())
		return
	}

	httpDeleteTaskSuccessResponse(c, "delete task success", nil)
	return
}

func httpDeleteTaskFailedResponse(c *gin.Context, err error, msg string) {
	utils.FailOnError(err, "delete task request failed")
	c.JSON(http.StatusOK, gin.H{
		"code": 401,
		"msg":  msg,
		"data": nil,
	})
}

func httpDeleteTaskSuccessResponse(c *gin.Context, msg string, obj interface{}) {
	utils.LoggerInfo(msg)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
		"data": obj,
	})
}
