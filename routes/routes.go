package routes

import (
	"net/http"

	user_dao "github.com/chenLe1232/luck-go/daos/users"
	handlers "github.com/chenLe1232/luck-go/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// 初始化 DAO
	userDAO := user_dao.NewUserDAO(db)

	// 初始化 Handler
	userHandler := handlers.NewUserHandler(userDAO)

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

		// 用户相关路由
		api.POST("/users", userHandler.CreateUser)
		api.GET("/users/email", userHandler.FindByEmail)
	}
}
