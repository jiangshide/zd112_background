package controllers

import (
	"zd112/models"
	"github.com/astaxie/beego"
	"time"
	"zd112/utils"
	"strconv"
	"strings"
)

type AdminController struct {
	BaseController
}

func (this *AdminController) List() {
	this.pageTitle("管理员管理")
	this.display("backstage/admin/list")
}

func (this *AdminController) Add() {
	this.pageTitle("新增管理员")
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, _ := models.RoleList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["role_name"] = v.Name
		list[k] = row
	}
	this.Data["role"] = list
	this.display("backstage/admin/add")
}

func (this *AdminController) Edit() {
	this.pageTitle("编辑管理员")
	id := this.getInt("id", 0)
	admin, _ := models.AdminGetById(id)
	row := make(map[string]interface{})
	row["id"] = admin.Id
	row["login_name"] = admin.Name
	row["real_name"] = admin.RealName
	row["phone"] = admin.Phone
	row["email"] = admin.Email
	row["role_ids"] = admin.RoleIds
	this.Data["admin"] = row

	role_ids := strings.Split(admin.RoleIds, ",")

	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, _ := models.RoleList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["checked"] = 0
		for i := 0; i < len(role_ids); i++ {
			role_id, _ := strconv.Atoi(role_ids[i])
			if role_id == v.Id {
				row["checked"] = 1
			}
		}
		row["id"] = v.Id
		row["role_name"] = v.Name
		list[k] = row
	}
	this.Data["role"] = list
	this.display("backstage/admin/edit")
}

func (this AdminController) AjaxSave() {
	id := this.getInt("id", 0)
	if id == 0 {
		admin := new(models.Admin)
		this.getParam(admin)
		admin.Salt = utils.GetRandomString(10)
		admin.Password = utils.Md5(this.defaultPsw + admin.Salt)
		admin.CreateTime = time.Now().Unix()
		admin.CreateId = this.userId
		if _, err := models.AdminAdd(admin); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	admin, _ := models.AdminGetById(id)
	admin.Id = id
	admin.UpdateTime = time.Now().Unix()
	admin.UpdateId = this.userId
	this.getParam(admin)

	resetPsw := this.getInt("reset_pwd", 0)
	if resetPsw == 1 {
		admin.Password = utils.Md5(this.defaultPsw + admin.Salt)
		admin.Salt = utils.GetRandomString(10)
	}
	if err := admin.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg(strconv.Itoa(resetPsw), MSG_OK)
}

func (this AdminController) getParam(admin *models.Admin) {
	admin.Name = this.getString("login_name", "账号不能为空!", defaultMinSize)
	admin.RealName = this.getString("real_name", "真实名字不能为空!", defaultMinSize)
	admin.Phone = this.getString("phone", "联系电话不能为空!", defaultMinSize)
	admin.Email = this.getString("email", "邮箱不能为空!", defaultMinSize)
	admin.RoleIds = this.getString("roleids", "没有选择权限项!", 0)
	admin.UpdateTime = time.Now().Unix()
	admin.UpdateId = this.userId
	admin.Status = 1
}

func (this AdminController) AjaxDel() {
	id := this.getInt("id", 0)
	beego.Info("---------------id:", id)
	admin, _ := models.AdminGetById(id)
	admin.UpdateTime = time.Now().Unix()
	admin.Status = 0
	admin.Id = id
	if id == 1 {
		this.ajaxMsg("超级管理员不允许删除!", MSG_ERR)
	}
	if err := admin.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *AdminController) Table() {
	page, err := this.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := this.GetInt("limit")
	if err != nil {
		limit = 30
	}
	this.pageSize = limit
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, count := models.AdminList(page, this.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["login_name"] = v.Name
		row["real_name"] = v.RealName
		row["phone"] = v.Phone
		row["email"] = v.Email
		row["role_ids"] = v.RoleIds
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		list[k] = row
	}
	this.ajaxList("成功", MSG_OK, count, list)
}
