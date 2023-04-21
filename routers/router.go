// @APIVersion 1.0.0
// @Title beego 中药 API
// @Description 中药相关接口，包括用户、banner、中药、药方等
// @Contact 2366850717@qq.com
package routers

import (
	"beego-chinese-herb/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/banner",
			beego.NSInclude(
				&controllers.BannerController{},
			),
		),
		beego.NSNamespace("/chineseHerb",
			beego.NSInclude(
				&controllers.ChineseHerbController{},
			),
		),
		beego.NSNamespace("/upload",
			beego.NSInclude(
				&controllers.UploadController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
