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

func TaskAddService(c *gin.Context) {
	token := c.GetHeader("token")

	userId, err := utils.ExtractUserIDFromToken(token)
	if err != nil {
		httpAddTaskFailedResponse(c, err, "ExtractUserIDFromToken Failed")
		return
	}

	// Register custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("repeat_rule", IsValidRepeatRule)
		v.RegisterValidation("status", IsValidStatus)
	}

	task := new(models.TaskRequest)
	if err := c.ShouldBindJSON(&task); err != nil {
		httpAddTaskFailedResponse(c, err, "json to struct failed")
		return
	}

	err = handler.AddTask(utils.ToPtr(userId), task.TaskInfo)
	if err != nil {
		httpAddTaskFailedResponse(c, err, err.Error())
		return
	}

	httpAddTaskSuccessResponse(c, "add task success", nil)
	return
}

// IsValidRepeatRule is a custom validator for RepeatRule
func IsValidRepeatRule(fl validator.FieldLevel) bool {
	rule := fl.Field().Interface().(models.RepeatRule)
	switch rule {
	case models.RepeateInit, models.RepeateByDay, models.RepeateByWeek, models.RepeateByMouth:
		return true
	}
	return false
}

// IsValidStatus is a custom validator for RepeatRule
func IsValidStatus(fl validator.FieldLevel) bool {
	status := fl.Field().Interface().(models.Status)
	switch status {
	case models.IsValidStatus, models.IsInValidStatus:
		return true
	}
	return false
}

func httpAddTaskFailedResponse(c *gin.Context, err error, msg string) {
	utils.FailOnError(err, "add task request failed")
	c.JSON(http.StatusOK, gin.H{
		"code": 401,
		"msg":  msg,
		"data": nil,
	})
}

func httpAddTaskSuccessResponse(c *gin.Context, msg string, obj interface{}) {
	utils.LoggerInfo(msg)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
		"data": obj,
	})
}
