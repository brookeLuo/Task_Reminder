package dao

import (
	"fmt"
	"simple_project/Task_Reminder/models"
	"simple_project/Task_Reminder/utils"
)

func AddTask(u *models.TaskInfo) error {

	result := DB.Create(u)

	if result.Error != nil {
		return result.Error
	}

	utils.LoggerInfo("create task info success")
	return nil
}

// GetTask 根据不同的查询条件获取用户信息列表
func GetTask(conditions map[string]interface{}) ([]models.TaskInfo, error) {
	var tasks []models.TaskInfo
	query := DB.Model(&models.TaskInfo{})

	// 动态添加查询条件
	for key, value := range conditions {
		if key == "task_name" {
			// 使用 LIKE 模糊查询
			query = query.Where(fmt.Sprintf("%s LIKE ?", key), fmt.Sprintf("%%%s%%", value))
		} else {
			query = query.Where(fmt.Sprintf("%s = ?", key), value)
		}
	}

	if err := query.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
