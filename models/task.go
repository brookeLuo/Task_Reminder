package models

import "gorm.io/gorm"

type TaskInfo struct {
	gorm.Model
	TaskName        string `json:"taskName"`
	Status          string `json:"Status"`
	IsRepeat        bool   `json:"IsRepeat"`
	RepeatRule      int    `json:"repeatRule"`
	RepeatTime      int    `json:"repeatTime"`
	TaskDescreption string `json:"taskDescreption"`
	TaskOwner       string `json:"taskOwner"`
}
