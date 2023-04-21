package component

import (
	"beego-chinese-herb/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

// 初始化
func init() {
	// 异常处理
	web.BConfig.RecoverPanic = true
	web.BConfig.RecoverFunc = recoverPanic
	// 图片路径配置
	picturePath, _ := web.AppConfig.String("imagepath")
	web.SetStaticPath("/picture", picturePath)
	// swagger配置
	if web.BConfig.RunMode == "dev" {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.SetStaticPath("/swagger", "swagger")
	}
	// filter配置
	// auth校验
	web.InsertFilter("/*", web.BeforeRouter, auth())
	// 支持跨域
	web.InsertFilter("/*", web.BeforeRouter, allowCors())
	// filterChain配置
	web.InsertFilterChain("/*", func(next web.FilterFunc) web.FilterFunc {
		return func(ctx *context.Context) {
			next(ctx)
		}
	})

}

// auth校验（登录态等）
func auth() web.FilterFunc {
	return func(ctx *context.Context) {
			//userId := ctx.Input.Session("user_id")
			//if(userId == nil) {
			//	panic(any(models.NewAuthError("请登录后再访问")))
			//}
		}
}

// 跨域
func allowCors() web.FilterFunc {
	return cors.Allow(&cors.Options{
		// 允许访问所有源
		AllowAllOrigins: true,
		// 可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// 指的是允许的Header的种类
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	})
}

// 恢复异常
func recoverPanic(ctx *context.Context, config *web.Config) {
	if err := interface{}(recover()); err != nil {
		logs.Error("error[%v]", err)
		if err == any(web.ErrAbort) {
			return
		}
		hasIndent := web.BConfig.RunMode != web.PROD
		switch obj := interface{}(err).(type) {
		case models.AuthError:
			ctx.Output.JSON(models.ResponseFail(obj.Code, obj.Message), hasIndent, false)
		case models.ParamError:
			ctx.Output.JSON(models.ResponseFail(obj.Code, obj.Message), hasIndent, false)
		case models.ServiceError:
			ctx.Output.JSON(models.ResponseFail(obj.Code, obj.Message), hasIndent, false)
		default:
			ctx.Output.JSON(models.ResponseFastFail("服务器内部异常～"), hasIndent, false)
		}
	}
}