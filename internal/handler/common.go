package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"reflect"
	"slices"
)

// 提前符合structFiled中的map数据 ShouldBindBodyWithJSON
func getTransmitFields(c *gin.Context, structFild interface{}) (map[string]interface{}, error) {
	reqFields := make(map[string]interface{})
	err := c.ShouldBindBodyWithJSON(&reqFields)


	if err != nil && err == io.EOF {
		err = errors.New("no data please use ShouldBindBodyWithJSON bind Data")
		return nil, err
	}

	var reqFieldKeys []string
	for k, _ := range reqFields {
		reqFieldKeys = append(reqFieldKeys, k)
	}

	val := reflect.ValueOf(structFild)
	typ := reflect.TypeOf(structFild)

	// 如果传入的是指针，则获取其指向的值
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}


	// 遍历结构体字段
	fields := make(map[string]interface{})
	for i := 0; i < val.NumField(); i++ {
		fieldName := typ.Field(i).Tag.Get("json")
		if slices.Contains(reqFieldKeys, fieldName) {
			fields[fieldName] = reqFields[fieldName]
		}
	}
	return fields, nil
}