package migrations

import (
	"log"
	"os"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	// 检查环境变量
	if os.Getenv("AUTO_MIGRATE") != "true" {
		return
	}

	log.Println("开始数据库迁移...")

	if err := CreateUsersTable(db); err != nil {
		log.Printf("用户表迁移失败: %v", err)
	}

	log.Println("数据库迁移完成")
}
