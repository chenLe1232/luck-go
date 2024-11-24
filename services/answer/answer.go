package answerService

import (
	"errors"

	commitDao "github.com/chenLe1232/luck-go/daos/commit"
	"github.com/chenLe1232/luck-go/dtos/answer"
)

type AnswerService struct {
	commitDao *commitDao.CommitDAO
}

// 构造函数
func NewAnswerService(commitDao *commitDao.CommitDAO) *AnswerService {
	return &AnswerService{
		commitDao: commitDao,
	}
}

func (s *AnswerService) FirstAnswer(request answer.FirstAnswerRequest) {

}

// QueryAnswerByCommitID 查询回答
func (s *AnswerService) QueryAnswerByCommitID(request answer.QueryAnswerByCommitIDRequest) (*answer.QueryAnswerResponse, error) {
	// 参数校验
	if request.CommitID == "" {
		return nil, errors.New("commitID 不能为空")
	}

	// 调用 DAO 层查询
	commits, err := s.commitDao.CommitQueryAllCommitID(request.CommitID)
	if err != nil {
		return nil, err
	}

	// 转换为 DTO 响应
	response := &answer.QueryAnswerResponse{
		Commits: commits, // 注意：实际使用时可能需要将 models.Commit 转换为对应的 DTO
	}

	return response, nil
}

// QueryAnswerByQuestion 根据问题查询回答
func (s *AnswerService) QueryAnswerByQuestion(request answer.QueryAnswerByQuestionAndCommitIDRequest) (*answer.QueryAnswerByQuestionRequest, error) {
	// 参数校验
	if request.Question == "" || request.CommitID == "" {
		return nil, errors.New("question 和 commitID 不能为空")
	}

	// 调用 DAO 层查询
	commits, err := s.commitDao.CommitQuery(request.Question, request.CommitID)
	if err != nil {
		return nil, err
	}

	// 转换为正确的返回类型
	response := &answer.QueryAnswerByQuestionRequest{
		Commits: commits,
	}

	return response, nil
}
