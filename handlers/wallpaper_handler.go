package handlers

import (
	"luck-go/dtos/wallpaper"
	wallpaperService "luck-go/services/wallpaper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWallpaper(c *gin.Context) {
	var req wallpaper.WallpaperRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := wallpaperService.GetWallpaperURL(req.Category, req.PageSize, req.PageStart)
	// 发起get请求url服务数据
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	// 使用 DataFromReader 直接流式传输响应
	c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
}
