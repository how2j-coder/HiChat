package common

import (
	"HiChat/global"
	"HiChat/utils"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var priPemFile = utils.GetRootPath() + "/common/pem/private-key.pem"
var pubPemFile = utils.GetRootPath() + "/common/pem/public-key.pem"

type CuClaims struct {
	InfoT interface{} `json:"info_t"`
	jwt.RegisteredClaims
}

// GenerateTaken 生成token 非对称加密RS256
func GenerateTaken(payload interface{}, salt string) string {
	claims := CuClaims{
		InfoT: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Auth_Server",                                   // 签发者
			Subject:   "Auth_Server",                                   // 签发对象
			Audience:  jwt.ClaimStrings{"Android_APP", "Web_APP"},      //签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),   //过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)), //最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                  //签发时间
			ID:        salt,                                            // wt ID, 类似于盐值
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

func DecryptTaken(tokenStr string) (payload interface{}, err error) {
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
		return claims.InfoT, nil
	} else {
		return nil, err
	}
}
