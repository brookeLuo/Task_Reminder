package models

import "gorm.io/gorm"

const (
	RepeateInit    = RepeatRule(0)
	RepeateByDay   = RepeatRule(100)
	RepeateByWeek  = RepeatRule(200)
	RepeateByMouth = RepeatRule(300)

	IsValidStatus   = Status(1)
	IsInValidStatus = Status(2)
)

type RepeatRule int

type Status int

type TaskInfo struct {
	gorm.Model
	TaskName        *string     `json:"task_name" binding:"required"`
	Status          *Status     `json:"status" binding:"status_jude"`
	IsRepeat        *bool       `json:"is_repeat" binding:"required"`
	RepeatRule      *RepeatRule `json:"repeat_rule" binding:"repeat_rule_jude"`
	RepeatTime      *int        `json:"repeat_time" binding:"required"`
	TaskDescreption *string     `json:"task_descreption" binding:"required"`
	TaskOwner       *string     `json:"task_owner"`
}

type TaskListRule struct {
	OperatorID *int
	TaskNames  *string
}

// TaskRequest 定义 TaskRequest 结构体
type TaskRequest struct {
	TaskInfo *TaskInfo `json:"task_info" binding:"required"`
}

type ListTaskRequest struct {
	TaskNames *string `json:"task_names" binding:"required"`
}

type DeleteTaskRequest struct {
	TaskInfo *DeleteTaskInfo `json:"task_info" binding:"required"`
}

type DeleteTaskInfo struct {
	ID uint `json:"id" binding:"required"`
}
