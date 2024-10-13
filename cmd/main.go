package main

import (
	"log"
	"luck-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	err := r.Run(":7001")
	if err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
