package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"math/rand"
	"reflect"
)

// ValidateErrorMsg 消息类型
type ValidateErrorMsg map[string]string

func GetErrorMsg(err error, structure interface{}) string {
	var validationErrors validator.ValidationErrors
	// 如果err是validator.ValidationErrors类型，获取第一个错误信息
	if errors.As(err, &validationErrors) && len(validationErrors) > 0 {
		s := reflect.TypeOf(structure)
		errMsg := validationErrors[0]
		filed, _ := s.FieldByName(errMsg.Field())
		tag := ToCamelCaseLower(errMsg.Tag())
		errText := filed.Tag.Get(tag + "Msg")
		// 如果没有自定义消息，返回错误成员本身的错误信息
		if errText == "" {
			return err.Error()
		}
		return errText
	}

	// 其他类型的错误直接返回错误信息
	return err.Error()
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// GetJsonAndExistField 获取传递的字段 TODO: chose
func GetJsonAndExistField(ctx *gin.Context, target interface{}) (map[string]interface{}, error) {
	if err := ctx.ShouldBind(&target); err != nil {
		return nil, err
	}

	var jsonMap map[string]interface{}
	if err := ctx.ShouldBindBodyWithJSON(&jsonMap); err != nil {
		return nil, err
	}

	return jsonMap, nil
}