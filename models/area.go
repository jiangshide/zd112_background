package models

import "github.com/astaxie/beego/orm"

type Continent struct {
	Id         int
	Name       string
	Icon       string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (this *Continent) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Continent) Del(id int) {
	this.Id = id
}

func (this *Continent) Update() {

}

func (this *Continent) Query() {

}

type State struct {
	Id          int
	ContinentId int
	Name        string
	Icon        string
	CreateId    int
	UpdateId    int
	CreateTime  int64
	UpdateTime  int64
}

func (this *State) Add() {

}

type Province struct {
	Id         int
	StateId    int
	Name       string
	Icon       string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

type City struct {
	Id         int
	ProvinceId int
	Name       string
	Icon       string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

type Region struct {
	Id         int
	CityId     int
	Name       string
	Icon       string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

type County struct {
	Id         int
	RegionId   int
	Name       string
	Icon       string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

type Town struct {
	Id         int
	CountyId   int
	Name       string
	Icon       string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

type Country struct {
	Id         int
	TownId     int
	Name       string
	Icon       string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

type Village struct {
	Id         int
	CountryId  int
	Name       string
	Icon       string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

type Group struct {
	Id         int
	VillageId  int
	Name       string
	Icon       string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

type Team struct {
	Id         int
	GroupId    int
	Name       string
	Icon       string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}
