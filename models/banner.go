package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
)

// banner
type Banner struct {
	Id          int64		`json:"id" form:"id" orm:"column(id)"`
	Title       string		`json:"title" form:"title" valid:"Required" orm:"column(title)"`
	Url         string		`json:"url" form:"url" valid:"Required" orm:"column(url)"`
	Link        string		`json:"link" form:"link" orm:"column(link)"`
	Background  string		`json:"background" form:"background" orm:"column(background)"`
	Sort        int32		`json:"sort" form:"sort" orm:"column(sort)"`
	Desc        string		`json:"desc" form:"desc" orm:"column(desc)"`
	Base
}

func init() {
	// 注册模型
	orm.RegisterModel(new(Banner))
}

// 自定义函数验证
func (u *Banner) Valid(v * validation.Validation) {

}

// 自定义表名
func (u *Banner) TableName() string {
	return "tbl_banner"
}

// 插入banner
func AddBanner(banner Banner) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(&banner)
	if err != nil {
		return 0, err
	}
	return id, nil
}


// 更新banner
func UpdateBanner(banner Banner) (int64, error) {
	o := orm.NewOrm()
	count, err := o.Update(&banner)
	if err != nil {
		return 0, err
	}
	return count, err
}

// 删除banner
func DeleteBanner(banner Banner) (int64, error) {
	o := orm.NewOrm()
	banner.Deleted = 1
	count, err := o.Update(&banner, "deleted")
	if err != nil {
		return 0, err
	}
	return count, err
}

// 查找banner
func GetBanner(id int64) (*Banner, error) {
	o := orm.NewOrm()
	banner := Banner{Id : id}
	err := o.Read(&banner)
	if err != nil {
		return nil, err
	}
	return &banner, nil
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



