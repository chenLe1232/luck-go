package dataDao

import (
	"luck-go/database/models"
)

func GetData() ([]models.Data, error) {
	// This is a placeholder implementation. Replace with actual data access logic.
	return []models.Data{
		{ID: 1, Name: "Sample Data 1", Value: "Value 1"},
		{ID: 2, Name: "Sample Data 2", Value: "Value 2"},
	}, nil
}
