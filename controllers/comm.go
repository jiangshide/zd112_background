package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"strconv"
	"zd112/utils"
	"fmt"
	"zd112/models"
	"time"
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
	userIcon   string
	isLogin    bool
	logo       string
	page       int
	pageSize   int
	offSet     int
	allowUrl   string
	user       *models.Admin
	defaultPsw string
	upload     string
}

var defaultTips = "该项不能为空!"
var defaultMinSize = 6

func (this *BaseController) pageTitle(pageTitle string) {
	this.Data["pageTitle"] = pageTitle
}

func (this *BaseController) showTips(errorMsg interface{}) {
	beego.Info("errorMsg:", errorMsg)
	flash := beego.NewFlash()
	flash.Error(fmt.Sprint(errorMsg))
	flash.Store(&this.Controller)
	controller, action := this.GetControllerAndAction()
	this.redirect(beego.URLFor(controller + "." + action))
}

func (this *BaseController) getString(key, tips string, minSize int) string {
	value := strings.TrimSpace(this.GetString(key, ""))
	errorMsg := ""
	beego.Info("----------key:", key, " | value:", value)
	if len(value) == 0 {
		errorMsg = tips
	} else if len(value) < minSize {
		errorMsg = "长度不能小于" + fmt.Sprint(minSize) + "字符"
	}
	if errorMsg != "" {
		this.showTips(errorMsg)
	}
	return value
}

func (this *BaseController) getInt(key string, defaultValue int) int {
	resInt, _ := this.GetInt(key, defaultValue)
	errorMsg := ""
	if errorMsg != "" {
		this.showTips(errorMsg)
	}
	return resInt
}

func (this *BaseController) Prepare() {
	this.pageSize = 20
	controller, action := this.GetControllerAndAction()
	this.controller = strings.ToLower(controller[0:len(controller)-10])
	this.action = strings.ToLower(action)
	this.userIcon = "/static/mingzu/img/3.jpg"
	this.upload = "/static/upload/"
	this.page = this.getInt("page", 1)
	this.pageSize = this.getInt("limit", 30)
	this.offSet = (this.page - 1) * this.pageSize

	this.Data["route"] = this.controller + "." + this.action
	this.Data["action"] = this.action
	this.defaultPsw = beego.AppConfig.String("defaultPsw")
	this.currParam()
}

func (this *BaseController) currParam() {
	this.auth()
	this.Data["userId"] = this.userId
	this.Data["userName"] = this.userName
	this.Data["userIcon"] = this.userIcon
	this.Data["isLogin"] = this.isLogin
	this.Data["logo"] = "/static/mingzu/img/11.jpg"
	this.Data["version"] = beego.AppConfig.String("version")
	this.Data["siteName"] = beego.AppConfig.String("site.name")
}

func (this *BaseController) auth() {
	arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
	this.userId = 0
	if len(arr) == 2 {
		idstr, psw := arr[0], arr[1]
		userId, _ := strconv.Atoi(idstr)
		if userId > 0 {
			user, err := models.AdminGetById(userId)
			if err == nil && psw == utils.Md5(this.getClientIp()+"|"+user.Password+user.Salt) {
				this.userId = user.Id
				this.loginName = user.Name
				this.userName = user.RealName
				this.userIcon = "/static/mingzu/img/1.jpg"
				this.isLogin = true
				this.user = user
			}
			if strings.Contains(this.controller, "backstage") {
				this.AdminAuth()
			}
			//isHashAuth := strings.Contains(this.allowUrl, this.controller+"/"+this.action)
			//noAuth := "ajaxsave/ajaxdel/table/loginin/getnodes/start/show/ajaxapisace"
			//isNoAuth := strings.Contains(noAuth, this.action)
			//if isHashAuth == false && isNoAuth == false {
			//	this.Ctx.WriteString("没有权限")
			//	this.ajaxMsg("没有权限", MSG_ERR)
			//	return
			//}
		}
	}
	//if this.userId == 0 && (this.controller != "login" && this.action != "loginin") {
	//	this.redirect(beego.URLFor("LoginController.LoginIn"))
	//}
}

