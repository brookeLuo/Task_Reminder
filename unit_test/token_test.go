package unit_test

import (
	"fmt"
	"simple_project/Task_Reminder/utils"
	"testing"
)

func TestToken(t *testing.T) {
	token, err := utils.GenerateToken(1)
	fmt.Println(token, "|||||", err)

	fromToken, err := utils.ExtractUserIDFromToken(token)
	fmt.Println(fromToken, "||||", err)
}
