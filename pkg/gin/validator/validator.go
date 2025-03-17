// Package validator is gin request parameter check library.
package validator

import (
	"reflect"
	"sync"

	valid "github.com/go-playground/validator/v10"
)

// Init request body file valid
func Init() *CustomValidator {
	validator := NewCustomValidator()
	validator.Engine()
	return validator
}

// CustomValidator Custom valid objects
type CustomValidator struct {
	Once     sync.Once
	Validate *valid.Validate
}

// NewCustomValidator Instantiate
func NewCustomValidator() *CustomValidator {
	return &CustomValidator{}
}

// ValidateStruct Instantiate struct valid 实例化valid结构
func (v *CustomValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		v.lazyInit()
		if err := v.Validate.Struct(obj); err != nil {
			return err
		}
	}

	return nil
}

// Engine Instantiate valid
func (v *CustomValidator) Engine() interface{} {
	v.lazyInit()
	return v.Validate
}

func (v *CustomValidator) lazyInit() {
	v.Once.Do(func() {
		v.Validate = valid.New()
		v.Validate.SetTagName("binding")

		// 注册自定义校验规则
		// optional_not_empty 空效验 字段类型要设置为指针类型
		_ = v.Validate.RegisterValidation("optional_not_empty", validateOptionalNotEmpty, true)
	})
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

func validateOptionalNotEmpty(fl valid.FieldLevel) bool {
	// 获取字段的值
	field := fl.Field()
	if field.Kind() == reflect.Ptr && field.IsNil() {
		return true
	}

	// 如果字段传入了，则检查其值是否为空
	switch field.Kind() {
	case reflect.String:
		return field.String() != ""
	case reflect.Slice, reflect.Map, reflect.Array:
		return field.Len() > 0
	case reflect.Ptr, reflect.Interface:
		return !field.IsNil()
	default:
		return true // 其他类型默认通过
	}
}
