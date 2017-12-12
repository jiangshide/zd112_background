package controllers

import (
	"zd112/models"
	"time"
)

type ChannelController struct {
	BaseController
}

func (this *BaseController) row(row map[string]interface{}, id int, name, descript string) {
	row["id"] = id
	row["name"] = name
	row["descript"] = descript
	this.Data["row"] = row
}

func (this *ChannelController) List() {
	this.pageTitle("渠道列表")
	this.display(this.getBgAppAction("channel/list"))
}

func (this *ChannelController) Add() {
	this.pageTitle("增加渠道名称")
	this.display(this.getBgAppAction("channel/add"))
}

func (this *ChannelController) Edit() {
	this.pageTitle("编辑渠道名称")
	channel := new(models.Channel)
	channel.Id = this.getInt("id", 0)
	if err := channel.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{})
	row["friend_id"] = channel.FriendId
	this.row(row, channel.Id, channel.Name, channel.Descript)
	this.display(this.getBgAppAction("channel/edit"))
}

func (this *ChannelController) AjaxSave() {
	channel := new(models.Channel)
	channel.Id = this.getInt("id", 0)
	channel.Name = this.getString("name", "名称不能为空!", 1)
	channel.FriendId = this.getString("friend_id", "", 0)
	channel.Descript = this.getString("descript", "", 0)
	var err error
	if channel.Id == 0 {
		channel.CreateId = this.userId
		channel.CreateTime = time.Now().Unix()
		_, err = channel.Add()
	} else {
		channel.UpdateId = this.userId
		channel.UpdateTime = time.Now().Unix()
		_, err = channel.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *ChannelController) Table() {
	channel := new(models.Channel)
	result, count := channel.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["name"] = v.Name
		row["friend_id"] = v.FriendId
		row["descript"] = v.Descript
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *ChannelController) AjaxDel() {
	channel := new(models.Channel)
	channel.Id = this.getInt("id", 0)
	if _, err := channel.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type ApplicationController struct {
	BaseController
}

func (this ApplicationController) List() {
	this.pageTitle("应用名称列表")
	this.display(this.getBgAppAction("application/list"))
}

func (this ApplicationController) Add() {
	this.pageTitle("增加应用名称")
	this.display(this.getBgAppAction("application/add"))
}

func (this ApplicationController) Edit() {
	this.pageTitle("编辑应用名称")
	application := new(models.Application)
	application.Id = this.getInt("id", 0)
	if err := application.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{})
	this.row(row, application.Id, application.Name, application.Descript)
	this.display(this.getBgAppAction("application/edit"))
}

func (this ApplicationController) AjaxSave() {
	application := new(models.Application)
	application.Id = this.getInt("id", 0)
	application.Name = this.getString("name", "应用名称不能为空!", 1)
	application.Descript = this.getString("descript", "", 0)
	var err error
	if application.Id == 0 {
		application.CreateId = this.userId
		application.CreateTime = time.Now().Unix()
		_, err = application.Add()
	} else {
		application.UpdateId = this.userId
		application.UpdateTime = time.Now().Unix()
		_, err = application.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this ApplicationController) Table() {
	application := new(models.Application)
	result, count := application.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["name"] = v.Name
		row["descript"] = v.Descript
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *ApplicationController) AjaxDel() {
	application := new(models.Application)
	application.Id = this.getInt("id", 0)
	if _, err := application.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type PkgController struct {
	BaseController
}

func (this PkgController) List() {
	this.pageTitle("应用包列表")
	this.display(this.getBgAppAction("pkg/list"))
}

func (this PkgController) Add() {
	this.pageTitle("增加应用包")
	this.display(this.getBgAppAction("pkg/add"))
}

func (this PkgController) Edit() {
	this.pageTitle("编辑应用包")
	pkg := new(models.Pkg)
	pkg.Id = this.getInt("id", 0)
	if err := pkg.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{})
	this.row(row, pkg.Id, pkg.Name, pkg.Descript)
	this.display(this.getBgAppAction("pkg/edit"))
}

func (this *PkgController) AjaxSave() {
	pkg := new(models.Pkg)
	pkg.Id = this.getInt("id", 0)
	pkg.Name = this.getString("name", "包名称不能为空!", 1)
	pkg.Descript = this.getString("descript", "", 0)
	var err error
	if pkg.Id == 0 {
		pkg.CreateId = this.userId
		pkg.CreateTime = time.Now().Unix()
		_, err = pkg.Add()
	} else {
		pkg.UpdateId = this.userId
		pkg.UpdateTime = time.Now().Unix()
		_, err = pkg.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *PkgController) Table() {
	pkg := new(models.Pkg)
	result, count := pkg.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["name"] = v.Name
		row["descript"] = v.Descript
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *PkgController) AjaxDel() {
	pkg := new(models.Pkg)
	pkg.Id = this.getInt("id", 0)
	if _, err := pkg.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type VersionController struct {
	BaseController
}

func (this *VersionController) List() {
	this.pageTitle("应用版本列表")
	this.display(this.getBgAppAction("version/list"))
}

func (this *VersionController) Add() {
	this.pageTitle("增加应用版本")
	this.display(this.getBgAppAction("version/add"))
}

func (this *VersionController) Edit() {
	this.pageTitle("编辑应用版本")
	version := new(models.Version)
	version.Id = this.getInt("id", 0)
	if err := version.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{})
	this.row(row, version.Id, version.Name, version.Descript)
	this.display(this.getBgAppAction("version/edit"))
}

func (this *VersionController) AjaxSave() {
	version := new(models.Version)
	version.Id = this.getInt("id", 0)
	version.Name = this.getString("name", "版本名称不能为空!", 1)
	version.Descript = this.getString("descript", "", 0)
	var err error
	if version.Id == 0 {
		version.CreateId = this.userId
		version.CreateTime = time.Now().Unix()
		_, err = version.Add()
	} else {
		version.UpdateId = this.userId
		version.UpdateTime = time.Now().Unix()
		_, err = version.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *VersionController) Table() {
	version := new(models.Version)
	result, count := version.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row ["name"] = v.Name
		row["descript"] = v.Descript
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *VersionController) AjaxDel() {
	version := new(models.Version)
	version.Id = this.getInt("id", 0)
	if _, err := version.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type CodeController struct {
	BaseController
}

func (this *CodeController) List() {
	this.pageTitle("应用编号列表")
	this.display(this.getBgAppAction("code/list"))
}

func (this *CodeController) Add() {
	this.pageTitle("增加应用编号列表")
	this.display(this.getBgAppAction("code/add"))
}

func (this *CodeController) Edit() {
	this.pageTitle("编辑应用编号")
	code := new(models.Code)
	code.Id = this.getInt("id", 0)
	if err := code.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{})
	row["id"] = code.Id
	row["code"] = code.Code
	row["descript"] = code.Descript
	this.Data["row"] = row
	this.display(this.getBgAppAction("code/edit"))
}

func (this *CodeController) AjaxSave() {
	code := new(models.Code)
	code.Id = this.getInt("id", 0)
	code.Code = this.getInt("code", 0)
	code.Descript = this.getString("descript", "", 0)
	var err error
	if code.Id == 0 {
		code.CreateId = this.userId
		code.CreateTime = time.Now().Unix()
		_, err = code.Add()
	} else {
		code.UpdateId = this.userId
		code.UpdateTime = time.Now().Unix()
		_, err = code.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *CodeController) Table() {
	code := new(models.Code)
	result, count := code.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["code"] = v.Code
		row["descript"] = v.Descript
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *CodeController) AjaxDel() {
	code := new(models.Code)
	code.Id = this.getInt("id", 0)
	if _, err := code.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type EnvController struct {
	BaseController
}

func (this *EnvController) List() {
	this.pageTitle("应用环境列表")
	this.display(this.getBgAppAction("env/list"))
}

func (this *EnvController) Add() {
	this.pageTitle("增加应用环境")
	this.display(this.getBgAppAction("env/add"))
}

func (this *EnvController) Edit() {
	this.pageTitle("编辑应用环境")
	env := new(models.Env)
	env.Id = this.getInt("id", 0)
	if err := env.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{})
	this.row(row, env.Id, env.Name, env.Descript)
	this.display(this.getBgAppAction("env/edit"))
}

func (this *EnvController) AjaxSave() {
	env := new(models.Env)
	env.Id = this.getInt("id", 0)
	env.Name = this.getString("name", "应用环境名称不能为空！", 1)
	env.Descript = this.getString("descript", "", 0)
	var err error
	if env.Id == 0 {
		env.CreateId = this.userId
		env.CreateTime = time.Now().Unix()
		_, err = env.Add()
	} else {
		env.UpdateId = this.userId
		env.UpdateTime = time.Now().Unix()
		_, err = env.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *EnvController) Table() {
	env := new(models.Env)
	result, count := env.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["name"] = v.Name
		row["descript"] = v.Descript
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *EnvController) AjaxDel() {
	env := new(models.Env)
	env.Id = this.getInt("id", 0)
	if _, err := env.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type BuildController struct {
	BaseController
}

func (this *BuildController) List() {
	this.pageTitle("构建类型列表")
	this.display(this.getBgAppAction("build/list"))
}

func (this *BuildController) Add() {
	this.pageTitle("增加构建类型")
	this.display(this.getBgAppAction("build/add"))
}

func (this *BuildController) Edit() {
	this.pageTitle("编辑构建类型")
	build := new(models.Build)
	build.Id = this.getInt("id", 0)
	if err := build.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{})
	this.row(row, build.Id, build.Name, build.Descript)
	this.display(this.getBgAppAction("build/edit"))
}

func (this *BuildController) AjaxSave() {
	build := new(models.Build)
	build.Id = this.getInt("id", 0)
	build.Name = this.getString("name", "名称不能为空!", 1)
	build.Descript = this.getString("descript", "", 0)
	var err error
	if build.Id == 0 {
		build.CreateId = this.userId
		build.CreateTime = time.Now().Unix()
		_, err = build.Add()
	}
	build.UpdateId = this.userId
	build.UpdateTime = time.Now().Unix()
	_, err = build.Update()
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *BuildController) Table() {
	build := new(models.Build)
	result, count := build.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["name"] = v.Name
		row["descript"] = v.Descript
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *BuildController) AjaxDel() {
	build := new(models.Build)
	build.Id = this.getInt("id", 0)
	if _, err := build.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type TypeController struct {
	BaseController
}

func (this *TypeController) List() {
	this.pageTitle("构建平台列表")
	this.display(this.getBgAppAction("type/list"))
}

func (this *TypeController) Add() {
	this.pageTitle("增加平台名称")
	this.display(this.getBgAppAction("type/add"))
}

func (this *TypeController) Edit() {
	this.pageTitle("编辑平台名称")
	types := new(models.Type)
	types.Id = this.getInt("id", 0)
	if err := types.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{})
	this.row(row, types.Id, types.Name, types.Descript)
	this.display(this.getBgAppAction("type/edit"))
}

func (this *TypeController) AjaxSave() {
	types := new(models.Type)
	types.Id = this.getInt("id", 0)
	types.Name = this.getString("name", "名称不能为空!", 1)
	types.Descript = this.getString("descript", "", 0)
	var err error
	if types.Id == 0 {
		types.CreateId = this.userId
		types.CreateTime = time.Now().Unix()
		_, err = types.Add()
	} else {
		types.UpdateId = this.userId
		types.UpdateTime = time.Now().Unix()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *TypeController) Table() {
	types := new(models.Type)
	result, count := types.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["name"] = v.Name
		row["descript"] = v.Descript
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *TypeController) AjaxDel() {
	types := new(models.Type)
	types.Id = this.getInt("id", 0)
	if _, err := types.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type AppController struct {
	BaseController
}

func (this *AppController) List() {
	this.pageTitle("应用列表")
	this.display(this.getBgAppAction("app/list"))
}

func (this *AppController) Add() {
	this.pageTitle("增加应用")
	this.display(this.getBgAppAction("app/add"))
}

func (this *AppController) Edit() {
	app := new(models.App)
	app.Id = this.getInt("id", 0)
	if err := app.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{})
	row["id"] = app.Id
	row["project_id"] = app.ProjectId
	row["test_id"] = app.TestId
	row["icon"] = app.Icon
	row["type"] = app.Type
	row["application"] = app.Application
	row["pkg"] = app.Pkg
	row["version"] = app.Version
	row["code"] = app.Code
	row["env"] = app.Env
	row["build"] = app.Build
	row["channel"] = app.Channel
	row["friend_id"] = app.FriendId
	row["descript"] = app.Descript
	row["status"] = app.Status
	row["times"] = app.Times
	row["url"] = app.Url
	row["downs"] = app.Downs
	this.Data["row"] = row
	this.display(this.getBgAppAction("app/edit"))
}

func (this *AppController) AjaxSave() {
	app := new(models.App)
	app.Id = this.getInt("id", 0)
	app.TestId = this.getInt("test_id", 0)
	app.Icon = this.getString("icon", "Logo不能为空!", 1)
	app.Type = this.getString("type", "构建平台不能为空", 1)
	app.Application = this.getString("application", "应用名称不能为空!", 1)
	app.Pkg = this.getString("pkg", "应用包名不能为空!", 1)
	app.Version = this.getString("version", "应用版本不能为空!", 1)
	app.Code = this.getInt("code", 0)
	app.Env = this.getString("env", "构建应用环境不能为空", 1)
	app.Build = this.getString("build", "构建应用类型不能为空", 1)
	app.Channel = this.getString("channel", "构建应用渠道不能为空!", 1)
	app.FriendId = this.getString("friend_id", "FriendId不能为空!", 1)
	app.Descript = this.getString("descript", "", 0)
	app.Status = this.getInt("status", 0)
	app.Times = this.getInt("times", 0)
	app.Url = this.getString("url", "", 0)
	app.Downs = this.getInt("downs", 0)
	var err error
	if app.Id == 0 {
		app.CreateId = this.userId
		app.CreateTime = time.Now().Unix()
		_, err = app.Add()
	} else {
		app.UpdateId = this.userId
		app.UpdateTime = time.Now().Unix()
		_, err = app.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *AppController) Table() {
	app := new(models.App)
	result, count := app.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["test_id"] = v.TestId
		row["icon"] = v.Icon
		row["type"] = v.Type
		row["application"] = v.Application
		row["pkg"] = v.Pkg
		row["version"] = v.Version
		row["code"] = v.Code
		row["env"] = v.Env
		row["build"] = v.Build
		row["channel"] = v.Channel
		row["friend_id"] = v.FriendId
		row["descript"] = v.Descript
		row["status"] = v.Status
		row["times"] = v.Times
		row["url"] = v.Url
		row["downs"] = v.Downs
		this.parse(list, row, k, v.Id, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *AppController) AjaxDel() {
	app := new(models.App)
	app.Id = this.getInt("id", 0)
	if _, err := app.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}
