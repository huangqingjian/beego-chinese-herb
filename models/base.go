package models

type Base struct {
	CreateBy   int64		`json:"createBy" orm:"column(create_by)"`
	CreateTime Time			`json:"createTime" orm:"auto_now_add;column(create_time)"`
	UpdateBy   int64		`json:"updateBy" orm:"column(update_by)"`
	UpdateTime Time			`json:"updateTime" orm:"auto_now;column(update_time)"`
	Deleted    int8			`json:"deleted" orm:"column(deleted)"`
}