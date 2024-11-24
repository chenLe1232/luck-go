package main

import (
	"fmt"
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
	//! 配置日志 logs/main.log // 如果不存在就创建这个目录与文件
	os.MkdirAll("logs", os.ModePerm)
	f, err := os.Create("logs/main.log")
	if err != nil {
		log.Fatalf("创建日志文件失败: %v", err)
	}
	defer f.Close()
	// 设置日志输出到文件
	log.SetOutput(f)
	// 数据库连接配置
	dsn := config.GetDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	autoEnv := os.Getenv("AUTO_MIGRATE")
	fmt.Printf("当前 AUTO_MIGRATE 环境变量值: %q\n", autoEnv)
	// 只在特定条件下运行迁移
	if autoEnv == "true" {
		migrations.AutoMigrate(db)
	}

	r := gin.Default()
	routes.SetupRoutes(r, db)
	err = r.Run(":7001")
	if err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
