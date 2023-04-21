package models

import "encoding/json"

// 参数错误
type ParamError struct {
	Code int
	Message string
}

func (p ParamError) Error() string {
	err, _ := json.Marshal(p)
	return string(err)
}

func NewParamError(message string) error {
	return ParamError{
		Code: -1,
		Message: message,
	}
}

// 鉴权错误
type AuthError struct {
	Code int
	Message string
}

func (a AuthError) Error() string {
	err, _ := json.Marshal(a)
	return string(err)
}

func NewAuthError(message string) error {
	return AuthError{
		Code: 401,
		Message: message,
	}
}

// 服务错误
type ServiceError struct {
	Code int
	Message string
}

func (s ServiceError) Error() string {
	err, _ := json.Marshal(s)
	return string(err)
}

func NewServiceError(message string) error {
	return ServiceError{
		Code: -1,
		Message: message,
	}
}