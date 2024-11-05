package main

import (
	"log"
	"os"

	"github.com/chenLe1232/luck-go/database/migrations"
	config "github.com/chenLe1232/luck-go/pkg/config"
	"github.com/chenLe1232/luck-go/routes"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func main() {
	// 数据库连接配置
	dsn := config.GetDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 只在特定条件下运行迁移
	if os.Getenv("AUTO_MIGRATE") == "true" {
		if err := migrations.CreateUsersTable(db); err != nil {
			log.Printf("警告：数据库迁移失败: %v", err)
		}
	}

	r := gin.Default()
	routes.SetupRoutes(r, db)
	err = r.Run(":7001")
	if err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
