package controllers

import (
	"beego-chinese-herb/models"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/i18n"
	"strconv"
)

// 用户API
type UserController struct {
	BaseController
}

// @Title 新增用户
// @Description 新增用户
// @Param body body models.User	true "body for user content"
// @Success 200 {object} models.Response
// @Failure 200 error response object
// @router /add [post]
func (u *UserController) AddUser() {
	user := models.User{}
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		panic(any(models.NewParamError(i18n.Tr(u.Lang, "error.paramformat"))))
	}
	// 参数验证
	valid := validation.Validation{}
	b, err := valid.Valid(&user)
	if err != nil {
		panic(any(models.NewParamError(i18n.Tr(u.Lang, "error.paramvalid"))))
	}
	// 验证不通过
	if !b {
		// fast fail
		for _, err := range valid.Errors {
			panic(any(models.NewParamError(err.Message)))
		}
	}
	result, err := models.AddUser(&user)
	if err != nil {
		panic(any(models.NewParamError("error.userinsert")))
	}
	u.Data["json"] = models.ResponseSuccess(result)
	u.ServeJSON()
}

// @Title 查找用户
// @Description 通过id查找用户详情
// @param id path int true "user id"
// @Success 200 {object} models.Response
// @Failure 200 a error response object
// @router /get/:id [get]
func (u *UserController) GetUser() {
	id := u.Ctx.Input.Param(":id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(any(models.NewParamError(i18n.Tr(u.Lang, "error.userid"))))
	}
	user, err := models.GetUser(userId)
	if err != nil {
		if err == orm.ErrNoRows {
			panic(any(models.NewServiceError(i18n.Tr(u.Lang, "error.usernotfound"))))
		}
		logs.Error("GetUser Error[%v]", err)
		panic(any(models.NewServiceError(i18n.Tr(u.Lang, "error.finduser"))))
	}
	u.Data["json"] = models.ResponseSuccess(&user)
	u.ServeJSON()
}

// @Title 查询用户列表
// @Description 通过条件查询用户列表
// @param q query string true "用户名或手机号"
// @param pageSize query int false "分页大小"
// @param pageNo query int false "页码"
// @Success 200 {object} models.Response
// @Failure 200 a error response object
// @router /list [get]
func (u *UserController) GetUserList() {
	query := getUserQuery(u)
	users, err := models.GetUserList(query)
	if err != nil {
		logs.Error("GetUserList Error[%v]", err)
		panic(any(models.NewServiceError(i18n.Tr(u.Lang, "error.finduser"))))
	}
	u.Data["json"] = models.ResponseSuccess(&users)
	u.ServeJSON()
}

// 获取用户查询条件
func getUserQuery(u *UserController) *models.UserQuery {
	query := models.UserQuery{}
	query.Q = u.GetString("q")
	query.PageNum, _ = u.GetInt32("page", 1)
	query.PageSize, _ = u.GetInt32("limit", 10)
	return &query
}
