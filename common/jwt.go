package common

import (
	"HiChat/global"
	"HiChat/utils"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"time"
)

var priPemFile = utils.GetRootPath() + "/common/pem/private-key.pem"
var pubPemFile = utils.GetRootPath() + "/common/pem/public-key.pem"

type CuClaims struct {
	Data interface{} `json:"info"`
	jwt.RegisteredClaims
}

// EncryptTaken 生成token 非对称加密RS256
func EncryptTaken(payload interface{}, iss string, expiresTime time.Duration) string {
	uuidValue := uuid.New()
	claims := CuClaims{
		Data: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    iss,                                                            // 签发者
			Subject:   "Auth_Server",                                                  // 签发对象
			Audience:  jwt.ClaimStrings{"Android_APP", "Web_APP"},                     //签发受众
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                 //签发时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)),                //最早使用时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresTime)), //过期时间
			ID:        uuidValue.String(),                                             // wt ID, 类似于盐值
		},
	}
	// 读取私钥
	privateKeyDataPem, err := os.ReadFile(priPemFile)
	if err != nil {
		global.Logger.Error(err.Error())
	}

	// ParseRSAPrivateKeyFromPEM
	privateKeyData, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyDataPem)
	if err != nil {
		global.Logger.Error(err.Error())
	}
	// 2 把token加密
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(privateKeyData)
	return token
}

// DecryptTaken 解密效验token
func DecryptTaken(tokenStr string) (payload *CuClaims, errs error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CuClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 读取公钥
		publicKeyDataPem, err := os.ReadFile(pubPemFile)
		if err != nil {
			global.Logger.Error(err.Error())
		}
		// ParseRSAPublicKeyFromPEM
		return jwt.ParseRSAPublicKeyFromPEM(publicKeyDataPem)
	})
	if claims, ok := token.Claims.(*CuClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// GenerateTaken 生成token(access & refresh)
func GenerateTaken(payload interface{}, iss string, expiresTime time.Duration) (accessToken, refreshToken string) {
	accessToken = EncryptTaken(payload, iss, expiresTime)
	// 刷新的token 有效期7天
	refreshToken = EncryptTaken(payload, iss,time.Hour * 24 * 7)
	return
}

func RefreshToken(aToken, rToken string) (newAToken, newRToken string)  {
	// refresh token无效直接返回
	_, err := DecryptTaken(rToken)
	if err != nil {
		return "", ""
	}

	v, err := DecryptTaken(aToken)
	if errors.Is(err, jwt.ErrTokenExpired) {
		newAToken = EncryptTaken(v.Data, "Auth_Server", time.Hour * 2)
		return newAToken, ""
	}
	return
}