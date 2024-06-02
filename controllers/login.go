package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple_project/Task_Reminder/handler"
	"simple_project/Task_Reminder/models"
	"simple_project/Task_Reminder/utils"
)

func UserLoginService(c *gin.Context) {
	userInfo := new(models.UserInfo)
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		HttpLoginFailedResponse(c, err, "json to struct failed")
		return
	}
	info, err := handler.UserLogin(userInfo)
	if err != nil {
		HttpLoginFailedResponse(c, err, err.Error())
		return
	}

	token, err := utils.GenerateToken(info.ID)
	if err != nil {
		HttpFailedResponse(c, err, "GenerateToken failed")
		return
	}

	HttpLoginSuccessResponse(c, "login success", token)
	return
}

func HttpLoginSuccessResponse(c *gin.Context, msg string, token string) {
	utils.LoggerInfo(msg)
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   msg,
		"token": token,
	})
}

func HttpLoginFailedResponse(c *gin.Context, err error, msg string) {
	utils.FailOnError(err, "request failed")
	c.JSON(http.StatusOK, gin.H{
		"code": 401,
		"msg":  msg,
		"data": nil,
	})
}
