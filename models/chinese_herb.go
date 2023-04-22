package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
)

// 中药
type ChineseHerb struct {
	Id          int64		`json:"id" form:"id" orm:"column(id)"`
	Type        int64		`json:"type" valid:"Required"`
	Name        string		`json:"name" valid:"Required"`
	EnName      string		`json:"enName" form:"enName" orm:"column(en_name)"`
	Alias       string		`json:"alias"`
	Pic         string		`json:"pic" valid:"Required"`
	Zwxt        string		`json:"zwxt"`
	Yybw        string		`json:"yybw"`
	Cdfb        string		`json:"cdfb"`
	Csjg        string		`json:"csjg"`
	Ycxz        string		`json:"ycxz"`
	Xwgj        string		`json:"xwgj"`
	Gxzy        string		`json:"gxzy"`
	Lcyy        string		`json:"lcyy"`
	Ylyj        string		`json:"ylyj"`
	Hxcf        string		`json:"hxcf"`
	Syjj        string		`json:"syjj"`
	HerbPharmacys []HerbPharmacy `json:"herbPharmacys" orm:"-"`
	Base
}

// 查询查询
type ChineseHerbQuery struct {
	Q           string
	Type        int64
	PageQuery
}

func init() {
	// 注册模型
	orm.RegisterModel(new(ChineseHerb))
}

// 自定义函数验证
func (c *ChineseHerb) Valid(v * validation.Validation) {

}

// 自定义表名
func (c *ChineseHerb) TableName() string {
	return "tbl_chinese_herb"
}

// 插入中药
func AddChineseHerb(c *ChineseHerb) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(c)
	if err != nil {
		return 0, err
	}
	return id, nil
}


// 更新中药
func UpdateChineseHerb(c *ChineseHerb) (int64, error) {
	o := orm.NewOrm()
	count, err := o.Update(c)
	if err != nil {
		return 0, err
	}
	return count, err
}

// 删除中药
func DeleteChineseHerb(c *ChineseHerb) (int64, error) {
	o := orm.NewOrm()
	c.Deleted = 1
	count, err := o.Update(c, "deleted")
	if err != nil {
		return 0, err
	}
	return count, err
}

// 查找中药
func GetChineseHerb(id int64) (*ChineseHerb, error) {
	o := orm.NewOrm()
	c := ChineseHerb{Id : id}
	err := o.Read(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// 查询中药
func GetChineseHerbList(query *ChineseHerbQuery) (*Page, error) {
	var herbs []ChineseHerb
	o := orm.NewOrm()
	qs := o.QueryTable("tbl_chinese_herb")
	qs = qs.Filter("deleted", 0)
	if query.Type != 0 {
		qs = qs.Filter("type", query.Type)
	}
	if query.Q != "" {
		qs = qs.Filter("name__startswith", query.Q)
	}
	qs = qs.Limit(query.PageSize, (query.PageNum - 1) * query.PageSize)
	count, err := qs.Count()
	if err != nil {
		return nil, err
	}
	_, err = qs.All(&herbs)
	if err != nil {
		return nil, err
	}
	return NewPage(query.PageNum, query.PageSize, int32(len(herbs)), int32(count), &herbs), nil
}



