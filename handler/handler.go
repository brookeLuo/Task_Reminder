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

func UserLogin(u *models.UserInfo) (UserInfo *models.UserInfo, err error) {
	h := new(user.LoginHandler)
	h.UserInfo = u
	h.Handler()
	return h.UserInfo, h.Error
}
