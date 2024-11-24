package commitDao

import (
	"github.com/chenLe1232/luck-go/database/models"
	"gorm.io/gorm"
)

type CommitDAO struct {
	db *gorm.DB
}

func NewCommitDAO(db *gorm.DB) *CommitDAO {
	return &CommitDAO{db: db}
}

func (dao *CommitDAO) CommitCreate(commit *models.Commit) error {
	return dao.db.Create(commit).Error
}

// ! 查询所有的 commit_id
func (dao *CommitDAO) CommitQueryAllCommitID(commit_id string) ([]*models.Commit, error) {
	var commits []*models.Commit
	err := dao.db.Where("commit_id = ?", commit_id).Find(&commits).Error
	if err != nil {
		return nil, err
	}
	return commits, nil
}

// ! 根据用户 question 关键字查询，返回所有匹配的会话记录
func (dao *CommitDAO) CommitQuery(question string, commit_id string) ([]*models.Commit, error) {
	//! 从数据库中查询
	var commits []*models.Commit
	err := dao.db.Where("question LIKE ?", "%"+question+"%").Find(&commits).Error
	if err != nil {
		return nil, err
	}
	//! 把当前的commit_id 以及 question 创建一条新纪录
	dao.db.Create(&models.Commit{
		CommitID: commit_id,
		Question: question,
	})
	return commits, nil
}
