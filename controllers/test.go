package controllers

import (
	"zd112/models"
	"time"
)

type EnvironmnetController struct {
	BaseController
}

func (this *EnvironmnetController) List() {
	this.pageTitle("系统环境列表")
	this.display(this.getBgTestAction("environment/list"))
}

func (this *EnvironmnetController) Add() {
	this.pageTitle("增加系统环境变量")
	this.display(this.getBgTestAction("environment/add"))
}

func (this *EnvironmnetController) Edit() {
	this.pageTitle("编辑系统环境变量")
	environment := new(models.Environment)
	environment.Id = this.getInt("id", 0)
	if err := environment.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{})
	row["id"] = environment.Id
	row["name"] = environment.Name
	row["jdk"] = environment.Jdk
	row["git"] = environment.Git
	row["gradle"] = environment.Gradle
	row["adb"] = environment.Adb
	this.Data["row"] = row
	this.display(this.getBgTestAction("environment/edit"))
}

func (this *EnvironmnetController) AjaxSave() {
	environment := new(models.Environment)
	environment.Id = this.getInt("id", 0)
	environment.Name = this.getString("name", "环境名称不能为空!", 2)
	environment.Jdk = this.getString("jdk", "jdk没有安装!", 2)
	environment.Git = this.getString("git", "git没有安装!", 2)
	environment.Gradle = this.getString("gradle", "Gradle没有安装!", 2)
	environment.Adb = this.getString("adb", "Adb没有安装!", 2)
	var err error
	if environment.Id == 0 {
		environment.CreateId = this.userId
		environment.CreateTime = time.Now().Unix()
		_, err = environment.Add()
	} else {
		environment.UpdateId = this.userId
		environment.UpdateTime = time.Now().Unix()
		_, err = environment.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *EnvironmnetController) Table() {
	environment := new(models.Environment)
	result, count := environment.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["name"] = v.Name
		row["jdk"] = v.Jdk
		row["git"] = v.Git
		row["gradle"] = v.Gradle
		row["adb"] = v.Adb
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *EnvironmnetController) AjaxDel() {
	environment := new(models.Environment)
	environment.Id = this.getInt("id", 0)
	if _, err := environment.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type ProjectController struct {
	BaseController
}

func (this *ProjectController) List() {
	this.pageTitle("项目列表")
	this.display(this.getBgTestAction("project/list"))
}

func (this *ProjectController) Add() {
	this.pageTitle("增加工程项目名称")
	this.parent(0)
	this.display(this.getBgTestAction("project/add"))
}

func (this *ProjectController) Edit() {
	this.pageTitle("编辑工程项目名称")
	project := new(models.Project)
	project.Id = this.getInt("id", 0)
	if err := project.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.parent(project.ParentId)
	row := make(map[string]interface{})
	row["parent_id"] = project.ParentId
	row["icon"] = project.Icon
	row["address"] = project.Address
	row["account"] = project.Account
	row["psw"] = project.Psw
	row["branch"] = project.Branch
	row["tag"] = project.Tag
	this.row(row, project.Id, project.Name, project.Descript)
	this.display(this.getBgTestAction("project/edit"))
}

func (this *ProjectController) AjaxSave() {
	project := new(models.Project)
	project.Id = this.getInt("id", 0)
	project.Icon = this.getString("file", "Icon不能为空!", defaultMinSize)
	project.Name = this.getString("name", "项目名称不能为空!", 1)
	project.Icon = this.getString("icon", "Icon不能为空!", 1)
	project.Address = this.getString("address", "项目地址不能为空!", 2)
	project.Account = this.getString("account", "", 0)
	project.Psw = this.getString("psw", "", 0)
	var err error
	if project.Id == 0 {
		project.CreateId = this.userId
		project.CreateTime = time.Now().Unix()
		_, err = project.Add()
	} else {
		project.UpdateId = this.userId
		project.UpdateTime = time.Now().Unix()
		_, err = project.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *ProjectController) Table() {
	project := new(models.Project)
	result, count := project.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["name"] = v.Name
		row["file"] = v.Icon
		row["address"] = v.Address
		row["account"] = v.Account
		row["psw"] = v.Psw
		row["branch"] = v.Branch
		row["tag"] = v.Tag
		row["descript"] = v.Descript
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *ProjectController) parent(id int) {
	environment := new(models.Environment)
	result, count := environment.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["id"] = v.Id
		row["name"] = v.Name
		if v.Id == id {
			row["selected"] = true
		} else {
			row["selected"] = false
		}
		list[k] = row
	}
	this.Data["group"] = list
}

func (this *ProjectController) AjaxDel() {
	project := new(models.Project)
	project.Id = this.getInt("id", 0)
	if _, err := project.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}
