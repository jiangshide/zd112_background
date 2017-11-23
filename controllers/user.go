package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"strconv"
	"zd112/models"
	"zd112/utils"
	"fmt"
)

type UserController struct {
	BaseController
}

func (this *UserController) Admin() {
	beego.ReadFromRequest(&this.Controller)
	if this.isPost() {
		userName := this.getString("username", "账号不能为空!", defaultMinSize)
		password := this.getString("password", "密码不能为空!", defaultMinSize)
		realName := this.getString("realname", "真实名称不能为空!", 1)
		email := this.getString("email", "邮箱不能为空!", defaultMinSize)
		phone := this.getString("phone", "电话号码不能为空!", defaultMinSize)
		sex := this.getInt("radio")
		motto := this.getString("motto", "请君赠言...", defaultMinSize)
		admin := new(models.Admin)
		admin.Name = userName
		admin.RealName = realName
		admin.Salt = utils.GetRandomString(10)
		admin.LastIp = this.getClientIp()
		admin.Password = utils.Md5(password + admin.Salt)
		admin.Email = email
		admin.RoleIds = "1"
		admin.Phone = phone
		admin.Sex = sex
		admin.Motto = motto
		admin.CreateTime = time.Now().Unix()
		if _, err := models.AdminAdd(admin); err != nil {
			this.showTips(fmt.Sprint(err))
		}
		this.redirect(beego.URLFor("UserController.Login"))
	}
	this.TplName = "login/admin.html"
}

func (this *UserController) Reg() {
	beego.ReadFromRequest(&this.Controller)
	if this.isPost() {
		userName := this.getString("username", "账号不能为空!", defaultMinSize)
		password := this.getString("password", "密码不能为空!", defaultMinSize)
		email := this.getString("email", "邮箱不能为空!", defaultMinSize)
		sex := this.getInt("radio")

		admin := new(models.Admin)
		admin.Name = userName
		admin.Password = password
		admin.Email = email
		admin.Sex = sex
		admin.LastIp = this.getClientIp()
		admin.Salt = utils.GetRandomString(10)
		admin.Password = utils.Md5(password + admin.Salt)
		admin.RoleIds = "2"
		admin.CreateTime = time.Now().Unix()
		if _, err := models.AdminAdd(admin); err != nil {
			this.showTips(fmt.Sprint(err))
		}
		this.redirect(beego.URLFor("UserController.Login"))
	}
	this.TplName = "login/reg.html"
}

func (this *UserController) Login() {
	if this.userId > 0 {
		this.redirect(beego.URLFor("MainController.Index"))
	}
	beego.ReadFromRequest(&this.Controller)
	if this.isPost() {
		userName := this.getString("username", "账号不能为空!", defaultMinSize)
		password := this.getString("password", "密码不能为空!", defaultMinSize)
		user, err := models.AdminGetByName(userName)
		beego.Info("--------user:", user, " | err:", err)
		if err != nil || user.Password != utils.Md5(password+user.Salt) {
			this.showTips(err)
		} else if user.Status == -1 {
			this.showTips("该账号已禁用!")
		}
		user.LastIp = this.getClientIp()
		user.LastLogin = time.Now().Unix()
		user.Update()
		authKey := utils.Md5(this.getClientIp() + "|" + user.Password + user.Salt)
		this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authKey, 7*86400)
		this.redirect(beego.URLFor("MainController.Index"))
	}
	this.TplName = "login/login.html"
}

func (this *UserController) LoginOut() {
	beego.Info("----------loginout")
	this.Ctx.SetCookie("auth", "")
	this.redirect(beego.URLFor("IndexController.Index"))
}

func (this *UserController) NoAuth() {
	this.Ctx.WriteString("没有权限")
}

func(this *UserController)Edit(){
	beego.Info("---------------edit")
	this.Data["pageTitle"]="资料修改"
	id:=this.userId
	Admin,_:=models.AdminGetById(id)
	row := make(map[string]interface{})
	row["id"]=Admin.Id
	row["login_name"]=Admin.Name
	row["real_name"]=Admin.RealName
	row["phone"]=Admin.Phone
	row["email"]=Admin.Email
	this.Data["admin"]=row
	this.display("backstage/user/edit")
}