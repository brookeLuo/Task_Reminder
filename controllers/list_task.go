package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple_project/Task_Reminder/handler"
	"simple_project/Task_Reminder/models"
	"simple_project/Task_Reminder/utils"
)

func TaskListService(c *gin.Context) {
	token := c.GetHeader("token")

	userId, err := utils.ExtractUserIDFromToken(token)
	if err != nil {
		httpAddTaskFailedResponse(c, err, "ExtractUserIDFromToken Failed")
		return
	}

	task := new(models.ListTaskRequest)
	if err := c.ShouldBindJSON(&task); err != nil || task == nil {
		httpAddTaskFailedResponse(c, err, "json to struct failed")
		return
	}
	rule := new(models.TaskListRule)
	rule.TaskNames = task.TaskNames
	rule.OperatorID = utils.ToPtr(userId)

	infos, err := handler.ListTask(rule)
	if err != nil {
		httpListTaskFailedResponse(c, err, err.Error())
		return
	}

	httpListTaskSuccessResponse(c, "list task success", infos)
}

func httpListTaskFailedResponse(c *gin.Context, err error, msg string) {
	utils.FailOnError(err, "list task request failed")
	c.JSON(http.StatusOK, gin.H{
		"code": 401,
		"msg":  msg,
		"data": nil,
	})
}

func httpListTaskSuccessResponse(c *gin.Context, msg string, obj interface{}) {
	utils.LoggerInfo(msg)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
		"data": obj,
	})
}
