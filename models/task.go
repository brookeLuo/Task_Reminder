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
	TaskName        *string     `json:"taskName" binding:"required"`
	Status          *Status     `json:"Status" binding:"status"`
	IsRepeat        *bool       `json:"IsRepeat" binding:"required"`
	RepeatRule      *RepeatRule `json:"repeatRule" binding:"repeat_rule"`
	RepeatTime      *int        `json:"repeatTime" binding:"required"`
	TaskDescreption *string     `json:"taskDescreption" binding:"required"`
	TaskOwner       *string     `json:"taskOwner"`
}

type TaskListRule struct {
	OperatorID *int
	TaskNames  *string
}

// TaskRequest 定义 TaskRequest 结构体
type TaskRequest struct {
	TaskInfo *TaskInfo `json:"taskinfo" binding:"required"`
}

type ListTaskRequest struct {
	TaskNames *string `json:"taskname" binding:"required"`
}

type DeleteTaskRequest struct {
	TaskInfo *DeleteTaskInfo `json:"taskinfo" binding:"required"`
}

type DeleteTaskInfo struct {
	ID uint `json:"id" binding:"required"`
}
