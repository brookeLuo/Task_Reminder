package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"simple_project/Task_Reminder/handler"
	"simple_project/Task_Reminder/models"
	"simple_project/Task_Reminder/utils"
)

func UserRegisterService(c *gin.Context) {

	userInfo := new(models.UserInfo)
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		utils.HttpFailedResponse(c, err, "json to struct failed")
		return
	}

	fmt.Println(*c)
	fmt.Println(*userInfo)

	err := handler.UserRegister(userInfo)
	if err != nil {
		utils.HttpFailedResponse(c, err, err.Error())
		return
	}

	utils.HttpSuccessResponse(c, "register success", nil)
	return
}
