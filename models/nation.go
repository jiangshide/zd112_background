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

func (this *Nation) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Nation) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Nation) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Nation) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter("name", this.Name).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Nation) List(pageSize, offSet int) ([]*Nation, int64) {
	list := make([]*Nation, 0)
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return list, total
}
