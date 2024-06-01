package dao

import (
	"simple_project/Task_Reminder/models"
	"simple_project/Task_Reminder/utils"
)

func UserRegisterDao(u *models.UserInfo) error {

	result := DB.Create(u)

	if result.Error != nil {
		return result.Error
	}

	utils.LoggerInfo("create user info success")
	return nil
}

// GetUser 根据不同的查询条件获取用户信息
func GetUser(conditions map[string]interface{}) (*models.UserInfo, error) {
	var user models.UserInfo
	query := DB.Model(&models.UserInfo{})

	// 动态添加查询条件
	for key, value := range conditions {
		query = query.Where(key+" = ?", value)
	}

	if err := query.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
