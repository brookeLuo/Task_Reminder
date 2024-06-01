package models

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	UserName *string `json:"userName"`
	Password *string `json:"password"`
	Email    *string `json:"email"`
}

type Resp struct {
	Code  *int    `json:"code"`
	Msg   *string `json:"msg"`
	Error error
}
