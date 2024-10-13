package handlers

import (
	"log"
	"luck-go/dtos/wallpaper"
	wallpaperService "luck-go/services/wallpaper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWallpaper(c *gin.Context) {
	log.Printf("Received GET request for wallpaper: %+v", c.Request.URL.Query())

	var req wallpaper.WallpaperRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Printf("Error binding query parameters: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := wallpaperService.GetWallpaperURL(req.Category, req.PageSize, req.PageStart)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching wallpaper: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch wallpaper"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code from wallpaper service: %d", resp.StatusCode)
		c.JSON(resp.StatusCode, gin.H{"error": "Wallpaper service returned an error"})
		return
	}

	// 使用 DataFromReader 直接流式传输响应
	c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
}
