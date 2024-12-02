package middlewear

import (
	. "HiChat/common"
	"HiChat/global"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 authorization header：获取前端传过来的信息的
		tokenStr := ctx.GetHeader("Authorization")

		// 验证是否传入token
		if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, Error.WithMsg("认证失败"))
			ctx.Abort()
			return
		}

		tokenStr = tokenStr[len("Bearer "):]

		// 解析Token并验证
		jwtClaims, err := DecryptTaken(tokenStr)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				ctx.JSON(http.StatusUnauthorized, Error.WithMsg("认证已过期"))
			} else {
				ctx.JSON(http.StatusUnauthorized, Error.WithMsg("认证失败"))
			}
			global.Logger.Error(err.Error())
			ctx.Abort()
			return
		}
		ctx.Set("userId", jwtClaims.Data)
		ctx.Next()
	}
}
