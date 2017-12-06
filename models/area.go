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

func (this *Continent) TableName() string {
	return TableName("area_continent")
}

func (this *Continent) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Continent) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Continent) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Continent) Query() error {
	return orm.NewOrm().Read(this)
}

func ContinentList(pageSize, offSet int) ([]*Continent, int64) {
	list := make([]*Continent, 0)
	query := orm.NewOrm().QueryTable(TableName("area_continent"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return list, total
}

type State struct {
	Id         int
	ParentId   int
	Name       string
	Icon       string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (this *State) TableName() string {
	return TableName("area_state")
}

func (this *State) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *State) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *State) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *State) Query() error {
	return orm.NewOrm().Read(this)
}

func StateList(pageSize, offSet int) ([]*State, int64) {
	list := make([]*State, 0)
	query := orm.NewOrm().QueryTable(TableName("area_state"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return list, total
}

type Province struct {
	Id         int
	ParentId   int
	Name       string
	Icon       string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (this *Province) TableName() string {
	return TableName("area_province")
}

func (this *Province) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Province) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Province) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Province) Query() error {
	return orm.NewOrm().Read(this)
}

func ProvinceList(pageSize, offSet int) ([]*Province, int64) {
	list := make([]*Province, 0)
	query := orm.NewOrm().QueryTable(TableName("area_province"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return list, total
}

type City struct {
	Id         int
	ParentId   int
	Name       string
	Icon       string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (this *City) TableName() string {
	return TableName("area_city")
}

func (this *City) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *City) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *City) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *City) Query() error {
	return orm.NewOrm().Read(this)
}

func CityList(pageSize, offSet int) ([]*City, int64) {
	list := make([]*City, 0)
	query := orm.NewOrm().QueryTable(TableName("area_city"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return list, total
}

type Region struct {
	Id         int
	ParentId   int
	Name       string
	Icon       string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (this *Region) TableName() string {
	return TableName("area_region")
}

func (this *Region) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Region) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Region) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Region) Query() error {
	return orm.NewOrm().Read(this)
}

func RegionList(pageSize, offSet int) ([]*Region, int64) {
	list := make([]*Region, 0)
	query := orm.NewOrm().QueryTable(TableName("area_region"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return list, total
}

type County struct {
	Id         int
	ParentId   int
	Name       string
	Icon       string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (this *County) TableName() string {
	return TableName("area_county")
}

func (this *County) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *County) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *County) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *County) Query() error {
	return orm.NewOrm().Read(this)
}

func CountyList(pageSize, offSet int) ([]*County, int64) {
	list := make([]*County, 0)
	query := orm.NewOrm().QueryTable(TableName("area_county"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return list, total
}

type Town struct {
	Id         int
	ParentId   int
	Name       string
	Icon       string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (this *Town) TableName() string {
	return TableName("area_town")
}

func (this *Town) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Town) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Town) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Town) Query() error {
	return orm.NewOrm().Read(this)
}

func TownList(pageSize, offSet int) ([]*Town, int64) {
	list := make([]*Town, 0)
	query := orm.NewOrm().QueryTable(TableName("area_town"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return list, total
}

type Country struct {
	Id         int
	ParentId   int
	Name       string
	Icon       string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (this *Country) TableName() string {
	return TableName("area_country")
}

func (this *Country) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Country) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Country) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Country) Query() error {
	return orm.NewOrm().Read(this)
}

func CountryList(pageSize, offSet int) ([]*Country, int64) {
	list := make([]*Country, 0)
	query := orm.NewOrm().QueryTable(TableName("area_country"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return list, total
}

type Village struct {
	Id         int
	ParentId   int
	Name       string
	Icon       string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (this *Village) TableName() string {
	return TableName("area_village")
}

func (this *Village) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Village) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Village) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Village) Query() error {
	return orm.NewOrm().Read(this)
}

func VillageList(pageSize, offSet int) ([]*Village, int64) {
	list := make([]*Village, 0)
	query := orm.NewOrm().QueryTable(TableName("area_village"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return list, total
}

type Group struct {
	Id         int
	ParentId   int
	Name       string
	Icon       string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (this *Group) TableName() string {
	return TableName("area_group")
}

func (this *Group) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Group) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Group) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Group) Query() error {
	return orm.NewOrm().Read(this)
}

func GroupList(pageSize, offSet int) ([]*Group, int64) {
	list := make([]*Group, 0)
	query := orm.NewOrm().QueryTable(TableName("area_group"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return list, total
}

type Team struct {
	Id         int
	ParentId   int
	Name       string
	Icon       string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (this *Team) TableName() string {
	return TableName("area_team")
}

func (this *Team) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Team) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Team) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Team) Query() error {
	return orm.NewOrm().Read(this)
}

func TeamList(pageSize, offSet int) ([]*Team, int64) {
	list := make([]*Team, 0)
	query := orm.NewOrm().QueryTable(TableName("area_team"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return list, total
}
