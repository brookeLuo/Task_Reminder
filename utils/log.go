package utils

import "fmt"

func FailOnError(err error, msg string) {
	if err != nil {
		fmt.Println("msg =", msg, " err=", err)
	}
}

func LoggerInfo(msg string) {
	fmt.Println("Log Info: ", msg)
}
