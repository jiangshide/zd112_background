package controllers

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	this.currParam()
	this.Data["img"] = "/static/mingzu/img/1.jpg"
	this.display("index")
}
