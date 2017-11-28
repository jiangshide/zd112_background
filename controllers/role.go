package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"zd112/models"
)

type RoleController struct {
	BaseController
}

func (this *RoleController) List() {
	this.pageTitle("角色管理")
	this.display("backstage/role/list")
}

func (this *RoleController) Add() {
	this.Data["zTree"] = true
	this.pageTitle("增加角色")
	this.display("backstage/role/add")
}

func (this *RoleController) Edit() {
	this.Data["zTree"] = true
	this.pageTitle("编辑角色")
	id := this.getInt("id")
	role, _ := models.RoleGetById(id)
	this.Data["role"] = role
	roleAuth, _ := models.RoleAuthGetById(id)
	authId := make([]int, 0)
	for _, v := range roleAuth {
		authId = append(authId, v.AuthId)
	}
	this.Data["auth"] = authId
	this.display("backstage/role/edit")
}

func (this *RoleController) AjaxSave() {
	role := new(models.Role)
	role.Name = this.getString("role_name", "角色名称不能为空!", 0)
	role.Detail = this.getString("detail", "", -1)
	role.CreateTime = time.Now().Unix()
	role.Status = 1
}

func (this *RoleController) AjaxDel() {
	id := this.getInt("id")
	role, _ := models.RoleGetById(id)
	role.Status = 0
	role.Id = id
	role.UpdateTime = time.Now().Unix()
	if err := role.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *RoleController) Table() {
	beego.Info("--------------list")
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
	result, count := models.RoleList(page, this.pageSize, filters...)
	beego.Info("------------result:",result," | count:",count)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["role_name"] = v.Name
		row["detail"] = v.Detail
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		list[k] = row
	}
	this.ajaxList("成功", MSG_OK, count, list)
}
