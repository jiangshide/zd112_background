package controllers

type NationController struct{
	BaseController
}

func (this *NationController) List(){
	this.pageTitle("民族列表")
	this.display("backstage/nation/list")
}