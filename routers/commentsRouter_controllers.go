package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["beego-chinese-herb/controllers:BannerController"] = append(beego.GlobalControllerRouter["beego-chinese-herb/controllers:BannerController"],
        beego.ControllerComments{
            Method: "GetBannerList",
            Router: "/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-chinese-herb/controllers:ChineseHerbController"] = append(beego.GlobalControllerRouter["beego-chinese-herb/controllers:ChineseHerbController"],
        beego.ControllerComments{
            Method: "GetChineseHerb",
            Router: "/get/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-chinese-herb/controllers:ChineseHerbController"] = append(beego.GlobalControllerRouter["beego-chinese-herb/controllers:ChineseHerbController"],
        beego.ControllerComments{
            Method: "GetChineseHerbList",
            Router: "/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-chinese-herb/controllers:UploadController"] = append(beego.GlobalControllerRouter["beego-chinese-herb/controllers:UploadController"],
        beego.ControllerComments{
            Method: "UploadImg",
            Router: "/img",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-chinese-herb/controllers:UserController"] = append(beego.GlobalControllerRouter["beego-chinese-herb/controllers:UserController"],
        beego.ControllerComments{
            Method: "AddUser",
            Router: "/add",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-chinese-herb/controllers:UserController"] = append(beego.GlobalControllerRouter["beego-chinese-herb/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUser",
            Router: "/get/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-chinese-herb/controllers:UserController"] = append(beego.GlobalControllerRouter["beego-chinese-herb/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUserList",
            Router: "/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
