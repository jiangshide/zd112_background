package models

import "github.com/astaxie/beego/orm"

type Role struct {
	Id         int
	Name       string
	Detail     string
	Status     string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (this *Role) TableName() string {
	return TableName("role")
}

func (this *Role) Update(fields ...string) error {
	_, err := orm.NewOrm().Update(this, fields...)
	return err
}

func RoleAdd(this *Role) (int64, error) {
	return orm.NewOrm().Insert(this)
}

func RoleList(page, pageSize int, filters ...interface{}) ([]*Role, int64) {
	offSet := (page - 1) * pageSize
	list := make([]*Role, 0)
	query := orm.NewOrm().QueryTable(TableName("role"))
	if len(filters) > 0 {
		size := len(filters)
		for k := 0; k < size; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(page, offSet).All(&list)
	return list, total
}