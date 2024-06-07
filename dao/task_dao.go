package dao

import (
	"fmt"
	"gorm.io/gorm"
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

func DeleteTask(taskID uint) error {
	result := DB.Where("id = ?", taskID).Delete(&models.TaskInfo{})

	if result.Error != nil {
		return result.Error
	}

	utils.LoggerInfo("delete task success")
	return nil
}

// UpdateTask updates the TaskInfo record in the database
func UpdateTask(task *models.TaskInfo) error {
	// Check if the ID is valid
	if task.ID == 0 {
		return gorm.ErrRecordNotFound
	}

	// Create a map to hold the non-nil fields to be updated
	updates := make(map[string]interface{})

	if task.TaskName != nil {
		updates["task_name"] = task.TaskName
	}
	if task.Status != nil {
		updates["status"] = task.Status
	}
	if task.IsRepeat != nil {
		updates["is_repeat"] = task.IsRepeat
	}
	if task.RepeatRule != nil {
		updates["repeat_rule"] = task.RepeatRule
	}
	if task.RepeatTime != nil {
		updates["repeat_time"] = task.RepeatTime
	}
	if task.TaskDescreption != nil {
		updates["task_descreption"] = task.TaskDescreption
	}
	if task.TaskOwner != nil {
		updates["task_owner"] = task.TaskOwner
	}

	// Update the record with non-nil fields
	result := DB.Model(&models.TaskInfo{}).Where("id = ?", task.ID).Updates(updates)

	if result.Error != nil {
		return result.Error
	}

	utils.LoggerInfo("update task info success")
	return nil
}

// GetTask 根据不同的查询条件获取用户信息列表
func GetTask(conditions map[string]interface{}) ([]models.TaskInfo, error) {
	var tasks []models.TaskInfo
	query := DB.Model(&models.TaskInfo{})
	fmt.Println(conditions)

	// 动态添加查询条件
	for key, value := range conditions {
		if key == "task_name" {
			// 使用 LIKE 模糊查询
			str := conditions[key].(string)
			fmt.Println(str)
			query = query.Where(fmt.Sprintf("%s LIKE ?", key), "%"+str+"%")

		} else {
			query = query.Where(fmt.Sprintf("%s = ?", key), value)
		}
	}

	if err := query.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
