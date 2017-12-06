package controllers

import "zd112/models"

type BannerController struct{
	BaseController
}

func (this *BannerController) List(){
	this.pageTitle("焦点图列表")
	this.display(this.getBgWebAction("banner/list"))
}

func (this *BannerController) Add(){
	this.pageTitle("增加焦点图名称")
	this.display(this.getBgWebAction("banner/list"))
}

func (this *BannerController) Edit(){
	this.pageTitle("编辑焦点图名称")
	banner := new(models.Banner)
	banner.Id = this.getInt("id",0)
	if err := banner.Query();err != nil{
		this.ajaxMsg(err.Error(),MSG_ERR)
	}
	row := make(map[string]interface{},0)
	row["id"]=banner.Id
	row["name"]=banner.Name
	row["url"]=banner.Url
	row["icon"]=banner.Icon
	row["descript"]=banner.Descript
	this.Data["row"]=row
	this.display(this.getBgWebAction("banner/edit"))
}

func (this *BannerController) AjaxSave(){
	
}