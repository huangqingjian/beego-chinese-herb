package controllers

import (
	"beego-chinese-herb/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/i18n"
)

// 横幅API
type BannerController struct {
	BaseController
}

// @Title 查询banner列表
// @Description 查询全部banner
// @Success 200 {object} models.Response
// @Failure 200 a error response object
// @router /list [get]
func (b *BannerController) GetBannerList() {
	banners, err := models.GetBannerList()
	if err != nil {
		logs.Error("GetBannerList Error[%v]", err)
		panic(any(models.NewServiceError(i18n.Tr(b.Lang, "error.findbanner"))))
	}
	b.Data["json"] = models.ResponseSuccess(banners)
	b.ServeJSON()
}
