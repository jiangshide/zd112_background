package models

import (
	"github.com/astaxie/beego/orm"
)

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
	return TableName("uc_auth")
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
	query := orm.NewOrm().QueryTable(TableName("uc_auth"))
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

func (this *Auth) List(pageSize, offSet int) (list []*Auth, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

func AuthGetById(id int) (this *Auth, err error) {
	this = &Auth{
		Id: id,
	}
	return this, orm.NewOrm().Read(this)
}

func AuthDelById(id int) (int64, error) {
	return orm.NewOrm().QueryTable(TableName("uc_auth")).Filter("id", id).Delete()
}

func AuthGetListByIds(authIds string, userId int) ([]*Auth, error) {
	list := make([]*Auth, 0)
	var lists []orm.Params
	var err error
	if userId == 1 {
		_, err = orm.NewOrm().Raw("select id,name,url,pid,sort,icon,is_show from zd_uc_auth where status=? order by pid asc,sort asc", 1).Values(&lists)
	} else {
		_, err = orm.NewOrm().Raw("select id,name,url,pid,sort,icon,is_show from zd_uc_auth where status=1 and id in("+authIds+") order by pid asc,sort asc", authIds).Values(&lists)
	}
	return list, err
}
