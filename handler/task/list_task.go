package task

import (
	"errors"
	"simple_project/Task_Reminder/dao"
	"simple_project/Task_Reminder/models"
	"simple_project/Task_Reminder/utils"
	"strings"
)

type ListTaskHandler struct {
	*models.Resp
	*models.TaskListRule
	TaskInfos *[]models.TaskInfo
}

func (l *ListTaskHandler) Handler() {
	if l.load(); l.Resp != nil {
		return
	}

	var taskInfos []models.TaskInfo

	//查询 username
	user, err := dao.GetUser(map[string]interface{}{
		"id": utils.FromPtr(l.TaskListRule.OperatorID),
	})
	if err != nil || user == nil {
		l.Resp = &models.Resp{
			Code:  utils.ToPtr(400),
			Msg:   utils.ToPtr("operator is invalid"),
			Error: errors.New("operator is invalid"),
		}
		return
	}
	//切分 task name
	taskNames := strings.Split(utils.FromPtr(l.TaskNames), ",")
	for _, value := range taskNames {
		ts, err := dao.GetTask(map[string]interface{}{
			"task_name":  value,
			"task_owner": utils.FromPtr(user.UserName),
		})
		if err != nil {
			l.Resp = &models.Resp{
				Code:  utils.ToPtr(400),
				Msg:   utils.ToPtr(err.Error()),
				Error: err,
			}
			return
		}
		taskInfos = append(taskInfos, ts...)
	}

	l.TaskInfos = utils.ToPtr(taskInfos)
	l.Resp = &models.Resp{
		Code: utils.ToPtr(0),
		Msg:  utils.ToPtr("list task success"),
	}
	return

}

func (l *ListTaskHandler) load() {
	if l.TaskListRule == nil || l.TaskListRule.TaskNames == nil || l.TaskListRule.OperatorID == nil {
		l.Resp = &models.Resp{
			Code:  utils.ToPtr(400),
			Msg:   utils.ToPtr("task list rule is invalid"),
			Error: errors.New("task list rule is invalid"),
		}
		return
	}
}
