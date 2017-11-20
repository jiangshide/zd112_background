package controllers

type HomeController struct{
	BaseController
}

func (this *HomeController) Index(){
	this.Data["pageTitle"]="系统首页"
	this.TplName="public/main.html"
}

func(this *HomeController)Start(){
	this.Data["pageTitle"]="控制面板"
	this.display()
}