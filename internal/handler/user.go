package handler

import (
	"com/chat/service/internal/cache"
	"com/chat/service/internal/dao"
	"com/chat/service/internal/database"
	"com/chat/service/internal/ecode"
	"com/chat/service/internal/handler/common"
	"com/chat/service/internal/model"
	"com/chat/service/internal/types"
	"com/chat/service/pkg/gin/middleware"
	"com/chat/service/pkg/gin/response"
	"com/chat/service/pkg/gocrypto"
	"com/chat/service/pkg/jwt"
	"com/chat/service/pkg/logger"
	"com/chat/service/pkg/srand"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/redis/go-redis/v9"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
	"strconv"
)

var _ UserHandler = (*userHandler)(nil)

type UserHandler interface {
	Create(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
	UpdateByID(c *gin.Context)
	List(c *gin.Context)
	GetByID(c *gin.Context)
}

type userHandler struct {
	iDao dao.UserDao
}

func NewUserHandler() UserHandler {
	return &userHandler{
		iDao: dao.NewUserDao(database.GetDB()),
	}
}

// Create 注册
func (u *userHandler) Create(c *gin.Context) {
	form := &types.CreateUserReq{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	form.Account = srand.String(6)
	form.Password, _ = gocrypto.HashAndSaltPassword(form.Password)

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

// Login 登录
func (u *userHandler) Login(c *gin.Context) {
	cRedis := database.GetRedisClient()
	defer func(cRedis *redis.Client) {
		err := cRedis.Close()
		if err != nil {
			logger.Warn("Redis.Close() error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
			response.Output(c, ecode.InternalServerError.ToHTTPCode())
			return
		}
	}(cRedis)

	uCache := cache.NewUserCache(cRedis)

	loginType := []string{"account", "email"}
	// 获取登录类型
	q := c.Request.URL.Query()
	reqType := q.Get("type")

	// 判断是否为定义的登录类型
	isHas := slices.Contains(loginType, reqType)
	if !isHas {
		response.Error(c, ecode.InvalidParams.RewriteMsg("未知的登录类型"))
		return
	}

	var form types.UserReq
	if reqType == "account" {
		form = &types.UserAccountLogoutReq{}
	}
	if reqType == "email" {
		form = &types.UserEmailLoginReq{}
	}

	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	var findUser *model.User

	// 根据类型不同，调用不同的查询方法
	switch v := form.(type) {
	case *types.UserAccountLogoutReq:
		findUser, err = u.iDao.FindByAccount(ctx, v.Account)
	case *types.UserEmailLoginReq:
		findUser, err = u.iDao.FindByEmail(ctx, v.Email)
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, ecode.ErrLoginUser.RewriteMsg("用户不存在"))
			return
		}
		response.Error(c, ecode.InternalServerError)
		return
	}

	// 验证密码
	correct := gocrypto.VerifyPassword(form.GetPassword(), findUser.Password)
	if !correct {
		response.Error(c, ecode.ErrLoginUser.RewriteMsg("密码错误"))
		return
	}
	uid := strconv.FormatUint(findUser.ID, 10)
	token, _ := jwt.GenerateToken(uid, findUser.Username)

	// redis 缓存token
	err = uCache.Set(ctx, findUser.ID, &token, cache.UserCacheExpireTime)

	response.Success(c, map[string]string{
		"token": token,
	})
}

// Logout 退出
func (u *userHandler) Logout(c *gin.Context) {
	uid, _ := c.Get("uid")
	fmt.Println(uid)
	response.Success(c, nil)
}

// UpdateByID 更新
func (u *userHandler) UpdateByID(c *gin.Context) {
	_, id, isAbort := common.GetIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	form := &types.UpdateUserReq{}

	// 获取需要更新的数据
	update, err := common.GetTransmitFields(c, form)

	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	user := &model.User{}
	user.ID = id

	ctx := middleware.WrapCtx(c)
	err = u.iDao.UpdateByID(ctx, user, update)

	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}
	response.Success(c)
}
func (u *userHandler) List(c *gin.Context)    {}
func (u *userHandler) GetByID(c *gin.Context) {}
