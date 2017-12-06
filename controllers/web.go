package controllers

import (
	"zd112/models"
	"time"
)

type BannerController struct {
	BaseController
}

func (this *BannerController) List() {
	this.pageTitle("焦点图列表")
	this.display(this.getBgWebAction("banner/list"))
}

func (this *BannerController) Add() {
	this.pageTitle("增加焦点图名称")
	this.display(this.getBgWebAction("banner/add"))
}

func (this *BannerController) Table() {
	result, count := models.BannerList(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["name"] = v.Name
		row["url"] = v.Url
		row["icon"] = v.Icon
		row["descript"] = v.Descript
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *BannerController) Edit() {
	this.pageTitle("编辑焦点图名称")
	banner := new(models.Banner)
	banner.Id = this.getInt("id", 0)
	if err := banner.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{}, 0)
	row["id"] = banner.Id
	row["name"] = banner.Name
	row["url"] = banner.Url
	row["icon"] = banner.Icon
	row["descript"] = banner.Descript
	this.Data["row"] = row
	this.display(this.getBgWebAction("banner/edit"))
}

func (this *BannerController) AjaxSave() {
	banner := new(models.Banner)
	banner.Id = this.getInt("id", 0)
	if banner.Id == 0 {
		banner.Name = this.getString("name", "名称不能为空!", 1)
		banner.Url = this.getString("url", "", 0)
		banner.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
		banner.Descript = this.getString("descript", "", 0)
		banner.CreateId = this.userId
		banner.CreateTime = time.Now().Unix()
		if _, err := banner.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	banner.Query()
	banner.Name = this.getString("name", "名称不能为空!", 1)
	banner.Url = this.getString("url", "", 0)
	banner.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
	banner.Descript = this.getString("descript", "", 0)
	banner.UpdateId = this.userId
	banner.UpdateTime = time.Now().Unix()
	if _, err := banner.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *BannerController) AjaxDel() {
	banner := new(models.Banner)
	banner.Id = this.getInt("id", 0)
	if _, err := banner.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}