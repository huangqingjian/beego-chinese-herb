package models

type Base struct {
	CreateBy   int64		`json:"createBy" orm:"column(create_by)"`					// 创建人
	CreateTime Time			`json:"createTime" orm:"auto_now_add;column(create_time)"`  // 创建时间
	UpdateBy   int64		`json:"updateBy" orm:"column(update_by)"`                   // 更新人
	UpdateTime Time			`json:"updateTime" orm:"auto_now;column(update_time)"`      // 更新时间
	Deleted    int8			`json:"deleted" orm:"column(deleted)"`						// 是否删除
}