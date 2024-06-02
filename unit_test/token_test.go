package unit_test

import (
	"fmt"
	"simple_project/Task_Reminder/dao"
	"simple_project/Task_Reminder/utils"
	"testing"
)

func TestToken(t *testing.T) {
	token, err := utils.GenerateToken(1)
	fmt.Println(token, "|||||", err)

	fromToken, err := utils.ExtractUserIDFromToken(token)
	fmt.Println(fromToken, "||||", err)
}

func TestListTask(t *testing.T) {
	dao.InitDB()
	task, err := dao.GetTask(map[string]interface{}{
		"task_name":  "以办也联达克",
		"task_owner": "蔡强",
	})

	for _, v := range task {
		fmt.Println(v)
	}
	fmt.Println(err)
}
