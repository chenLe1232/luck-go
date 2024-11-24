package answerDao

import (
	"github.com/chenLe1232/luck-go/database/models"
	"gorm.io/gorm"
)

type AnswerDAO struct {
	db *gorm.DB
}

// ! 如果用户在表里不存在 则创建
func (dao *AnswerDAO) AnswerCreate(answer *models.Answer) error {
	return dao.db.Create(answer).Error
}
