package handler

import (
	"com/chat/service/internal/ecode"
	"com/chat/service/internal/handler/common"
	"com/chat/service/pkg/gin/middleware"
	"com/chat/service/pkg/gin/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

var _ RoleHandler = (*roleHandler)(nil)

type roleHandler struct {

}

type RoleHandler interface {
	Create(c *gin.Context)
}

func NewRoleHandler() RoleHandler {
	return &roleHandler{}
}

func (h *roleHandler) Create(c *gin.Context) {
	ctx := middleware.WrapCtx(c)
	enforcer, e := common.CasbinEnforcer(ctx)
	if e != nil {
		response.Error(c, ecode.InvalidParams, http.StatusForbidden)
		return
	}
	_, e = enforcer.AddPolicy("root", "menu_id", "All")
	response.Success(c, gin.H{
		"message": "success",
	})
}