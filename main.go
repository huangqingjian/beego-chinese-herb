package main

import (
	_ "beego-chinese-herb/component"
	_ "beego-chinese-herb/routers"
	"github.com/beego/beego/v2/adapter/toolbox"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// 开启定时任务
	initTask()
	toolbox.StartTask()
	defer toolbox.StopTask()

	beego.Run()
}

// 初始化定时任务
func initTask() {
	//// 添加全局定时任务
	//ht := toolbox.NewTask("heatbeat", "*/1 * * * * *", func() error {
	//	fmt.Println("i am health");
	//	return nil
	//})
	//toolbox.AddTask("heatbeat", ht)
}
