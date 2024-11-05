package handlers

import (
	"fmt"
	"net/http"

	userDao "github.com/chenLe1232/luck-go/daos/users"
	"github.com/chenLe1232/luck-go/database/models"
	"github.com/chenLe1232/luck-go/dtos/user"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userDAO *userDao.UserDAO
}

func NewUserHandler(userDAO *userDao.UserDAO) *UserHandler {
	return &UserHandler{userDAO: userDAO}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var dto user.CreateUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{
		Username: dto.Username,
		Password: dto.Password, // 注意：实际应用中应该对密码进行加密
		Email:    dto.Email,
	}

	if err := h.userDAO.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": user.ID})
}

// ! 根据邮箱查找用户
func (h *UserHandler) FindByEmail(c *gin.Context) {
	var dto user.FindByEmailDTO
	if err := c.ShouldBindQuery(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "请提供有效的邮箱地址",
			"details": err.Error(),
		})
		return
	}
	//! 打印dto
	fmt.Println("dto", dto)
	user, err := h.userDAO.FindByEmail(dto.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
