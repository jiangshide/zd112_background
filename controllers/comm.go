package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/george518/PPGo_ApiAdmin/models"
	"strconv"
	"zd112/utils"
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
	user *models.Admin
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
	arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
	this.userId = 0
	if len(arr) == 2 {
		idstr, psw := arr[0], arr[1]
		userId, _ := strconv.Atoi(idstr)
		if userId > 1 {
			user, err := models.AdminGetById(userId)
			if err == nil && psw == utils.Md5([]byte(this.getClientIp()+"|"+user.Password+user.Salt)) {
				this.userId = user.Id
				this.loginName = user.LoginName
				this.userName = user.RealName
				this.user = user
			}
			this.AdminAuth()
			isHashAuth := strings.Contains(this.allowUrl,this.controller+"/"+this.action)
			noAuth := "ajaxsave/ajaxdel/table/loginin/getnodes/start/show/ajaxapisace"
			isNoAuth := strings.Contains(noAuth,this.action)
			if isHashAuth == false && isNoAuth == false{
				this.Ctx.WriteString("没有权限")
				this.ajaxMsg("没有权限",MSG_ERR)
				return
			}
		}
	}
	if this.userId == 0 && (this.controller!="login" && this.action != "loginin"){
		this.redirect(beego.URLFor("LoginController.LoginIn"))
	}
}

func (this *BaseController) AdminAuth() {
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if this.userId != 1 {
		adminAuthIds, _ := models.RoleAuthGetByIds(this.user.RoleIds)
		adminAuthIdArr := strings.Split(adminAuthIds,",")
		filters = append(filters,"id__in",adminAuthIdArr)
	}
	result,_:=models.AuthGetList(1,100,filters...)
	list := make([]map[string]interface{},len(result))
	lists := make([]map[string]interface{},len(result))
	allow_url:=""
	i,j := 0,0
	for _,v := range result{
		if v.AuthUrl != " " || v.AuthUrl != "/"{
			allow_url+=v.AuthUrl
		}
		row :=make(map[string]interface{})
		if v.Pid == 1 && v.IsShow == 1{
			row["Id"]=int(v.Id)
			row["Sort"]=v.Sort
			row["AuthName"]=v.AuthName
			row["AuthUrl"]=v.AuthUrl
			row["Icon"]=v.Icon
			row["Pid"]=int(v.Pid)
			list[i]=row
			i++
		}
		if v.Pid != 1 && v.IsShow == 1{
			row["Id"]=int(v.Id)
			row["Sort"]=v.Sort
			row["AuthName"]=v.AuthName
			row["AuthUrl"]=v.AuthUrl
			row["Icon"]=v.Icon
			row["Pid"]=int(v.Pid)
			lists[j]=row
			j++
		}
	}
	this.Data["SideMenu1"]=list[:i]
	this.Data["SideMenu2"]=lists[:j]
}

func (this *BaseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

func (this *BaseController) getClientIp() string {
	return strings.Split(this.Ctx.Request.RemoteAddr, ":")[0]
}

func (this *BaseController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

func (this *BaseController) display(tpl ...string) {
	var tplName string
	if len(tpl) > 0 {
		tplName = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplName = this.controller + "/" + this.action + ".html"
	}
	this.Layout = "public/layout.html"
	this.TplName = tplName
}

func (this *BaseController) ajaxMsg(msg interface{}, msgNo int) {
	out := make(map[string]interface{})
	out["status"]=msgNo
	out["msg"]=msg
	this.Data["json"]=out
	this.ServeJSON()
	this.StopRun()
}

func(this *BaseController) ajaxList(msg interface{},msgNo int,count int64,data interface{}){
	out := make(map[string]interface{})
	out["code"]=msgNo
	out["msg"]=msg
	out["count"]=count
	out["data"]=data
	this.Data["json"]=out
	this.ServeJSON()
	this.StopRun()
}
