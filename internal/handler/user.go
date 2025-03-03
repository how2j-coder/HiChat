package handler

import "github.com/gin-gonic/gin"

var _ UserHandler = (*userHandler)(nil)

type UserHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	GetByID(c *gin.Context)
}

type userHandler struct {
}

func NewUserHandler() UserHandler {
	return &userHandler{}
}

func (u *userHandler) Create(c *gin.Context)  {}
func (u *userHandler) List(c *gin.Context)    {}
func (u *userHandler) GetByID(c *gin.Context) {}
