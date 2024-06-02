package task

import (
	"errors"
	"simple_project/Task_Reminder/dao"
	"simple_project/Task_Reminder/models"
	"simple_project/Task_Reminder/utils"
)

type AddTaskHandler struct {
	*models.Resp
	UserInfo *models.UserInfo
	TaskInfo *models.TaskInfo
}

func (a *AddTaskHandler) Handler() {
	if a.load(); a.Resp != nil {
		return
	}
	if err := dao.AddTask(a.TaskInfo); err != nil {
		a.Resp = &models.Resp{
			Code:  utils.ToPtr(400),
			Msg:   utils.ToPtr("create task dao failed"),
			Error: err,
		}
		return
	}

	a.Resp = &models.Resp{
		Code: utils.ToPtr(0),
		Msg:  utils.ToPtr("add task success"),
	}
	return

}

func (a *AddTaskHandler) load() {
	//插 owner
	if a.UserInfo == nil || a.UserInfo.ID == 0 {
		a.Resp = &models.Resp{
			Code:  utils.ToPtr(400),
			Msg:   utils.ToPtr("user info is invalid"),
			Error: errors.New("user info is invalid"),
		}
		return
	}

	user, err := dao.GetUser(map[string]interface{}{
		"id": a.UserInfo.ID,
	})
	if err != nil {
		a.Resp = &models.Resp{
			Code:  utils.ToPtr(400),
			Msg:   utils.ToPtr("user info is not exit"),
			Error: errors.New("user info is not exit"),
		}
		return
	}

	//判参数
	if a.TaskInfo == nil || a.TaskInfo.IsRepeat == nil || a.TaskInfo.Status == nil || a.TaskInfo.TaskDescreption == nil || a.TaskInfo.TaskName == nil {
		a.Resp = &models.Resp{
			Code:  utils.ToPtr(400),
			Msg:   utils.ToPtr("task info is invalid"),
			Error: errors.New("task info is invalid"),
		}
		return
	}

	//success
	a.TaskInfo.TaskOwner = user.UserName
	return

}
