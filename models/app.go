package models

import (
	"github.com/astaxie/beego/orm"
)

type Channel struct {
	Id         int
	Name       string
	FriendId   string
	Descript   string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
	Views      int
}

func (this *Channel) TableName() string {
	return TableName("app_channel")
}

func (this *Channel) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Channel) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Channel) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Channel) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Channel) List(pageSize, offSet int) (list []*Channel, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Application struct {
	Id         int
	Name       string
	Descript   string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
	Views      int
}

func (this *Application) TableName() string {
	return TableName("app_application")
}

func (this *Application) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Application) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Application) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Application) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Application) List(pageSize, offSet int) (list []*Application, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Pkg struct {
	Id         int
	Name       string
	Descript   string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
	Views      int
}

func (this *Pkg) TableName() string {
	return TableName("app_pkg")
}

func (this *Pkg) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Pkg) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Pkg) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Pkg) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Pkg) List(pageSize, offSet int) (list []*Pkg, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Version struct {
	Id         int
	Name       string
	Descript   string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
	Views      int
}

func (this *Version) TableName() string {
	return TableName("app_version")
}

func (this *Version) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Version) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Version) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Version) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Version) List(pageSize, offSet int) (list []*Version, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Code struct {
	Id         int
	Code       int
	Descript   string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
	Views      int
}

func (this *Code) TableName() string {
	return TableName("app_code")
}

func (this *Code) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Code) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Code) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Code) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Code) List(pageSize, offSet int) (list []*Code, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Env struct {
	Id         int
	Name       string
	Descript   string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
	Views      int
}

func (this *Env) TableName() string {
	return TableName("app_env")
}

func (this *Env) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Env) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Env) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Env) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Env) List(pageSize, offSet int) (list []*Env, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Build struct {
	Id         int
	Name       string
	Descript   string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
	Views      int
}

func (this *Build) TableName() string {
	return TableName("app_build")
}

func (this *Build) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Build) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Build) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Build) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Build) List(pageSize, offSet int) (list []*Build, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(list)
	return
}

type Type struct {
	Id         int
	Name       string
	Descript   string
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
	Views      int
}

func (this *Type) TableName() string {
	return TableName("app_type")
}

func (this *Type) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Type) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Type) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Type) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Type) List(pageSize, offSet int) (list []*Type, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type App struct {
	Id          int
	ProjectId   int
	TestId      int
	Icon        string
	Type        string
	Application string
	Pkg         string
	Version     string
	Code        int
	Env         string
	Build       string
	Channel     string
	FriendId    string
	Descript    string
	Status      int
	Times       int
	Url         string
	Downs       int
	CreateId    int
	UpdateId    int
	CreateTime  int64
	UpdateTime  int64
	Views       int
}

func (this *App) TableName() string {
	return TableName("app")
}

func (this *App) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *App) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *App) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *App) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *App) List(pageSize, offSet int) (list []*App, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return list, total
}
