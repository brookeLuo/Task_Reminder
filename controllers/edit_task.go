package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"simple_project/Task_Reminder/handler"
	"simple_project/Task_Reminder/models"
	"simple_project/Task_Reminder/utils"
)

func TaskEditService(c *gin.Context) {
	token := c.GetHeader("token")

	userId, err := utils.ExtractUserIDFromToken(token)
	if err != nil {
		httpEditTaskFailedResponse(c, err, "ExtractUserIDFromToken Failed")
		return
	}

	// Register custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("repeat_rule", IsValidRepeatRule)
		v.RegisterValidation("status", IsValidStatus)
	}

	task := new(models.TaskRequest)
	if err := c.ShouldBindJSON(&task); err != nil {
		httpEditTaskFailedResponse(c, err, "json to struct failed")
		return
	}

	err = handler.EditTask(utils.ToPtr(userId), task.TaskInfo)
	if err != nil {
		httpEditTaskFailedResponse(c, err, err.Error())
		return
	}

	httpEditTaskSuccessResponse(c, "edit task success", nil)
	return
}

func httpEditTaskFailedResponse(c *gin.Context, err error, msg string) {
	utils.FailOnError(err, "edit task request failed")
	c.JSON(http.StatusOK, gin.H{
		"code": 401,
		"msg":  msg,
		"data": nil,
	})
}

func httpEditTaskSuccessResponse(c *gin.Context, msg string, obj interface{}) {
	utils.LoggerInfo(msg)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
		"data": obj,
	})
}
