package task

import (
	"errors"
	"simple_project/Task_Reminder/dao"
	"simple_project/Task_Reminder/models"
	"simple_project/Task_Reminder/utils"
)

type EditTaskHandler struct {
	*models.Resp
	UserInfo *models.UserInfo
	TaskInfo *models.TaskInfo
}

func (a *EditTaskHandler) Handler() {
	if a.load(); a.Resp != nil {
		return
	}
	if err := dao.UpdateTask(a.TaskInfo); err != nil {
		a.Resp = &models.Resp{
			Code:  utils.ToPtr(400),
			Msg:   utils.ToPtr("edit task dao failed"),
			Error: err,
		}
		return
	}

	a.Resp = &models.Resp{
		Code: utils.ToPtr(0),
		Msg:  utils.ToPtr("edit task success"),
	}
	return

}

func (a *EditTaskHandler) load() {
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

	a.UserInfo = user

	//判参数
	if a.TaskInfo == nil || a.TaskInfo.IsRepeat == nil || a.TaskInfo.Status == nil || a.TaskInfo.TaskDescreption == nil || a.TaskInfo.TaskName == nil {
		a.Resp = &models.Resp{
			Code:  utils.ToPtr(400),
			Msg:   utils.ToPtr("task info is invalid"),
			Error: errors.New("task info is invalid"),
		}
		return
	}

	//重复规则
	if utils.FromPtr(a.TaskInfo.IsRepeat) {
		if a.TaskInfo.RepeatRule == nil {
			a.Resp = &models.Resp{
				Code:  utils.ToPtr(400),
				Msg:   utils.ToPtr("RepeatRule info is invalid"),
				Error: errors.New("RepeatRule info is invalid"),
			}
			return
		}
	}

	//判断 token 解析 userinfo 与 task owner 是否一致
	task, err := dao.GetTask(map[string]interface{}{
		"id": a.TaskInfo.ID,
	})
	if task == nil || len(task) != 1 {
		a.Resp = &models.Resp{
			Code:  utils.ToPtr(400),
			Msg:   utils.ToPtr("taskId  is invalid"),
			Error: errors.New("taskId is invalid"),
		}
		return
	}
	if utils.FromPtr(task[0].TaskOwner) != utils.FromPtr(user.UserName) {
		a.Resp = &models.Resp{
			Code:  utils.ToPtr(400),
			Msg:   utils.ToPtr("task is not belong the user"),
			Error: errors.New("task is not belong the user"),
		}
		return
	}

	//success
	a.TaskInfo.TaskOwner = user.UserName
	return

}
