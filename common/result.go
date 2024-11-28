package common

import (
	"encoding/json"
)

type Response struct {
	Code    int         `json:"code"`    // 业务状态码 200
	Status  bool        `json:"status"`  // 状态信息 True | False
	Message string      `json:"message"` // 返回信息 Operation successful
	Data    interface{} `json:"data"`    // 返回数据 Data
}

// WithMsg 自定义响应信息
func (res *Response) WithMsg(message string) Response {
	return Response{
		Code:    res.Code,
		Status:  res.Status,
		Message: message,
		Data:    res.Data,
	}
}

// WithData 追加响应数据
func (res *Response) WithData(data interface{}) Response {
	return Response{
		Code:    res.Code,
		Status:  res.Status,
		Message: res.Message,
		Data:    data,
	}
}

// ToString 返回 JSON 格式的错误详情
func (res *Response) ToString() string {
	err := &struct {
		Code    int         `json:"code"`
		Status  bool        `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Code:    res.Code,
		Message: res.Message,
		Data:    res.Data,
	}
	raw, _ := json.Marshal(err)
	return string(raw)
}

// 构造函数
func response(code int, status bool, msg string) *Response {
	return &Response{
		Code:    code,
		Status:  status,
		Message: msg,
		Data:    nil,
	}
}

var (
	SuccessCode = 1000
	ErrorCode   = 1001
	ParamsErrorCode   = 1002 // 参数错误
	NotFoundCode = 1004
)

var (
	Success        = response(SuccessCode, true, "success")
	Error          = response(ErrorCode, true, "error")
	NotFound       = response(NotFoundCode, true, "not found")
	ParamsNilError = response(ParamsErrorCode, false, "the parameter is the default")
)
