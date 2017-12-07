package models

import "github.com/astaxie/beego/orm"

type FormatType struct {
	Id         int
	Name       string
	Descript   string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (this *FormatType) TableName() string {
	return TableName("tools_format_type")
}

func (this *FormatType) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *FormatType) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *FormatType) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *FormatType) Query() error {
	return orm.NewOrm().Read(this)
}

func FormatTypeList(pageSize, offSet int) ([]*FormatType, int64) {
	list := make([]*FormatType, 0)
	query := orm.NewOrm().QueryTable(TableName("tools_format_type"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return list, total
}

type Format struct {
	Id         int
	ParentId   int
	Name       string
	Descript   string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (this *Format) TableName() string {
	return TableName("tools_format")
}

func (this *Format) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Format) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Format) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Format) Query() error {
	return orm.NewOrm().Read(this)
}

func FormatList(pageSize, offSet int) ([]*Format, int64) {
	list := make([]*Format, 0)
	query := orm.NewOrm().QueryTable(TableName("tools_format"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return list, total
}

type Compress struct {
	Id         int
	Name       string
	Url        string
	Type       int
	Format     string
	Descript   string
	Size       int64
	ReSize     int
	Compress   int
	Downs      int
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
	Views      int
}

func (this *Compress) TableName() string {
	return TableName("tools_compress")
}

func (this *Compress) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Compress) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Compress) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Compress) Query() error {
	return orm.NewOrm().Read(this)
}

func CompressList(pageSize, offSet int) ([]*Compress, int64) {
	list := make([]*Compress, 0)
	query := orm.NewOrm().QueryTable(TableName("tools_compress"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return list, total
}
