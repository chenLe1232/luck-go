package wallpaper

import (
	"fmt"
)

var keyMap = map[string]string{
	"start":   "5109e05248d5b9368bb559dc",
	"story":   "4fb47a465ba1c65561000028",
	"man":     "4e4d610cdf714d2966000006",
	"machine": "4e4d610cdf714d2966000005",
	"city":    "4fb47a305ba1c60ca5000223",
	"animal":  "4e4d610cdf714d2966000001",
	"design":  "4fb47a195ba1c60ca5000222",
	"emotion": "4ef0a35c0569795756000000",
	"view":    "4fb479f75ba1c65561000027",
	"word":    "5109e04e48d5b9364ae9ac45",
	"game":    "4e4d610cdf714d2966000007",
	"scenery": "4e4d610cdf714d2966000002",
	"comic":   "4e4d610cdf714d2966000003",
	"beauty":  "4e4d610cdf714d2966000000",
}

func GetWallpaperURL(category string, pageSize int, pageStart int) string {
	if _, ok := keyMap[category]; !ok {
		category = "beauty"
	}
	categoryID := keyMap[category]
	return fmt.Sprintf("https://wallpaper.shiyongmodule.com/v1/vertical/category/%s/vertical?appver=2.1.19&limit=%d&skip=%d", categoryID, pageSize, pageStart)
}
