package migrations

import (
	"log"
	"os"

	"github.com/chenLe1232/luck-go/database/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	// 确保 logs 目录存在
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Printf("无法创建logs目录: %v", err)
		return
	}

	// 配置日志文件
	logFile, err := os.OpenFile("logs/migration.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("无法创建日志文件: %v", err)
		return
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	// 检查环境变量
	if os.Getenv("AUTO_MIGRATE") != "true" {
		return
	}

	log.Println("开始数据库迁移...")

	// 自动迁移所有模型
	if err := db.AutoMigrate(
		&models.User{},
		&models.Commit{},
		&models.Answer{},
	); err != nil {
		log.Printf("数据库迁移失败: %v", err)
		return
	}

	log.Println("数据库迁移完成")
}
