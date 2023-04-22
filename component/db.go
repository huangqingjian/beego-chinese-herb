package component

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

// 变量
var (
	sqlAlias = "default"
	sqlDriver = "sqldriver"
	sqlConn = "sqlconn"
)

// 初始化数据库
func InitDB() {
	//数据库连接
	sqlDriver, _ := web.AppConfig.String(sqlDriver)
	sqlConn, _ := web.AppConfig.String(sqlConn)

	orm.RegisterDriver(sqlDriver, orm.DRMySQL)
	orm.RegisterDataBase(sqlAlias, sqlDriver, sqlConn)

	// dev环境开启日志打印
	if web.BConfig.RunMode == "dev" {
		orm.Debug = true
	}
}
