package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
)

// 药方
type HerbPharmacy struct {
	Id          int64		`json:"id"`												// id
	HerbId      int64		`json:"herbId" valid:"Required" orm:"column(herb_id)"`  // 中药Id
	Content     string		`json:"content"`										// 内容
	Base
}

// 药方查询
type HerbPharmacyQuery struct {
	HerbId      int64		// 中药Id
	HerbIds     []int64     // 中药Id（多个）
}

func init() {
	// 注册模型
	orm.RegisterModel(new(HerbPharmacy))
}

// 自定义函数验证
func (h *HerbPharmacy) Valid(v * validation.Validation) {

}

// 自定义表名
func (h *HerbPharmacy) TableName() string {
	return "tbl_herb_pharmacy"
}

// 插入药方
func AddHerbPharmacy(h *HerbPharmacy) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(h)
	if err != nil {
		return 0, err
	}
	return id, nil
}


// 更新药方
func UpdateHerbPharmacy(h *HerbPharmacy) (int64, error) {
	o := orm.NewOrm()
	count, err := o.Update(h)
	if err != nil {
		return 0, err
	}
	return count, err
}

// 删除药方
func DeleteHerbPharmacy(h *HerbPharmacy) (int64, error) {
	o := orm.NewOrm()
	h.Deleted = 1
	count, err := o.Update(h, "deleted")
	if err != nil {
		return 0, err
	}
	return count, err
}

// 查找药方
func GetHerbPharmacy(id int64) (*HerbPharmacy, error) {
	o := orm.NewOrm()
	herbPharmacy := HerbPharmacy{Id : id}
	err := o.Read(&herbPharmacy)
	if err != nil {
		return nil, err
	}
	return &herbPharmacy, nil
}

// 通过药物Id查询药方
func GetHerbPharmacyList(query *HerbPharmacyQuery) ([]HerbPharmacy, error) {
	var herbPharmacys []HerbPharmacy
	o := orm.NewOrm()
	qs := o.QueryTable("tbl_herb_pharmacy")
	qs = qs.Filter("deleted", 0)
	if query.HerbId != 0 {
		qs = qs.Filter("herbId", query.HerbId)
	}
	if query.HerbIds != nil && len(query.HerbIds) > 0 {
		qs = qs.Filter("herbId__in", query.HerbIds)
	}
	_, err := qs.All(&herbPharmacys)
	if err != nil {
		return nil, err
	}
	return herbPharmacys, nil
}



