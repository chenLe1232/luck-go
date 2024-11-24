package answer

import "github.com/chenLe1232/luck-go/database/models"

// ! 首次回答
type FirstAnswerRequest struct {
	Uid string `json:"uid" binding:"required"`
	//！ 问题
	Question string `json:"question" binding:"required"`
	//！ 入参的 commitID
	CommitID string `json:"commit_id" binding:"required"`
}

// ! 根据 commitID 查询所有相关联的回答
type QueryAnswerByCommitIDRequest struct {
	CommitID string `json:"commit_id" binding:"required"`
}

// ! 输入的问题 以及 commitId
type QueryAnswerByQuestionAndCommitIDRequest struct {
	Question string `json:"question" binding:"required"`
	CommitID string `json:"commit_id" binding:"required"`
}

// ! 根据问题查询回答 返回的是 符合条件的所有的 model 的 Commit 结构体 或者是空的结构体
type QueryAnswerByQuestionRequest struct {
	Commits []*models.Commit `json:"commits"`
}

type QueryAnswerResponse struct {
	Commits interface{} // 根据实际的 Commit 结构体类型调整
}
