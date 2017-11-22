package controllers

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	this.currParam()
	this.display("index")
}
