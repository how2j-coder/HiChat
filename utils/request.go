package utils

import (
	"errors"
	"fmt"
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

// GetJsonAndExistField 获取传递的字段
func GetJsonAndExistField(ctx *gin.Context, target interface{}) (map[string]interface{}, error) {
	if err := ctx.ShouldBindBodyWithJSON(target); err != nil {
		return nil, err
	}

	var jsonMap map[string]interface{}
	if err := ctx.ShouldBindBodyWithJSON(&jsonMap); err != nil {
		return nil, err
	}

	tags := reflect.TypeOf(target).Elem()

	for i := 0; i < tags.NumField(); i++ {
		field := tags.Field(i)
		//判断字段是否有 has_required tag 验证字段是否传递
		//如果前端未传递某个字段，validator会将该字段设置为其类型的零值。
		//如果前端传递了该字段，并且是零值，validator同样会将其设置为相应的零值。
		//区分字段是未传递还是传递了零值
		errMsg := field.Tag.Get("has_required")
		if errMsg != "" {
			jsFiled := field.Tag.Get("json")
			value, ok := jsonMap[jsFiled]
			if ok {
				switch field.Type.Kind().String() {
				case "string":
					if value == "" {
						fmt.Println(errMsg)
						return nil, errors.New(errMsg)
					}
				}
			}
		}
	}



	for key, value := range jsonMap {
		jsonMap[CamelToSnake(key)] = value
	}
	return jsonMap, nil
}

//// GisterValidation 注册自定义参数效验
//var hasReqData validator.Func = func(fl validator.FieldLevel) bool {
//	fmt.Println(12312312)
//	return true
//}
//
//func GisterValidation() {
//	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
//			_ = v.RegisterValidation("has_required", hasReqData)
//	}
//}