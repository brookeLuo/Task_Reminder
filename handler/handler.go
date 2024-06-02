package handler

import (
	"gorm.io/gorm"
	"simple_project/Task_Reminder/handler/user"
	"simple_project/Task_Reminder/models"
	"simple_project/Task_Reminder/utils"
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

func UserInfo(userId *int) (UserInfo *models.UserInfo, err error) {
	h := new(user.UserInfoHandler)
	h.UserInfo = &models.UserInfo{
		Model: gorm.Model{
			ID: uint(utils.FromPtr(userId)),
		},
	}
	h.Handler()
	return h.UserInfo, h.Error
}
