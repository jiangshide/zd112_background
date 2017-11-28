package controllers

import (
	"zd112/models"
	"time"
	"github.com/astaxie/beego"
)

type AuthController struct {
	BaseController
}

func (this *AuthController) Index() {
	this.pageTitle("权限因子")
	this.display()
}

func (this *AuthController) List() {
	this.pageTitle("权限因子列表")
	this.Data["zTree"]=true
	this.display("backstage/auth/list")
}

func (this *AuthController) GetNodes() {
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, count := models.AuthList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["pId"] = v.Pid
		row["name"] = v.Name
		row["open"] = true
		list[k] = row
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *AuthController) GetNode() {
	id := this.getInt("id")
	result, _ := models.AuthGetById(id)
	row := make(map[string]interface{})
	row["id"] = result.Id
	row["pid"] = result.Pid
	row["name"] = result.Name
	row["sort"] = result.Sort
	row["is_show"] = result.IsShow
	row["icon"] = result.Icon
	this.ajaxList("成功", MSG_OK, 0, row)
}

func (this *AuthController) AjaxSave() {
	auth := new(models.Auth)
	auth.UserId = this.userId
	auth.Pid = this.getInt("pid")
	auth.Name = this.getString("name", "名称不能为空!", 1)
	auth.Url = this.getString("url", "访问Action不能为空!", 1)
	auth.Sort = this.getInt("sort")
	auth.IsShow = this.getInt("is_show")
	auth.Icon = this.getString("icon", "菜单图标不能为空!", 1)
	auth.UpdateTime = time.Now().Unix()
	auth.Status = 1
	id := this.getInt("id")
	beego.Info("----------------update:~id:",id," | usetId:",this.userId)
	if id == 0 {
		auth.CreateTime = time.Now().Unix()
		auth.CreateId = this.userId
		auth.UpdateId = this.userId
		if _, err := models.AuthAdd(auth); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
	} else {
		auth.Id = id
		auth.UpdateId = this.userId
		if err := auth.Update(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *AuthController) AjaxDel() {
	id := this.getInt("id")
	//auth, _ := models.AuthGetById(id)
	//auth.Id = id
	//auth.Status = 0
	//if err := auth.Update(); err != nil {
	//	this.ajaxMsg(err.Error(), MSG_ERR)
	//}
	//this.ajaxMsg("", MSG_OK)
	if _,err := models.AuthDelById(id);err != nil{
		this.ajaxMsg(err.Error(),MSG_ERR)
	}
	this.ajaxMsg("",MSG_OK)
}
