package routes

import (
	"net/http"

	handlers "github.com/chenLe1232/luck-go/handlers"

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

		// Add a new route for the data service API endpoint
		api.GET("/data", handlers.GetData)
		api.POST("/data", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "hello world",
			})
		})
	}
}
