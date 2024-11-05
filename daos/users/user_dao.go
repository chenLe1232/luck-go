package user_dao

import (
	"github.com/chenLe1232/luck-go/database/models"

	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

func (dao *UserDAO) Create(user *models.User) error {
	return dao.db.Create(user).Error
}

// ! 根据邮箱查找用户
func (dao *UserDAO) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := dao.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
