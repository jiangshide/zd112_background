package models

import "github.com/astaxie/beego/orm"

type Auth struct {
	Id         int
	Name       string
	Url        string
	UserId     int
	Pid        int
	Sort       int
	Icon       string
	IsShow     int
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (this *Auth) TableName() string {
	return TableName("auth")
}

func (this *Auth) Update(fields ...string) error {
	_, err := orm.NewOrm().Update(this, fields...)
	return err
}

func AuthAdd(this *Auth) (int64, error) {
	return orm.NewOrm().Insert(this)
}

func AuthList(page, pageSize int, filters ...interface{}) ([]*Auth, int64) {
	offSet := (page - 1) * pageSize
	list := make([]*Auth, 0)
	query := orm.NewOrm().QueryTable(TableName("auth"))
	if len(filters) > 0 {
		size := len(filters)
		for k := 0; k < size; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("pid", "sort").Limit(pageSize, offSet).All(&list)
	return list, total
}

func AuthGetById(id int)(this *Auth,err error){
	this = &Auth{
		Id:id,
	}
	return this,orm.NewOrm().Read(this)
}
