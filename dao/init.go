package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"simple_project/Task_Reminder/config"
	"simple_project/Task_Reminder/models"
)

var DB *gorm.DB

func InitDB() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MySqlUsername,
		config.MySqlPassword,
		config.MysqlHost,
		config.MySqlPort,
		config.MySqlDatabaseName,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	if err := AutoMigrate(); err != nil {
		return nil, err
	}
	return DB, nil
}

// AutoMigrate 自动迁移数据库表
func AutoMigrate() error {
	// 自动迁移模式
	return DB.AutoMigrate(&models.UserInfo{}, &models.TaskInfo{})
}
