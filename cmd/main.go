package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 添加新的路由
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "你好,世界!",
		})
	})

	r.POST("/echo", func(c *gin.Context) {
		var json struct {
			Message string `json:"message"`
		}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "收到的消息: " + json.Message,
		})
	})

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
