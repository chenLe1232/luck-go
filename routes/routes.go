package routes

import (
	handlers "luck-go/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// 现有的路由
	api := r.Group("/api")
	{
		api.POST("/wallpaper", handlers.GetWallpaper)
	}

	// 增加两个新路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})
}
