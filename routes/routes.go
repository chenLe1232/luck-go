package routes

import (
	handlers "luck-go/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// API 路由组
	api := r.Group("/api")
	{
		// 将 POST 改为 GET
		api.GET("/wallpaper", handlers.GetWallpaper)

		// 将这些路由移到 api 组内
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		api.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello, World!",
			})
		})
	}
}
