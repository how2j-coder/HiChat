package common

import (
	"HiChat/utils"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
)

var priPemFile = utils.GetRootPath() + "/common/pem/private-key.pem"
var pubPemFile = utils.GetRootPath() + "/common/pem/public-key.pem"

type Claims struct {
	info interface{}
	jwt.RegisteredClaims
}

// GenerateTaken 生成token 非对称加密RS256
func GenerateTaken(payload interface{}) string {
	infoData := payload
	claims := Claims{
		info:             infoData,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	// 读取私钥
	privateKeyDataPem, err := os.ReadFile(priPemFile)
	if err != nil {
		log.Fatal(err)
	}

	// ParseRSAPrivateKeyFromPEM
	privateKeyData, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyDataPem)
	if err != nil {
		log.Fatal(err)
	}
	// 2 把token加密
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(privateKeyData)
	return token
}

func DecryptTaken(tokenStr string) (payload interface{}, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 读取公钥
		publicKeyDataPem, err := os.ReadFile(pubPemFile)
		if err != nil {
			log.Fatal(err)
		}
		// ParseRSAPublicKeyFromPEM
		return jwt.ParseRSAPublicKeyFromPEM(publicKeyDataPem)
	})
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		fmt.Println(claims)
		return claims.info, nil
	} else {
		return nil, err
	}
}
