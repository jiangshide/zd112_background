package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"zd112/models"
	"zd112/utils"
	"time"
	"strconv"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Login() {
	beego.Info("--------------login")
	if this.userId > 0 {
		this.redirect(beego.URLFor("HomeController.Index"))
	}
	beego.ReadFromRequest(&this.Controller)
	if this.isPost() {
		userName := strings.TrimSpace(this.GetString("username"))
		psw := strings.TrimSpace(this.GetString("psw"))
		if userName != "" && psw != "" {
			user, err := models.AdminGetByName(userName)
			flash := beego.NewFlash()
			errorMsg := ""
			if err != nil || user.Password != utils.Md5([]byte(psw+user.Salt)) {
				errorMsg = "账号或密码错误"
			} else if user.Status == -1 {
				errorMsg = "该账号已禁用"
			} else {
				user.LastIp = this.getClientIp()
				user.LastLogin = time.Now().Unix()
				user.Update()
				authKey := utils.Md5([]byte(this.getClientIp() + "|" + user.Password + user.Salt))
				this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authKey, 7*86400)
				this.redirect(beego.URLFor("HomeController.Index"))
			}
			flash.Error(errorMsg)
			flash.Store(&this.Controller)
			this.redirect(beego.URLFor("LoginController.Login"))
		}
	}
	this.TplName = "login/login.html"
}

func (this *LoginController) LoginOut() {
	this.Ctx.SetCookie("auth", "")
	this.redirect(beego.URLFor("LoginController.LoginIn"))
}

func (this *LoginController) NoAuth() {
	this.Ctx.WriteString("没有权限")
}
