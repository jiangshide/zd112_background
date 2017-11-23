package controllers

type BackstageController struct {
	BaseController
}

func (this *BackstageController) Index() {
	this.Data["pageTitle"] = "系统首页"
	this.TplName = "backstage/index.html"
}

func (this *BackstageController) Start() {
	this.Data["pageTitle"] = "控制面板"
	this.display("backstage/home/start")
}
