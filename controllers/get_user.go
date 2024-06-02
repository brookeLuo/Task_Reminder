package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple_project/Task_Reminder/handler"
	"simple_project/Task_Reminder/utils"
)

func UserInfoService(c *gin.Context) {

	token := c.GetHeader("token")
	
	userId, err := utils.ExtractUserIDFromToken(token)
	if err != nil {
		HttpUserInfoFailedResponse(c, err, "ExtractUserIDFromToken Failed")
		return
	}
	info, err := handler.UserInfo(utils.ToPtr(userId))
	if err != nil {
		HttpUserInfoFailedResponse(c, err, "userInfo failed")
		return
	}

	HttpUserInfoSuccessResponse(c, "userinfo success", info)
	return
}

func HttpUserInfoFailedResponse(c *gin.Context, err error, msg string) {
	utils.FailOnError(err, "userinfo request failed")
	c.JSON(http.StatusOK, gin.H{
		"code": 401,
		"msg":  msg,
		"data": nil,
	})
}

func HttpUserInfoSuccessResponse(c *gin.Context, msg string, obj interface{}) {
	utils.LoggerInfo(msg)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
		"data": obj,
	})
}
