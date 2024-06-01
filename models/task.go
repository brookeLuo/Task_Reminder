package models

import "gorm.io/gorm"

type TaskInfo struct {
	gorm.Model
	TaskName        string `json:"taskName"`
	EnableTIme      string `json:"enableTIme"`
	IsRepeat        int    `json:"IsRepeat"`
	IsRepeatHours   int    `json:"IsRepeatHours"`
	IsRepeatDay     int    `json:"IsRepeatDay"`
	IsRepeatWeek    int    `json:"IsRepeatWeek"`
	IsRepeatMonth   int    `json:"IsRepeatMonth"`
	TaskDescreption string `json:"taskDescreption"`
}
