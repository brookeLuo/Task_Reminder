package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple_project/Task_Reminder/handler"
	"simple_project/Task_Reminder/models"
	"simple_project/Task_Reminder/utils"
)

func UserRegisterService(c *gin.Context) {

	userInfo := new(models.UserInfo)
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		httpFailedResponse(c, err, "json to struct failed")
		return
	}

	fmt.Println(*c)
	fmt.Println(*userInfo)

	err := handler.UserRegister(userInfo)
	if err != nil {
		httpFailedResponse(c, err, err.Error())
		return
	}

	httpSuccessResponse(c, "register success", nil)
	return
}

func httpFailedResponse(c *gin.Context, err error, msg string) {
	utils.FailOnError(err, "request failed")
	c.JSON(http.StatusOK, gin.H{
		"code": 401,
		"msg":  msg,
		"data": nil,
	})
}

func httpSuccessResponse(c *gin.Context, msg string, obj interface{}) {
	utils.LoggerInfo(msg)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
		"data": obj,
	})
}
