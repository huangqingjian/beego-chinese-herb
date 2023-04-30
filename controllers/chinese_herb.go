package controllers

import (
	"beego-chinese-herb/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/i18n"
	"strconv"
)

// 中药API
type ChineseHerbController struct {
	BaseController
}

// @Title 查询中药列表
// @Description 通过条件查询中药列表
// @param q query string false "中药名"
// @param type query int64 false "中药类型"
// @param pageSize query int false "分页大小"
// @param pageNo query int false "页码"
// @Success 200 {object} models.Response
// @Failure 200 a error response object
// @router /list [get]
func (c *ChineseHerbController) GetChineseHerbList() {
	// 查询中药
	query := getChineseHerbQuery(c)
	page, err := models.GetChineseHerbList(query)
	if err != nil {
		logs.Error("GetChineseHerbList Error[%v]", err)
		panic(any(models.NewServiceError(i18n.Tr(c.Lang, "error.findchineseherb"))))
	}
	c.Data["json"] = models.ResponseSuccess(&page)
	c.ServeJSON()
}

// @Title 查找中药
// @Description 通过id查找中药详情
// @param id path int true "中药id"
// @Success 200 {object} models.Response
// @Failure 200 a error response object
// @router /get/:id [get]
func (c *ChineseHerbController) GetChineseHerb() {
	// 参数校验
	id := c.Ctx.Input.Param(":id")
	chineseHerbId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(any(models.NewParamError(i18n.Tr(c.Lang, "error.chineseherbid"))))
	}
	// 查找中药
	chineseHerb, err := models.GetChineseHerb(chineseHerbId)
	if err != nil {
		if err == orm.ErrNoRows {
			panic(any(models.NewServiceError(i18n.Tr(c.Lang, "error.chineseherbnotfound"))))
		}
		logs.Error("GetUser Error[%v]", err)
		panic(any(models.NewServiceError(i18n.Tr(c.Lang, "error.findchineseherb"))))
	}
	// 查找药方
	herbPharmacys := getHerbPharmacyList(c, chineseHerbId)
	chineseHerb.HerbPharmacys = herbPharmacys
	c.Data["json"] = models.ResponseSuccess(&chineseHerb)
	c.ServeJSON()
}

// 查询药方
func getHerbPharmacyList(c *ChineseHerbController, chineseHerbId int64) []models.HerbPharmacy {
	hps, err := models.GetHerbPharmacyList(&models.HerbPharmacyQuery{HerbId: chineseHerbId})
	if err != nil {
		panic(any(models.NewServiceError(i18n.Tr(c.Lang, "error.findherbpharmacy"))))
	}
	return hps
}

// 获取中药查询条件
func getChineseHerbQuery(c *ChineseHerbController) *models.ChineseHerbQuery {
	query := models.ChineseHerbQuery{}
	query.Q = c.GetString("q")
	query.Type, _ = c.GetInt64("type")
	query.PageNum, _ = c.GetInt32("page", 1)
	query.PageSize, _ = c.GetInt32("limit", 10)
	return &query
}