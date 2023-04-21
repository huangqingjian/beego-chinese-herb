package component

import (
	"beego-chinese-herb/constant"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

var(
	log = "log"
	fileName = "filename"
	maxSize = "maxsize"
)

// 初始化日志
func init() {
	// 滚动日志
	logConf := make(map[string]interface{})
	logConf[fileName], _ = web.AppConfig.String(log + constant.SMH + fileName)
	logConf[maxSize], _ = web.AppConfig.Int(log + constant.SMH + maxSize)

	confJson, _ := json.Marshal(logConf)
	logs.SetLogger(logs.AdapterFile, string(confJson))
	logs.SetLevel(logs.LevelInformational)
	logs.SetLogFuncCall(true)
	logs.Async()
}
