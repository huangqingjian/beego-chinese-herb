package test

import (
	"beego-chinese-herb/component"
	"beego-chinese-herb/models"
	_ "beego-chinese-herb/routers"
	j "encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	// 初始化web配置
	component.InitWebConfig()
}

// TestGetChineseHerbList 测试中药列表查询
func TestGetChineseHerbList(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/chineseHerb/list?page=1&limit=2", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("测试中药列表查询API\n", t, func() {
		Convey("http状态码：200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		if w.Code != 200 {
			return
		}
		r := models.Response{}
		j.Unmarshal(w.Body.Bytes(), &r)
		Convey("返回对象code编码：0", func() {
			So(r.Code, ShouldEqual, 0)
		})
	})
}

// TestGetChineseHerbList 通过id查找中药详情
func TestGetChineseHerb(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/chineseHerb/get/1000", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("测试通过id查找中药详情API\n", t, func() {
		Convey("http状态码：200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		if w.Code != 200 {
			return
		}
		r := models.Response{}
		j.Unmarshal(w.Body.Bytes(), &r)
		Convey("返回对象code编码：0", func() {
			So(r.Code, ShouldEqual, 0)
		})
	})
}

