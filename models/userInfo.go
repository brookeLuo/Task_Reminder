package models

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
