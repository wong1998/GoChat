package config

import (
	"GoIM/models"
	//"gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	var err error
	//连接sqlite数据库
	DB, err = gorm.Open(sqlite.Open("chat_app.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate tables
	//自动生成相关的表结构
	err = DB.AutoMigrate(&models.User{}, &models.Message{}, &models.Group{}, &models.GroupMember{}, &models.GroupMessage{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate tables: %v", err)
	}
}
