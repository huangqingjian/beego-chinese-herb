package controllers

import (
	"beego-chinese-herb/constant"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
	"strings"
)

type BaseController struct {
	web.Controller
	i18n.Locale
}

// 初始化
func init() {
	// 自定义错误处理器
	web.ErrorController(&ErrorController{})
	// 本地化设置
	setLocales()
}

// 定义prepare方法
func (b *BaseController) Prepare() {
	setLang(b)
}

// 变量
var (
	lang = "lang"
	types = "types"
	acceptLanguage = "Accept-Language"
)

// 设置本地化
func setLocales() {
	langTypes, _ := web.AppConfig.String(lang + constant.SMH + types)
	langTypeArr := strings.Split(langTypes, constant.SX)
	for _, lang := range langTypeArr {
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			logs.Error("本地化文件设置失败[%v]", err)
			return
		}
	}
}
// 设置本地化语言
func setLang(b *BaseController) {
	// 从url参数获取lang，url中没有，从请求头获取
	lang := b.GetString(lang)
	if len(lang) == 0 {
		al := b.Ctx.Request.Header.Get(acceptLanguage)
		if len(al) > 4 {
			lang = al[:5]
		}
	}
	// 检查lang是否在i18n中存在
	if !i18n.IsExist(lang) {
		lang = ""
	}
	// 默认为中文
	if len(lang) == 0 {
		lang = "zh-CN";
	}
	b.Lang = lang
}