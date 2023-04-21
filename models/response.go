package models

// 响应数据
type Response struct {
	Code    int			`json:"code"`
	Message string		`json:"message"`
	Data    interface{} `json:"data"`
}

// 成功
func ResponseSuccess(data interface{}) *Response {
	return &Response{Code: 0, Data: data}
}

// 失败
func ResponseFail(code int, message string) *Response {
	return &Response{Code: code, Message: message}
}

// 失败
func ResponseFastFail(message string) *Response {
	return ResponseFail(1, message)
}

