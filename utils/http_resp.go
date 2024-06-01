package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpFailedResponse(c *gin.Context, err error, msg string) {
	FailOnError(err, "register failed")
	c.JSON(http.StatusOK, gin.H{
		"code": 401,
		"msg":  msg,
		"data": nil,
	})
}

func HttpSuccessResponse(c *gin.Context, msg string, obj interface{}) {
	LoggerInfo(msg)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
		"data": obj,
	})
}
