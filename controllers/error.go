package controllers

import (
	"beego-chinese-herb/models"
	"github.com/beego/beego/v2/server/web"
)

// 自定义错误 API
type ErrorController struct {
	web.Controller
}

// 400
func (e *ErrorController) Error400() {
	e.Data["json"] = models.ResponseFail(400, "参数错误～")
	e.ServeJSON()
}

// 401
func (e *ErrorController) Error401() {
	e.Data["json"] = models.ResponseFail(401, "访问非法～")
	e.ServeJSON()
}

// 404
func (e *ErrorController) Error404() {
	e.Data["json"] = models.ResponseFail(404, "路径非法～")
	e.ServeJSON()
}

// 500
func (e *ErrorController) Error500() {
	e.Data["json"] = models.ResponseFastFail("服务器内部错误")
	e.ServeJSON()
}