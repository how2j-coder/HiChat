package handler

import (
	"com/chat/service/internal/dao"
	"com/chat/service/internal/database"
	"com/chat/service/internal/types"
	"github.com/gin-gonic/gin"
)

type MenuHandler interface {
	Create(c *gin.Context)
	UpdateByID(c *gin.Context)
	DeleteByID(c *gin.Context)
}

type menuHandler struct {
	iDao dao.MenuDao
}

var _ MenuHandler = (*menuHandler)(nil)

func NewMenuHandler() MenuHandler {
	return &menuHandler{
		iDao: dao.NewMenuDao(database.GetDB()),
	}
}

func (menuHandler *menuHandler) Create(c *gin.Context) {
	form := &types.CreateMenuReq{}
}
func (menuHandler *menuHandler) UpdateByID(c *gin.Context) {

}
func (menuHandler *menuHandler) DeleteByID(c *gin.Context) {

}