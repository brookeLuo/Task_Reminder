package user

import (
	"errors"
	"simple_project/Task_Reminder/dao"
	"simple_project/Task_Reminder/models"
	"simple_project/Task_Reminder/utils"
)

type UserInfoHandler struct {
	*models.Resp
	UserInfo *models.UserInfo
}

func (u *UserInfoHandler) Handler() {
	if u.UserInfo == nil || u.UserInfo.ID == 0 {
		u.Resp = &models.Resp{
			Code:  utils.ToPtr(400),
			Msg:   utils.ToPtr("UserId is Invalid"),
			Error: errors.New("UserId is Invalid"),
		}
		return
	}

	user, err := dao.GetUser(map[string]interface{}{
		"id": u.UserInfo.ID,
	})
	if err != nil || user == nil {
		u.Resp = &models.Resp{
			Code:  utils.ToPtr(400),
			Msg:   utils.ToPtr("UserId is Invalid"),
			Error: errors.New("UserId is Invalid"),
		}
		return
	}
	u.UserInfo = user
	u.Resp = &models.Resp{
		Code: utils.ToPtr(0),
		Msg:  utils.ToPtr("userinfo success"),
	}
	return
}
