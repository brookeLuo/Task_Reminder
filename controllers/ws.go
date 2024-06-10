package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"simple_project/Task_Reminder/models"
	"simple_project/Task_Reminder/utils"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有源的连接，生产环境下需要根据实际情况进行修改
	},
}

func WsHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade: ", err)
		return
	}
	defer conn.Close()

	for {
		// 读取消息
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error: ", err)
			break
		}
		log.Printf("Received: %s", message)

		msg := []models.TaskInfo{
			{
				TaskOwner: utils.ToPtr("xxx"),
			},
			{
				TaskOwner: utils.ToPtr("xxx"),
			},
			{
				TaskOwner: utils.ToPtr("xxx"),
			},
			{
				TaskOwner: utils.ToPtr("xxx"),
			},
			{
				TaskOwner: utils.ToPtr("xxx"),
			},
			{
				TaskOwner: utils.ToPtr("xxx"),
			},
			{
				TaskOwner: utils.ToPtr("xxx"),
			}, {
				TaskOwner: utils.ToPtr("xxx"),
			}, {
				TaskOwner: utils.ToPtr("xxx"),
			},
			{
				TaskOwner: utils.ToPtr("xxx"),
			},
		}
		for _, value := range msg {
			// 将消息转换为 JSON 格式
			data, err := json.Marshal(value)
			if err != nil {
				log.Println("JSON marshaling error: ", err)
				break
			}

			// 发送 JSON 数据给客户端
			if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println("Write error: ", err)
				break
			}
		}

		// 发送消息
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Println("Write error: ", err)
			break
		}
	}
}
