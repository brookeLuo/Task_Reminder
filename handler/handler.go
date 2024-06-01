package handler

import (
	"simple_project/Task_Reminder/handler/user"
	"simple_project/Task_Reminder/models"
)

func UserRegister(u *models.UserInfo) error {
	h := new(user.RegisterHandler)
	h.UserInfo = u
	h.Handler()
	return h.Error
}
