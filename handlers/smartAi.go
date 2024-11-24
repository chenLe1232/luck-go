package handlers

import (
	"log"
	"net/http"

	commitDao "github.com/chenLe1232/luck-go/daos/commit"
	"github.com/chenLe1232/luck-go/dtos/answer"
	answerService "github.com/chenLe1232/luck-go/services/answer"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SmartAiHandler struct {
	db *gorm.DB
}

// 构造函数
func NewSmartAiHandler(db *gorm.DB) *SmartAiHandler {
	return &SmartAiHandler{db: db}
}

//! 智能 AI

// ! 根据 commitID 查询所有相关联的回答
func (h *SmartAiHandler) QueryAnswerByCommitID(c *gin.Context) {
	var request answer.QueryAnswerByCommitIDRequest
	log.Println("request", request)
	//! 修改为使用 Query 参数绑定
	commitID := c.Query("commit_id")
	if commitID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "commit_id is required"})
		return
	}
	request.CommitID = commitID

	//! 调用服务层
	queryRequest := answer.QueryAnswerByCommitIDRequest{
		CommitID: request.CommitID,
	}
	//! 这里其实已经没有必要在 dao 里面进行 db的 初始化行为操作了
	answerService := answerService.NewAnswerService(commitDao.NewCommitDAO(h.db))
	response, err := answerService.QueryAnswerByCommitID(queryRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

// ! 根据问题 返回所有相关的问题
func (h *SmartAiHandler) QueryAnswerByQuestion(c *gin.Context) {
	var request answer.QueryAnswerByQuestionAndCommitIDRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//! 调用服务层
	answerService := answerService.NewAnswerService(commitDao.NewCommitDAO(h.db))
	response, err := answerService.QueryAnswerByQuestion(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
