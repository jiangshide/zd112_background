package controllers

import (
	"zd112/models"
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	this.Data["img"] = "/static/mingzu/img/1.jpg"
	result, _ := models.BannerList(this.pageSize, this.offSet)
	this.Data["row"] = result
	this.display("index")
}
