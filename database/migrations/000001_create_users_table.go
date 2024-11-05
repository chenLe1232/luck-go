package migrations

import (
	"github.com/chenLe1232/luck-go/database/models"

	"gorm.io/gorm"
)

// ! 创建用户表
func CreateUsersTable(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{})
}
