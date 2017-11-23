package controllers

import "github.com/astaxie/beego"

type BackstageController struct {
	BaseController
}

func (this *BackstageController) Index() {
	this.Data["pageTitle"] = "系统首页"
	beego.Info("-----------------backstage")
	this.display("backstage/index")
}

func (this *BackstageController) Start() {
	this.Data["pageTitle"] = "控制面板"
	this.display()
}
