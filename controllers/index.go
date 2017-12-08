package controllers

import (
	"zd112/models"
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	this.Data["img"] = "/static/mingzu/img/1.jpg"
	banner := new(models.Banner)
	result, _ := banner.List(this.pageSize, this.offSet)
	this.Data["row"] = result
	this.display("index")
}
