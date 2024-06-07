package user

import (
	"errors"
	"simple_project/Task_Reminder/dao"
	"simple_project/Task_Reminder/models"
	"simple_project/Task_Reminder/utils"
)

type RegisterHandler struct {
	*models.Resp
	UserInfo *models.UserInfo
}

func (r *RegisterHandler) Handler() {
	r.load()
	if r.Resp != nil {
		return
	}

	err := dao.UserRegisterDao(r.UserInfo)
	if err != nil {
		r.Resp = &models.Resp{
			Code:  utils.ToPtr(400),
			Msg:   utils.ToPtr("create user failed"),
			Error: errors.New("create user failed"),
		}
		return
	}

	r.Resp = &models.Resp{
		Code: utils.ToPtr(200),
		Msg:  utils.ToPtr("register success"),
	}
	return
}

func (r *RegisterHandler) load() {
	if r.UserInfo == nil || r.UserInfo.UserName == nil || r.UserInfo.Email == nil {
		r.Resp = &models.Resp{
			Code:  utils.ToPtr(400),
			Msg:   utils.ToPtr("userInfo is invalided"),
			Error: errors.New("userInfo is invalided"),
		}
		return
	}

	//查询 userInfo 是否注册
	user, _ := dao.GetUser(map[string]interface{}{
		"email": utils.FromPtr(r.UserInfo.Email),
	})

	if user != nil {
		r.Resp = &models.Resp{
			Code:  utils.ToPtr(400),
			Msg:   utils.ToPtr("email already exists"),
			Error: errors.New("email already exists"),
		}
		return
	}

	//查询 userInfo 是否注册
	user, _ = dao.GetUser(map[string]interface{}{
		"user_name": utils.FromPtr(r.UserInfo.UserName),
	})

	if user != nil {
		r.Resp = &models.Resp{
			Code:  utils.ToPtr(400),
			Msg:   utils.ToPtr("user_name already exists"),
			Error: errors.New("user_name already exists"),
		}
		return
	}

	//load success
	return
}
