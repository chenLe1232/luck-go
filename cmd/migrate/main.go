package main

import (
	"log"

	"github.com/chenLe1232/luck-go/database/migrations"
	config "github.com/chenLe1232/luck-go/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := config.GetDSN()
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 运行迁移
	migrations.AutoMigrate(db)
}
