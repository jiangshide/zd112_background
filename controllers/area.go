package controllers

import (
	"zd112/models"
	"time"
	"github.com/astaxie/beego"
)

type ContinentController struct {
	BaseController
}

func (this *ContinentController) List() {
	this.pageTitle("洲列表")
	this.display(this.getBgAreaAction("continent/list"))
}

func (this *ContinentController) Add() {
	this.pageTitle("新增洲名称")
	this.display(this.getBgAreaAction("continent/add"))
}

func (this *ContinentController) Edit() {
	this.pageTitle("编辑洲名称")
	continent := new(models.Continent)
	continent.Id = this.getInt("id", 0)
	if err := continent.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{}, 0)
	row["id"] = continent.Id
	row["name"] = continent.Name
	row["icon"] = continent.Icon
	this.Data["row"] = row
	this.display(this.getBgAreaAction("continent/edit"))
}

func (this *ContinentController) AjaxSave() {
	continent := new(models.Continent)
	continent.Id = this.getInt("id", 0)
	if continent.Id == 0 {
		continent.Name = this.getString("name", "名称不能为空!", 1)
		continent.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
		continent.CreateId = this.userId
		continent.CreateTime = time.Now().Unix()
		if _, err := continent.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	continent.Query()
	continent.Name = this.getString("name", "名称不能为空!", 1)
	continent.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
	continent.UpdateId = this.userId
	continent.UpdateTime = time.Now().Unix()
	if _, err := continent.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *ContinentController) Table() {
	result, count := models.ContinentList(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["name"] = v.Name
		row["icon"] = v.Icon
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *ContinentController) AjaxDel() {
	continent := new(models.Continent)
	continent.Id = this.getInt("id", 0)
	if _, err := continent.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type StateController struct {
	BaseController
}

func (this *StateController) List() {
	this.pageTitle("国列表")
	this.display(this.getBgAreaAction("state/list"))
}

func (this *StateController) Add() {
	this.pageTitle("增加国名称")
	this.parent()
	this.display(this.getBgAreaAction("state/add"))
}

func (this *StateController) Edit() {
	this.pageTitle("编辑国名")
	state := new(models.State)
	state.Id = this.getInt("id", 0)
	if err := state.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.parent()
	row := make(map[string]interface{}, 0)
	row["id"] = state.Id
	row["name"] = state.Name
	row["icon"] = state.Icon
	this.Data["row"] = row
	this.display(this.getBgAreaAction("state/edit"))
}

func (this *StateController) AjaxSave() {
	state := new(models.State)
	state.Id = this.getInt("id", 0)
	if state.Id == 0 {
		state.ParentId = this.getInt("group_id",0)
		state.Name = this.getString("name", "名称不能为空!", 1)
		state.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
		state.CreateId = this.userId
		state.CreateTime = time.Now().Unix()
		if _, err := state.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	state.Query()
	state.ParentId = this.getInt("group_id",0)
	state.Name = this.getString("name", "名称不能为空!", 1)
	state.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
	state.UpdateId = this.userId
	state.UpdateTime = time.Now().Unix()
	if _, err := state.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *StateController) parent() {
	result, count := models.ContinentList(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["id"] = v.Id
		row["name"] = v.Name
		list[k] = row
	}
	this.Data["group"] = list
}

func (this *StateController) Table() {
	result, count := models.StateList(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["name"] = v.Name
		row["icon"] = v.Icon
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *StateController) AjaxDel() {
	state := new(models.State)
	state.Id = this.getInt("id", 0)
	if _, err := state.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type ProvinceController struct {
	BaseController
}

func (this *ProvinceController) List() {
	this.pageTitle("省列表")
	this.display(this.getBgAreaAction("province/list"))
}

func (this *ProvinceController) Add() {
	this.pageTitle("增加省名称")
	this.display(this.getBgAreaAction("province/add"))
}

func (this *ProvinceController) Edit() {
	this.pageTitle("编辑省名称")
	province := new(models.Province)
	province.Id = this.getInt("id", 0)
	if err := province.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{}, 0)
	row["id"] = province.Id
	row["name"] = province.Name
	row["icon"] = province.Icon
	this.Data["row"] = row
	this.display(this.getBgAreaAction("province/edit"))
}

func (this *ProvinceController) AjaxSave() {
	province := new(models.Province)
	province.Id = this.getInt("id", 0)
	if province.Id == 0 {
		province.Name = this.getString("name", "名称不能为空!", 1)
		province.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
		province.CreateId = this.userId
		province.CreateTime = time.Now().Unix()
		if _, err := province.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	province.Query()
	province.Name = this.getString("name", "名称不能为空!", 1)
	province.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
	province.UpdateId = this.userId
	province.UpdateTime = time.Now().Unix()
	if _, err := province.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *ProvinceController) Table() {
	result, count := models.ProvinceList(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["name"] = v.Name
		row["icon"] = v.Icon
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *ProvinceController) AjaxDel() {
	province := new(models.Province)
	province.Id = this.getInt("id", 0)
	if _, err := province.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type CityController struct {
	BaseController
}

func (this *CityController) List() {
	this.pageTitle("城市列表")
	this.display(this.getBgAreaAction("city/list"))
}

func (this *CityController) Add() {
	this.pageTitle("增加城市名称")
	this.display(this.getBgAreaAction("city/add"))
}

func (this *CityController) Edit() {
	this.pageTitle("编辑城市名称")
	city := new(models.City)
	city.Id = this.getInt("id", 0)
	if err := city.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{}, 0)
	row["id"] = city.Id
	row["name"] = city.Name
	row["icon"] = city.Icon
	this.Data["row"] = row
	this.display(this.getBgAreaAction("city/edit"))
}

func (this *CityController) AjaxSave() {
	city := new(models.City)
	city.Id = this.getInt("id", 0)
	if city.Id == 0 {
		city.Name = this.getString("name", "名称不能为空!", 1)
		city.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
		city.CreateId = this.userId
		city.CreateTime = time.Now().Unix()
		if _, err := city.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	city.Query()
	city.Name = this.getString("name", "名称不能为空!", 1)
	city.Icon = this.getString("icon", "Icon不能为空！", defaultMinSize)
	city.UpdateId = this.userId
	city.UpdateTime = time.Now().Unix()
	if _, err := city.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *CityController) Table() {
	result, count := models.CityList(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["name"] = v.Name
		row["icon"] = v.Icon
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *CityController) AjaxDel() {
	city := new(models.City)
	city.Id = this.getInt("id", 0)
	if _, err := city.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type RegionController struct {
	BaseController
}

func (this *RegionController) List() {
	this.pageTitle("地区列表")
	this.display(this.getBgAreaAction("region/list"))
}

func (this *RegionController) Add() {
	this.pageTitle("增加地区名称")
	this.display(this.getBgAreaAction("region/add"))
}

func (this *RegionController) Edit() {
	this.pageTitle("编辑地区名称")
	region := new(models.Region)
	region.Id = this.getInt("id", 0)
	if err := region.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{}, 0)
	row["id"] = region.Id
	row["name"] = region.Name
	row["icon"] = region.Icon
	this.Data["row"] = row
	this.display(this.getBgAreaAction("region/edit"))
}

func (this *RegionController) AjaxSave() {
	region := new(models.Region)
	region.Id = this.getInt("id", 0)
	if region.Id == 0 {
		region.Name = this.getString("name", "名称不能为空!", 1)
		region.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
		region.CreateId = this.userId
		region.CreateTime = time.Now().Unix()
		if _, err := region.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	region.Query()
	region.Name = this.getString("name", "名称不能为空!", 1)
	region.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
	region.UpdateId = this.userId
	region.UpdateTime = time.Now().Unix()
	if _, err := region.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *RegionController) Table() {
	result, count := models.RegionList(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["name"] = v.Name
		row["icon"] = v.Icon
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *RegionController) AjaxDel() {
	region := new(models.Region)
	region.Id = this.getInt("id", 0)
	if _, err := region.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type CountyController struct {
	BaseController
}

func (this *CountyController) List() {
	this.pageTitle("县列表")
	this.display(this.getBgAreaAction("county/list"))
}

func (this *CountyController) Add() {
	this.pageTitle("增加县名称")
	this.display(this.getBgAreaAction("county/add"))
}

func (this *CountyController) Edit() {
	county := new(models.County)
	county.Id = this.getInt("id", 0)
	if err := county.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{}, 0)
	row["id"] = county.Id
	row["name"] = county.Name
	row["icon"] = county.Icon
	this.Data["row"] = row
	this.display(this.getBgAreaAction("county/edit"))
}

func (this *CountyController) AjaxSave() {
	county := new(models.County)
	county.Id = this.getInt("id", 0)
	if county.Id == 0 {
		county.Name = this.getString("name", "名称不能为空!", 1)
		county.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
		county.CreateId = this.userId
		county.CreateTime = time.Now().Unix()
		if _, err := county.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	county.Query()
	county.Name = this.getString("name", "名称不能为空!", 1)
	county.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
	county.UpdateId = this.userId
	county.UpdateTime = time.Now().Unix()
	if _, err := county.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *CountyController) Table() {
	result, count := models.CountyList(this.pageSize, this.offSet)
	beego.Info("----------------result:", result, " | count:", count)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["name"] = v.Name
		row["icon"] = v.Icon
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *CountyController) AjaxDel() {
	county := new(models.County)
	county.Id = this.getInt("id", 0)
	if _, err := county.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type TownController struct {
	BaseController
}

func (this *TownController) List() {
	this.pageTitle("镇列表")
	this.display(this.getBgAreaAction("town/list"))
}

func (this *TownController) Add() {
	this.pageTitle("增加镇名称")
	this.display(this.getBgAreaAction("town/add"))
}

func (this *TownController) Edit() {
	this.pageTitle("编辑镇名称")
	town := new(models.Town)
	town.Id = this.getInt("id", 0)
	if err := town.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{}, 0)
	row["id"] = town.Id
	row["name"] = town.Name
	row["icon"] = town.Icon
	this.Data["row"] = row
	this.display(this.getBgAreaAction("town/edit"))
}

func (this *TownController) AjaxSave() {
	town := new(models.Town)
	town.Id = this.getInt("id", 0)
	if town.Id == 0 {
		town.Name = this.getString("name", "名称不能为空!", 1)
		town.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
		town.CreateId = this.userId
		town.CreateTime = time.Now().Unix()
		if _, err := town.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	town.Query()
	town.Name = this.getString("name", "名称不能为空!", 1)
	town.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
	town.UpdateId = this.userId
	town.UpdateTime = time.Now().Unix()
	if _, err := town.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *TownController) Table() {
	result, count := models.TownList(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["name"] = v.Name
		row["icon"] = v.Icon
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *TownController) AjaxDel() {
	town := new(models.Town)
	town.Id = this.getInt("id", 0)
	if _, err := town.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type CountryController struct {
	BaseController
}

func (this *CountryController) List() {
	this.pageTitle("乡列表")
	this.display(this.getBgAreaAction("country/list"))
}

func (this *CountryController) Add() {
	this.pageTitle("增加乡名称")
	this.display(this.getBgAreaAction("country/add"))
}

func (this *CountryController) Edit() {
	this.pageTitle("编辑乡名称")
	country := new(models.Country)
	country.Id = this.getInt("id", 0)
	if err := country.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{}, 0)
	row["id"] = country.Id
	row["name"] = country.Name
	row["icon"] = country.Icon
	this.Data["row"] = row
	this.display(this.getBgAreaAction("country/edit"))
}

func (this *CountryController) AjaxSave() {
	country := new(models.Country)
	country.Id = this.getInt("id", 0)
	if country.Id == 0 {
		country.Name = this.getString("name", "名称不能为空!", 1)
		country.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
		country.CreateId = this.userId
		country.CreateTime = time.Now().Unix()
		if _, err := country.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	country.Query()
	country.Name = this.getString("name", "名称不能为空!", 1)
	country.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
	country.UpdateId = this.userId
	country.UpdateTime = time.Now().Unix()
	if _, err := country.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *CountryController) Table() {
	result, count := models.CountryList(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["name"] = v.Name
		row["icon"] = v.Icon
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *CountryController) AjaxDel() {
	country := new(models.Country)
	country.Id = this.getInt("id", 0)
	if _, err := country.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type VillageController struct {
	BaseController
}

func (this *VillageController) List() {
	this.pageTitle("村列表")
	this.display(this.getBgAreaAction("village/list"))
}

func (this *VillageController) Add() {
	this.pageTitle("增加村名称")
	this.display(this.getBgAreaAction("village/add"))
}

func (this *VillageController) Edit() {
	this.pageTitle("编辑村名称")
	village := new(models.Village)
	village.Id = this.getInt("id", 0)
	if err := village.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{}, 0)
	row["id"] = village.Id
	row["name"] = village.Name
	row["icon"] = village.Icon
	this.Data["row"] = row
	this.display(this.getBgAreaAction("village/edit"))
}

func (this *VillageController) AjaxSave() {
	village := new(models.Village)
	village.Id = this.getInt("id", 0)
	if village.Id == 0 {
		village.Name = this.getString("name", "名称不能为空!", 1)
		village.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
		village.CreateId = this.userId
		village.CreateTime = time.Now().Unix()
		if _, err := village.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	village.Query()
	village.Name = this.getString("name", "名称不能为空!", 1)
	village.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
	village.UpdateId = this.userId
	village.UpdateTime = time.Now().Unix()
	if _, err := village.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *VillageController) Table() {
	result, count := models.VillageList(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["name"] = v.Name
		row["icon"] = v.Icon
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *VillageController) AjaxDel() {
	village := new(models.Village)
	village.Id = this.getInt("id", 0)
	if _, err := village.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type GroupController struct {
	BaseController
}

func (this *GroupController) List() {
	this.pageTitle("组列表")
	this.display(this.getBgAreaAction("group/list"))
}

func (this *GroupController) Add() {
	this.pageTitle("增加组名称")
	this.display(this.getBgAreaAction("group/add"))
}

func (this *GroupController) Edit() {
	this.pageTitle("编辑组名称")
	group := new(models.Group)
	group.Id = this.getInt("id", 0)
	if err := group.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{}, 0)
	row["id"] = group.Id
	row["name"] = group.Name
	row["icon"] = group.Icon
	this.Data["row"] = row
	this.display(this.getBgAreaAction("group/edit"))
}

func (this *GroupController) AjaxSave() {
	group := new(models.Group)
	group.Id = this.getInt("id", 0)
	if group.Id == 0 {
		group.Name = this.getString("name", "名称不能为空!", 1)
		group.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
		group.CreateId = this.userId
		group.CreateTime = time.Now().Unix()
		if _, err := group.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	group.Query()
	group.Name = this.getString("name", "名称不能为空!", 1)
	group.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
	group.UpdateId = this.userId
	group.UpdateTime = time.Now().Unix()
	if _, err := group.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *GroupController) Table() {
	result, count := models.GroupList(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["name"] = v.Name
		row["icon"] = v.Icon
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *GroupController) AjaxDel() {
	group := new(models.Group)
	group.Id = this.getInt("id", 0)
	if _, err := group.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type TeamController struct {
	BaseController
}

func (this *TeamController) List() {
	this.pageTitle("队列表")
	this.display(this.getBgAreaAction("team/list"))
}

func (this *TeamController) Add() {
	this.pageTitle("增加队名称")
	this.display(this.getBgAreaAction("team/add"))
}

func (this *TeamController) Edit() {
	this.pageTitle("编辑队名称")
	team := new(models.Team)
	team.Id = this.getInt("id", 0)
	if err := team.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{}, 0)
	row["id"] = team.Id
	row["name"] = team.Name
	row["icon"] = team.Icon
	this.Data["row"] = row
	this.display(this.getBgAreaAction("team/edit"))
}

func (this *TeamController) AjaxSave() {
	team := new(models.Team)
	team.Id = this.getInt("id", 0)
	if team.Id == 0 {
		team.Name = this.getString("name", "名称不能为空!", 1)
		team.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
		team.CreateId = this.userId
		team.CreateTime = time.Now().Unix()
		if _, err := team.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	team.Query()
	team.Name = this.getString("name", "名称不能为空!", 1)
	team.Icon = this.getString("icon", "Icon不能为空!", defaultMinSize)
	team.UpdateId = this.userId
	team.UpdateTime = time.Now().Unix()
	if _, err := team.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *TeamController) Table() {
	result, count := models.TeamList(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["name"] = v.Name
		row["icon"] = v.Icon
		this.parse(list, row, v.Id, k, v.CreateId, v.UpdateId, count, v.CreateTime, v.UpdateTime)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *TeamController) AjaxDel() {
	team := new(models.Team)
	team.Id = this.getInt("id", 0)
	if _, err := team.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}
