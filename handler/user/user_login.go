package user

import (
	"errors"
	"simple_project/Task_Reminder/dao"
	"simple_project/Task_Reminder/models"
	"simple_project/Task_Reminder/utils"
)

type LoginHandler struct {
	*models.Resp
	UserInfo *models.UserInfo
}

func (l *LoginHandler) Handler() {
	if l.load(); l.Resp != nil {
		return
	}

	//name 是否存在
	if user, _ := dao.GetUser(map[string]interface{}{
		"user_name": utils.FromPtr(l.UserInfo.UserName),
	}); user != nil {
		if utils.FromPtr(l.UserInfo.Password) == utils.FromPtr(user.Password) {
			l.Resp = &models.Resp{
				Code: utils.ToPtr(0),
				Msg:  utils.ToPtr("login success"),
			}
			l.UserInfo = user
			return
		}
	}

	//email 是否存在
	if user, _ := dao.GetUser(map[string]interface{}{
		"email": utils.FromPtr(l.UserInfo.Email),
	}); user != nil {
		if utils.FromPtr(l.UserInfo.Password) == utils.FromPtr(user.Password) {
			l.Resp = &models.Resp{
				Code: utils.ToPtr(0),
				Msg:  utils.ToPtr("login success"),
			}
			l.UserInfo = user
			return
		}
	}

	l.Resp = &models.Resp{
		Code:  utils.ToPtr(400),
		Msg:   utils.ToPtr("username or email is not exit"),
		Error: errors.New("username or email is not exit"),
	}
	return

}

func (l *LoginHandler) load() {
	if l.UserInfo == nil || (l.UserInfo.UserName == nil && l.UserInfo.Email == nil) {
		l.Resp = &models.Resp{
			Code:  utils.ToPtr(400),
			Msg:   utils.ToPtr("userInfo is invalided"),
			Error: errors.New("userInfo is invalided"),
		}
		return
	}

	//load success
	return
}
