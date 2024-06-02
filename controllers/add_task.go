package controllers

import (
	"github.com/gin-gonic/gin"
	"simple_project/Task_Reminder/models"
)

func TaskAdd(c *gin.Context) {
	task := new(models.TaskInfo)
	if err := c.ShouldBindJSON(&task); err != nil {
		HttpFailedResponse(c, err, "json to struct failed")
		return
	}

}
