package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type BaseController struct {
	beego.Controller
	controller string
	action     string
	userId     int
	userName   string
	loginName  string
	pageSize   int
	allowUrl   string
}

func (this *BaseController) Prepare() {
	this.pageSize = 20
	controller, action := this.GetControllerAndAction()
	this.controller = strings.ToLower(controller[0:len(controller)-10])
	this.action = strings.ToLower(action)
	this.Data["version"] = beego.AppConfig.String("version")
	this.Data["siteName"] = beego.AppConfig.String("siteName")
	this.Data["route"] = this.controller + "." + this.action
	this.Data["action"] = this.action

	this.auth()

	this.Data["userId"] = this.userId
	this.Data["userName"] = this.userName
}

func (this *BaseController) auth() {

}

func (this *BaseController) AminAuth() {
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if this.userId != 1 {
		//adminAuthIds,_ := models.
	}
}
