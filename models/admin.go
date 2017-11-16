package models

import "github.com/astaxie/beego/orm"

type Amin struct{
	Id int
	LoginName string
	RealName string
	Password string
	RoleIds string
	Phone string
	Email string
	Salt string
	LastLogin int64
	LastIp string
	Status int
	CreateId int
	UpdateId int
	CreateTime int64
	UpdateTime int
}

func (this *Amin)TableName()string{
	return TableName("admin")
}

func AdminAdd(this *Amin)(int64,error){
	return orm.NewOrm().Insert(this)
}