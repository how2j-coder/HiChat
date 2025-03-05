package middleware

import (
	"com/chat/service/pkg/errcode"
	"com/chat/service/pkg/gin/response"
	"com/chat/service/pkg/jwt"
	"com/chat/service/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
	"strings"
)

const (
	// HeaderAuthorizationKey http header authorization key
	HeaderAuthorizationKey = "Authorization"

	defaultAuthType = 1 // use default auth
	customAuthType  = 2 // use custom auth
)

// ExtraDefaultVerifyFn extra default verify function, tokenTail10 is the last 10 characters of the token.
type ExtraDefaultVerifyFn = func(claims *jwt.Claims, tokenTail10 string, c *gin.Context) error

// ExtraCustomVerifyFn extra custom verify function, tokenTail10 is the last 10 characters of the token.
type ExtraCustomVerifyFn = func(claims *jwt.CustomClaims, tokenTail10 string, c *gin.Context) error


type authOptions struct {
	isSwitchHTTPCode bool

	authType        int
	defaultVerifyFn ExtraDefaultVerifyFn
	customVerifyFn  ExtraCustomVerifyFn
}

// AuthOption set the auth options.
type AuthOption func(*authOptions)


var whiteList map[string]string
func ignoreWhiteList(url *url.URL, method string) bool {
	queryUrl := strings.Split(fmt.Sprint(url), "?")[0]
	if _, ok := whiteList[queryUrl]; ok {
		if whiteList[queryUrl] == method {
			return true
		}
		return false
	}
	return false
}


func defaultAuthOptions() *authOptions {
	return &authOptions{
		isSwitchHTTPCode: false,
		authType:         defaultAuthType,
	}
}

func (o *authOptions) apply(opts ...AuthOption) {
	for _, opt := range opts {
		opt(o)
	}
}

// WithSwitchHTTPCode switch to http code
func WithSwitchHTTPCode() AuthOption {
	return func(o *authOptions) {
		o.isSwitchHTTPCode = true
	}
}

// WithDefaultVerify set default verify type
func WithDefaultVerify(fn ...ExtraDefaultVerifyFn) AuthOption {
	return func(o *authOptions) {
		o.authType = defaultAuthType
		if len(fn) > 0 {
			o.defaultVerifyFn = fn[0]
		}
	}
}

// WithCustomVerify set custom verify type with extra verify function
func WithCustomVerify(fn ...ExtraCustomVerifyFn) AuthOption {
	return func(o *authOptions) {
		o.authType = customAuthType
		if len(fn) > 0 {
			o.customVerifyFn = fn[0]
		}
	}
}

// AddWhiteRouter add white route
func AddWhiteRouter(route map[string]string) AuthOption {
	return func(_ *authOptions) {
		// 如果 whiteList 还未初始化，则初始化
		if whiteList == nil {
			whiteList = make(map[string]string)
		}

		// 将传入的 route 中的键值对添加到 whiteList 中
		for key, value := range route {
			whiteList[key] = value
		}
	}
}

func responseUnauthorized(c *gin.Context, isSwitchHTTPCode bool) {
	if isSwitchHTTPCode {
		response.Out(c, errcode.Unauthorized)
	} else {
		response.Error(c, errcode.Unauthorized)
	}
}

// -------------------------------------------------------------------------------------------

// Auth authorization
func Auth(opts ...AuthOption) gin.HandlerFunc {
	o := defaultAuthOptions()
	o.apply(opts...)

	return func(c *gin.Context) {
		// 忽略白名单路由
		if ignoreWhiteList(c.Request.URL, c.Request.Method) {
			c.Next()
			return
		}

		authorization := c.GetHeader(HeaderAuthorizationKey)
		if len(authorization) < 100 {
			logger.Warn("authorization is illegal")
			responseUnauthorized(c, o.isSwitchHTTPCode)
			c.Abort()
			return
		}

		token := authorization[7:] // remove Bearer prefix

		if o.authType == customAuthType {
			// custom auth
			claims, err := jwt.ParseCustomToken(token)
			if err != nil {
				logger.Warn("ParseToken error", logger.Err(err))
				responseUnauthorized(c, o.isSwitchHTTPCode)
				c.Abort()
				return
			}
			// extra verify function
			if o.customVerifyFn != nil {
				tokenTail10 := token[len(token)-10:]
				if err = o.customVerifyFn(claims, tokenTail10, c); err != nil {
					// logger.Warn("verify error", logger.Err(err), logger.Any("fields", claims.Fields))
					responseUnauthorized(c, o.isSwitchHTTPCode)
					c.Abort()
					return
				}
			}
		} else {
			// default auth
			claims, err := jwt.ParseToken(token)
			if err != nil {
				logger.Warn("ParseToken error", logger.Err(err))
				responseUnauthorized(c, o.isSwitchHTTPCode)
				c.Abort()
				return
			}
			// extra verify function
			if o.defaultVerifyFn != nil {
				tokenTail10 := token[len(token)-10:]
				if err = o.defaultVerifyFn(claims, tokenTail10, c); err != nil {
					// logger.Warn("verify error", logger.Err(err), logger.String("uid", claims.UID), logger.String("name", claims.Name))
					responseUnauthorized(c, o.isSwitchHTTPCode)
					c.Abort()
					return
				}
			} else {
					c.Set("uid", claims.UID)
					c.Set("name", claims.Name)
			}
		}

		c.Next()
	}
}

// JwtOption set the auth options.
type JwtOption = AuthOption
type VerifyCustomFn = ExtraCustomVerifyFn

// AuthCustom custom authentication
func AuthCustom(fn VerifyCustomFn, opts ...JwtOption) gin.HandlerFunc {
	o := defaultAuthOptions()
	o.apply(opts...)

	return func(c *gin.Context) {
		authorization := c.GetHeader(HeaderAuthorizationKey)
		if len(authorization) < 150 {
			logger.Warn("authorization is illegal")
			responseUnauthorized(c, o.isSwitchHTTPCode)
			c.Abort()
			return
		}

		token := authorization[7:] // remove Bearer prefix
		claims, err := jwt.ParseCustomToken(token)
		if err != nil {
			logger.Warn("ParseToken error", logger.Err(err))
			responseUnauthorized(c, o.isSwitchHTTPCode)
			c.Abort()
			return
		}

		if fn != nil {
			tokenTail10 := token[len(token)-10:]
			if err = fn(claims, tokenTail10, c); err != nil {
				logger.Warn("verify error", logger.Err(err), logger.Any("fields", claims.Fields))
				responseUnauthorized(c, o.isSwitchHTTPCode)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}