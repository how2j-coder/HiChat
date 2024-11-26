package utils

import (
	"errors"
	"github.com/go-playground/validator/v10"
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
		errText := filed.Tag.Get(errMsg.Tag() + "Msg")
		// 如果没有自定义消息，返回错误成员本身的错误信息
		if errText == "" {
			return err.Error()
		}
		return errText
	}

	// 其他类型的错误直接返回错误信息
	return err.Error()

}
