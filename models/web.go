package models

import "github.com/astaxie/beego/orm"

type Banner struct {
	Id         int
	Name       string
	Url        string
	Icon       string
	Descript   string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (this *Banner) TableName() string {
	return TableName("web_banner")
}

func (this *Banner) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Banner) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Banner) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Banner) Query() error {
	return orm.NewOrm().Read(this)
}

func BannerList(pageSize, offSet int) ([]*Banner, int64) {
	list := make([]*Banner, 0)
	query := orm.NewOrm().QueryTable(TableName("web_banner"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return list, total
}
