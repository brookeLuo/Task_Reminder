package main

import (
	"simple_project/Task_Reminder/routes"
)

type A struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var data = make(map[string]A)

func main() {

	r := routes.Router()
	r.Run(":8080")
}
