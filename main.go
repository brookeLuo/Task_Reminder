package main

import (
	"simple_project/Task_Reminder/dao"
	"simple_project/Task_Reminder/routes"
	"simple_project/Task_Reminder/utils"
)

type A struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var data = make(map[string]A)

func main() {
	//初始化
	initDb, err := dao.InitDB()
	if err != nil {
		utils.FailOnError(err, "init db error")
		panic("Failed to connect to database!")
	}
	s, err := initDb.DB()
	defer s.Close()

	//http
	r := routes.Router()

	r.Run(":8080")
}