func (this *BaseController) AdminAuth() {
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if this.userId != 1 {
		adminAuthIds, _ := models.RoleAuthGetByIds(this.user.RoleIds)
		adminAuthIdArr := strings.Split(adminAuthIds, ",")
		filters = append(filters, "id__in", adminAuthIdArr)
	}
	result, _ := models.AuthList(1, 100, filters...)
	list := make([]map[string]interface{}, len(result))
	lists := make([]map[string]interface{}, len(result))
	allow_url := ""
	i, j := 0, 0
	for _, v := range result {
		if v.Url != " " || v.Url != "/" {
			allow_url += v.Url
		}
		row := make(map[string]interface{})
		if v.Pid == 1 && v.IsShow == 1 {
			row["Id"] = int(v.Id)
			row["Sort"] = v.Sort
			row["Name"] = v.Name
			row["Url"] = v.Url
			row["Icon"] = v.Icon
			row["Pid"] = int(v.Pid)
			list[i] = row
			i++
		}
		if v.Pid != 1 && v.IsShow == 1 {
			row["Id"] = int(v.Id)
			row["Sort"] = v.Sort
			row["Name"] = v.Name
			row["Url"] = v.Url
			row["Icon"] = v.Icon
			row["Pid"] = int(v.Pid)
			lists[j] = row
			j++
		}
	}
	this.Data["SideMenu1"] = list[:i]
	this.Data["SideMenu2"] = lists[:j]
	this.allowUrl = allow_url + "/backstage/index"
}

func (this *BaseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

func (this *BaseController) getClientIp() string {
	return this.Ctx.Input.IP()
}

func (this *BaseController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

func (this *BaseController) display(tpl ...string) {
	this.currParam()
	var tplName string
	if len(tpl) > 0 {
		tplName = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplName = this.controller + "/" + this.action + ".html"
	}
	if strings.Contains(tplName, "backstage") {
		this.Layout = this.getBgAction("comm/layout.html")
	}
	this.TplName = tplName
}

func (this *BaseController) getBgAction(action string) string {
	return "backstage/" + action
}

func (this *BaseController) getBgAreaAction(action string) string {
	return this.getBgAction("area/" + action)
}

func (this *BaseController) getBgWebAction(action string) string {
	return this.getBgAction("web/" + action)
}

func (this *BaseController) ajaxMsg(msg interface{}, msgNo int) {
	out := make(map[string]interface{})
	out["status"] = msgNo
	out["message"] = msg
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) ajaxList(msg interface{}, msgNo int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgNo
	out["msg"] = msg
	out["count"] = count
	out["data"] = data
	this.Data["json"] = out
	this.Data["data"] = data
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) parse(list []map[string]interface{}, row map[string]interface{}, id, k, createId, updateId int, count, createTime, updateTime int64) {
	row["id"] = id
	if createId != 0 {
		if admin, err := models.AdminGetById(createId); err == nil {
			row["create_name"] = admin.Name
		} else {
			row["create_name"] = createId
		}
	}
	if updateId != 0 {
		if admin, err := models.AdminGetById(updateId); err == nil {
			row["update_name"] = admin.Name
		} else {
			row["update_name"] = updateId
		}
	}
	if createTime != 0 {
		row["create_time"] = beego.Date(time.Unix(createTime, 0), "Y-m-d H:i:s")
	}
	if updateTime != 0 {
		row["update_time"] = beego.Date(time.Unix(updateTime, 0), "Y-m-d H:i:s")
	}
	list[k] = row
}

func (this *BaseController) Upload() {
	f, fh, err := this.GetFile("file")
	defer f.Close()
	fileName := fh.Filename
	sufix := "default"
	if strings.Contains(fh.Filename, ".") {
		sufix = fileName[strings.LastIndex(fileName, ".")+1:]
	}
	fileName = utils.Md5(this.userName+time.RubyDate+utils.GetRandomString(10)) + "_" + fmt.Sprint(time.Now().Unix()) + "." + sufix
	err = this.SaveToFile("file", utils.GetCurrentDir(this.upload+sufix+"/")+fileName)
	beego.Error(err)
	this.ajaxMsg(this.upload+sufix+"/"+fileName, MSG_OK)
}
