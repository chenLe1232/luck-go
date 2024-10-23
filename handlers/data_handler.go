package handlers

import (
	"log"
	"luck-go/services/data_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetData(c *gin.Context) {
	log.Printf("Received GET request for data: %+v", c.Request.URL.Query())

	data, err := data_service.FetchData()
	if err != nil {
		log.Printf("Error fetching data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func PostData(c *gin.Context) {
	log.Printf("Received POST request for data: %+v", c.Request.URL.Query())

	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
