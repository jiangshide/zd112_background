package models

import (
	"github.com/astaxie/beego/orm"
)

type Nation struct {
	Id         int
	Name       string
	Icon       string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (this *Nation) TableName() string {
	return TableName("nation")
}

func (this *Nation) Update(fields ...string) (int64, error) {
	return orm.NewOrm().Update(this, fields...)
}

func NationAdd(this *Nation) (int64, error) {
	return orm.NewOrm().Insert(this)
}

func NationList(pageSize,offSet int) ([]*Nation, int64) {
	list := make([]*Nation, 0)
	query := orm.NewOrm().QueryTable(TableName("nation"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return list, total
}

func NationById(id int) (this *Nation, err error) {
	this = &Nation{
		Id: id,
	}
	return this, orm.NewOrm().Read(this)
}

func NationDel(id int) (int64, error) {
	return orm.NewOrm().QueryTable(TableName("nation")).Filter("id", id).Delete()
}
