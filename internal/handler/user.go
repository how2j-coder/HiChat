package handler

import (
	"com/chat/service/internal/dao"
	"com/chat/service/internal/database"
	"com/chat/service/internal/ecode"
	"com/chat/service/internal/model"
	"com/chat/service/internal/types"
	"com/chat/service/pkg/gin/middleware"
	"com/chat/service/pkg/gin/response"
	"com/chat/service/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

var _ UserHandler = (*userHandler)(nil)

type UserHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	GetByID(c *gin.Context)
}

type userHandler struct {
	iDao dao.UserDao
}

func NewUserHandler() UserHandler {
	return &userHandler{
		iDao: dao.NewTeachDao(database.GetDB() ),
	}
}

func (u *userHandler) Create(c *gin.Context)  {
	form := &types.CreateUserReq{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	user := &model.User{}
	err = copier.Copy(user, form)
	if err != nil {
		response.Error(c, ecode.ErrCreateUser)
		return
	}

	ctx := middleware.WrapCtx(c)
	err = u.iDao.Create(ctx, user)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c, gin.H{"id": user.ID})
}
func (u *userHandler) List(c *gin.Context)    {}
func (u *userHandler) GetByID(c *gin.Context) {}
