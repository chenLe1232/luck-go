package wallpaper

type WallpaperRequest struct {
	Category  string `form:"category" binding:"required"`
	PageSize  int    `form:"pageSize,default=10"`
	PageStart int    `form:"pageStart,default=1"`
}

type WallpaperResponse struct {
	URL string `json:"url"`
}
