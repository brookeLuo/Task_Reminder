package controllers

import (
	"github.com/gin-gonic/gin"
	"simple_project/Task_Reminder/handler"
	"simple_project/Task_Reminder/models"
	"simple_project/Task_Reminder/utils"
)

func UserLoginService(c *gin.Context) {
	userInfo := new(models.UserInfo)
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		utils.HttpFailedResponse(c, err, "json to struct failed")
		return
	}
	info, err := handler.UserLogin(userInfo)
	if err != nil {
		utils.HttpFailedResponse(c, err, err.Error())
		return
	}

	token, err := utils.GenerateToken(info.ID)
	if err != nil {
		utils.HttpFailedResponse(c, err, "GenerateToken failed")
		return
	}

	utils.HttpLoginSuccessResponse(c, "login success", token)
	return
}
