package controllers

import (
	"zd112/models"
	"time"
	"zd112/utils"
	"github.com/astaxie/beego"
	"strings"
)

type BannerController struct {
	BaseController
}

func (this *BannerController) List() {
	this.pageTitle("焦点图列表")
	commRouter()
	this.display(this.getBgWebAction("banner/list"))
}

func commRouter() {
	actionStr := "list,add,edit,table,ajaxSave,ajaxDel"
	rout := make(map[string]interface{}, 0)
	rout["nation:"+actionStr] = &NationController{}
	rout["area/continent:"+actionStr] = &ContinentController{}
	rout["area/state:"+actionStr] = &StateController{}
	rout["area/province:"+actionStr] = &ProvinceController{}
	for k, v := range rout {
		kArr := strings.Split(k, ":")
		path := "/backstage/" + kArr[0]
		actions := strings.Split(kArr[1], ",")
		for _, action := range actions {
			rootPath := path + "/" + strings.ToLower(action)
			action = "*:" + utils.StrFirstToUpper(action)
			beego.Info("rootPath:", rootPath, " | action:", action, " | v:", v)
		}
	}

}

func (this *BannerController) Add() {
	this.pageTitle("增加焦点图名称")
	this.display(this.getBgWebAction("banner/add"))
}

func (this *BannerController) Table() {
	banner := new(models.Banner)
	result, count := banner.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["name"] = v.Name
		row["link"] = v.Link
		row["file"] = v.Icon
		row["descript"] = v.Descript
		row["clicks"] = v.Clicks
		row["views"] = v.Views
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
	row["link"] = banner.Link
	row["file"] = banner.Icon
	size, _ := utils.FileSize(banner.Icon)
	row["size"] = size
	row["descript"] = banner.Descript
	this.Data["row"] = row
	this.display(this.getBgWebAction("banner/edit"))
}

func (this *BannerController) AjaxSave() {
	banner := new(models.Banner)
	banner.Name = this.getString("name", "名称不能为空!", 1)
	banner.Link = this.getString("link", "", 0)
	banner.Icon = this.getString("file", "File不能为空!", defaultMinSize)
	banner.Descript = this.getString("descript", "", 0)
	banner.Id = this.getInt("id", 0)
	if banner.Id == 0 {

		banner.CreateId = this.userId
		banner.CreateTime = time.Now().Unix()
		if _, err := banner.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	banner.Query()
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
