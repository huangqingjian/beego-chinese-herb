package component

import "github.com/panjf2000/ants/v2"

// 异步任务
func Async(task func()) {
	ants.Submit(task)
}