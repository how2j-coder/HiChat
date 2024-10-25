package common

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`    // 状态码 200
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
		Message string      `json:"Message"`
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
		Status:  status, // 当我传入status时, 值为我传入的值, 否则Status默认为true
		Message: msg,
		Data:    nil,
	}
}

var (
	Success = response(http.StatusOK, true, "success")
	Error   = response(http.StatusInternalServerError, true, "error")
)
