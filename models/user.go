package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	"context"
)

// 用户
type User struct {
	Id         int64		`json:"id" form:"id" orm:"column(id)"`							 // id
	Name       string		`json:"name" form:"name" valid:"Required" orm:"column(name)"`    // 姓名
	Mobile     string		`json:"mobile" form:"mobile" valid:"Mobile" orm:"column(mobile)"`// 手机号
	Email      string		`json:"email" form:"email" valid:"Email" orm:"column(email)"`	 // 邮箱
	Password   string		`json:"password" form:"password" orm:"column(password)"`		 // 密码
	Face       string		`json:"face" form:"face" orm:"column(face)"`					 // 头像
	Sex        int8			`json:"sex" form:"sex" valid:"Range(1, 2)" orm:"column(sex)"`	 // 性别
	Desc       string		`json:"desc" form:"desc" orm:"column(desc)"`					 // 描述
	Base
}

// 查询条件
type UserQuery struct {
	Q          string
	PageQuery
}

func init() {
	// 注册模型
	orm.RegisterModel(new(User))
	// 自定义表名前缀
	//orm.RegisterModelWithPrefix("tbl_", new(User))
}

// 自定义函数验证
func (u *User) Valid(v * validation.Validation) {

}

// 自定义表名
func (u *User) TableName() string {
	return "tbl_user"
}

// 插入用户
func AddUser(u *User) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(u)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 带事务式插入用户
func AddUserWithTrans(u *User) (int64, error) {
	o := orm.NewOrm()
	err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		id, err := o.Insert(u)
		u.Id = id
		return err
	})
	if err != nil {
		return 0, err
	}
	return u.Id, nil
}

// 更新用户
func UpdateUser(u *User) (int64, error) {
	o := orm.NewOrm()
	count, err := o.Update(u)
	if err != nil {
		return 0, err
	}
	return count, err
}

// 带事务式更新用户
func UpdateUserWithTrans(u *User) error {
	o := orm.NewOrm()
	err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		_, err := o.Update(u)
		return err
	})
	return err
}

// 删除用户
func DeleteUser(u *User) (int64, error) {
	o := orm.NewOrm()
	u.Deleted = 1
	count, err := o.Update(u, "deleted")
	if err != nil {
		return 0, err
	}
	return count, err
}

// 带事务式删除用户
func DeleteUserWithTrans(u *User) error {
	o := orm.NewOrm()
	err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		_, err := o.Update(u, "deleted")
		return err
	})
	return err
}

// 查找用户
func GetUser(id int64) (*User, error) {
	o := orm.NewOrm()
	user := User{Id : id}
	err := o.Read(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 查询用户
func GetUserList(query *UserQuery) (*Page, error) {
	var users []User
	o := orm.NewOrm()
	qs := o.QueryTable("tbl_user")
	qs = qs.Filter("deleted", 0)
	if query.Q != "" {
		cond := orm.NewCondition()
		cond2 := orm.NewCondition()
		cond2.And("name__startswith", query.Q).Or("mobile__startswith", query.Q)
		cond.AndCond(cond2)
		qs = qs.SetCond(cond)
	}
	qs = qs.Limit(query.PageSize, (query.PageNum - 1) * query.PageSize)
	count, err := qs.Count()
	if err != nil {
		return nil, err
	}
	_, err = qs.All(&users)
	if err != nil {
		return nil, err
	}
	return NewPage(query.PageNum, query.PageSize, int32(len(users)), int32(count),  &users), nil
}



