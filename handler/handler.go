package handler

import (
	"gorm.io/gorm"
	"simple_project/Task_Reminder/handler/task"
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

func AddTask(uid *int, t *models.TaskInfo) error {
	h := new(task.AddTaskHandler)
	h.UserInfo = &models.UserInfo{
		Model: gorm.Model{
			ID: uint(utils.FromPtr(uid)),
		},
	}
	h.TaskInfo = t
	h.Handler()
	return h.Error
}

func ListTask(t *models.TaskListRule) (taskInfos *[]models.TaskInfo, err error) {
	h := new(task.ListTaskHandler)
	h.TaskListRule = t
	h.Handler()
	return h.TaskInfos, h.Error
}

func EditTask(uid *int, t *models.TaskInfo) error {
	h := new(task.EditTaskHandler)
	h.UserInfo = &models.UserInfo{
		Model: gorm.Model{
			ID: uint(utils.FromPtr(uid)),
		},
	}
	h.TaskInfo = t
	h.Handler()
	return h.Error
}

func DeleteTask(uid *int, t *models.TaskInfo) error {
	h := new(task.DeleteTaskHandler)
	h.UserInfo = &models.UserInfo{
		Model: gorm.Model{
			ID: uint(utils.FromPtr(uid)),
		},
	}
	h.TaskInfo = t
	h.Handler()
	return h.Error
}
