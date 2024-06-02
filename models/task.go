package models

import "gorm.io/gorm"

type TaskInfo struct {
	gorm.Model
	TaskName        *string `json:"taskName"`
	Status          *int    `json:"Status"`
	IsRepeat        *bool   `json:"IsRepeat"`
	RepeatRule      *int    `json:"repeatRule"`
	RepeatTime      *int    `json:"repeatTime"`
	TaskDescreption *string `json:"taskDescreption"`
	TaskOwner       *string `json:"taskOwner"`
}

type TaskListRule struct {
	OperatorID *int
	TaskNames  *string
}

// TaskRequest 定义 TaskRequest 结构体
type TaskRequest struct {
	TaskInfo *TaskInfo `json:"taskinfo"`
}
