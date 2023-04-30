package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
)

// banner
type Banner struct {
	Id          int64		`json:"id" form:"id" orm:"column(id)"`							// id
	Title       string		`json:"title" form:"title" valid:"Required" orm:"column(title)"`// 标题
	Url         string		`json:"url" form:"url" valid:"Required" orm:"column(url)"`		// 图片url
	Link        string		`json:"link" form:"link" orm:"column(link)"`					// 跳转链接
	Background  string		`json:"background" form:"background" orm:"column(background)"`	// 背景
	Sort        int32		`json:"sort" form:"sort" orm:"column(sort)"`					// 顺序
	Desc        string		`json:"desc" form:"desc" orm:"column(desc)"`					// 描述
	Base
}

func init() {
	// 注册模型
	orm.RegisterModel(new(Banner))
}

// 自定义函数验证
func (b *Banner) Valid(v * validation.Validation) {

}

// 自定义表名
func (b *Banner) TableName() string {
	return "tbl_banner"
}

// 插入banner
func AddBanner(b *Banner) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(b)
	if err != nil {
		return 0, err
	}
	return id, nil
}


// 更新banner
func UpdateBanner(b *Banner) (int64, error) {
	o := orm.NewOrm()
	count, err := o.Update(b)
	if err != nil {
		return 0, err
	}
	return count, err
}

// 删除banner
func DeleteBanner(b *Banner) (int64, error) {
	o := orm.NewOrm()
	b.Deleted = 1
	count, err := o.Update(b, "deleted")
	if err != nil {
		return 0, err
	}
	return count, err
}

// 查找banner
func GetBanner(id int64) (*Banner, error) {
	o := orm.NewOrm()
	b := Banner{Id : id}
	err := o.Read(&b)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

// 查询banner
func GetBannerList() ([]Banner, error) {
	var banners []Banner
	o := orm.NewOrm()
	qs := o.QueryTable("tbl_banner")
	qs = qs.Filter("deleted", 0)
	_, err := qs.All(&banners)
	if err != nil {
		return nil, err
	}
	return banners, nil
}



