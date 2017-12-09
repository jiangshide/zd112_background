package controllers

import (
	"zd112/models"
	"time"
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
	row["file"] = continent.Icon
	this.setFileSize(row, continent.Icon)
	this.Data["row"] = row
	this.display(this.getBgAreaAction("continent/edit"))
}

func (this *ContinentController) AjaxSave() {
	continent := new(models.Continent)
	continent.Name = this.getString("name", "名称不能为空!", 1)
	continent.Icon = this.getString("file", "Icon不能为空!", defaultMinSize)
	continent.Id = this.getInt("id", 0)
	if continent.Id == 0 {
		continent.CreateId = this.userId
		continent.CreateTime = time.Now().Unix()
		if _, err := continent.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	continent.UpdateId = this.userId
	continent.UpdateTime = time.Now().Unix()
	if _, err := continent.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *ContinentController) Table() {
	continent := new(models.Continent)
	result, count := continent.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["name"] = v.Name
		row["file"] = v.Icon
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
	this.parent(0)
	this.display(this.getBgAreaAction("state/add"))
}

func (this *StateController) Edit() {
	this.pageTitle("编辑国名")
	state := new(models.State)
	state.Id = this.getInt("id", 0)
	if err := state.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{}, 0)
	row["id"] = state.Id
	this.parent(state.ParentId)
	row["name"] = state.Name
	row["file"] = state.Icon
	this.setFileSize(row, state.Icon)
	this.Data["row"] = row
	this.display(this.getBgAreaAction("state/edit"))
}

func (this *StateController) AjaxSave() {
	state := new(models.State)
	state.Name = this.getString("name", "名称不能为空!", 1)
	state.Icon = this.getString("file", "Icon不能为空!", defaultMinSize)
	state.Id = this.getInt("id", 0)
	if state.Id == 0 {
		state.ParentId = this.getInt("group_id", 0)
		state.CreateId = this.userId
		state.CreateTime = time.Now().Unix()
		if _, err := state.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	state.ParentId = this.getInt("group_id", 0)
	state.UpdateId = this.userId
	state.UpdateTime = time.Now().Unix()
	if _, err := state.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *StateController) parent(id int) {
	continent := new(models.Continent)
	result, count := continent.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["id"] = v.Id
		row["name"] = v.Name
		if v.Id == id {
			row["selected"] = true
		} else {
			row["selected"] = false
		}
		list[k] = row
	}
	this.Data["group"] = list
}

func (this *StateController) Table() {
	state := new(models.State)
	result, count := state.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["name"] = v.Name
		row["file"] = v.Icon
		continent := new(models.Continent)
		continent.Id = v.ParentId
		if err := continent.Query(); err == nil {
			row["parent"] = continent.Name
		}
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
	this.parent(0)
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
	this.parent(province.ParentId)
	row["id"] = province.Id
	row["name"] = province.Name
	row["file"] = province.Icon
	this.setFileSize(row, province.Icon)
	this.Data["row"] = row
	this.display(this.getBgAreaAction("province/edit"))
}

func (this *ProvinceController) AjaxSave() {
	province := new(models.Province)
	province.Name = this.getString("name", "名称不能为空!", 1)
	province.Icon = this.getString("file", "File不能为空!", defaultMinSize)
	province.Id = this.getInt("id", 0)
	if province.Id == 0 {
		province.ParentId = this.getInt("group_id", 0)
		province.CreateId = this.userId
		province.CreateTime = time.Now().Unix()
		if _, err := province.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	province.ParentId = this.getInt("group_id", 0)
	province.UpdateId = this.userId
	province.UpdateTime = time.Now().Unix()
	if _, err := province.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *ProvinceController) parent(id int) {
	state := new(models.State)
	result, count := state.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["id"] = v.Id
		row["name"] = v.Name
		if v.Id == id {
			row["selected"] = true
		} else {
			row["selected"] = false
		}
		list[k] = row
	}
	this.Data["group"] = list
}

func (this *ProvinceController) Table() {
	province := new(models.Province)
	result, count := province.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["name"] = v.Name
		row["file"] = v.Icon
		state := new(models.State)
		state.Id = v.ParentId
		if err := state.Query(); err == nil {
			row["parent"] = state.Name
		}
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
	this.parent(0)
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
	this.parent(city.ParentId)
	row["id"] = city.Id
	row["name"] = city.Name
	row["file"] = city.Icon
	this.setFileSize(row, city.Icon)
	this.Data["row"] = row
	this.display(this.getBgAreaAction("city/edit"))
}

func (this *CityController) AjaxSave() {
	city := new(models.City)
	city.Name = this.getString("name", "名称不能为空!", 1)
	city.Icon = this.getString("file", "File不能为空!", defaultMinSize)
	city.Id = this.getInt("id", 0)
	if city.Id == 0 {
		city.ParentId = this.getInt("group_id", 0)
		city.CreateId = this.userId
		city.CreateTime = time.Now().Unix()
		if _, err := city.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	city.ParentId = this.getInt("group_id", 0)
	city.UpdateId = this.userId
	city.UpdateTime = time.Now().Unix()
	if _, err := city.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *CityController) parent(id int) {
	province := new(models.Province)
	result, count := province.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["id"] = v.Id
		row["name"] = v.Name
		if v.Id == id {
			row["selected"] = true
		} else {
			row["selected"] = false
		}
		list[k] = row
	}
	this.Data["group"] = list
}

func (this *CityController) Table() {
	city := new(models.City)
	result, count := city.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["name"] = v.Name
		row["file"] = v.Icon
		province := new(models.Province)
		province.Id = v.ParentId
		if err := province.Query(); err == nil {
			row["parent"] = province.Name
		}
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
	this.parent(0)
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
	this.parent(region.ParentId)
	row["id"] = region.Id
	row["name"] = region.Name
	row["file"] = region.Icon
	this.setFileSize(row, region.Icon)
	this.Data["row"] = row
	this.display(this.getBgAreaAction("region/edit"))
}

func (this *RegionController) AjaxSave() {
	region := new(models.Region)
	region.Name = this.getString("name", "名称不能为空!", 1)
	region.Icon = this.getString("file", "File不能为空!", defaultMinSize)
	region.Id = this.getInt("id", 0)
	if region.Id == 0 {
		region.ParentId = this.getInt("group_id", 0)
		region.CreateId = this.userId
		region.CreateTime = time.Now().Unix()
		if _, err := region.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	region.ParentId = this.getInt("group_id", 0)
	region.UpdateId = this.userId
	region.UpdateTime = time.Now().Unix()
	if _, err := region.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *RegionController) parent(id int) {
	city := new(models.City)
	result, count := city.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["id"] = v.Id
		row["name"] = v.Name
		if v.Id == id {
			row["selected"] = true
		} else {
			row["selected"] = false
		}
		list[k] = row
	}
	this.Data["group"] = list
}

func (this *RegionController) Table() {
	region := new(models.Region)
	result, count := region.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["name"] = v.Name
		row["file"] = v.Icon
		city := new(models.City)
		city.Id = v.ParentId
		if err := city.Query(); err == nil {
			row["parent"] = city.Name
		}
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
	this.parent(0)
	this.display(this.getBgAreaAction("county/add"))
}

func (this *CountyController) Edit() {
	county := new(models.County)
	county.Id = this.getInt("id", 0)
	if err := county.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	row := make(map[string]interface{}, 0)
	this.parent(county.ParentId)
	row["id"] = county.Id
	row["name"] = county.Name
	row["file"] = county.Icon
	this.setFileSize(row, county.Icon)
	this.Data["row"] = row
	this.display(this.getBgAreaAction("county/edit"))
}

func (this *CountyController) AjaxSave() {
	county := new(models.County)
	county.Name = this.getString("name", "名称不能为空!", 1)
	county.Icon = this.getString("file", "File不能为空!", defaultMinSize)
	county.Id = this.getInt("id", 0)
	if county.Id == 0 {
		county.ParentId = this.getInt("group_id", 0)
		county.CreateId = this.userId
		county.CreateTime = time.Now().Unix()
		if _, err := county.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	county.ParentId = this.getInt("group_id", 0)
	county.UpdateId = this.userId
	county.UpdateTime = time.Now().Unix()
	if _, err := county.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *CountyController) parent(id int) {
	city := new(models.City)
	result, count := city.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["id"] = v.Id
		row["name"] = v.Name
		if v.Id == id {
			row["selected"] = true
		} else {
			row["selected"] = false
		}
		list[k] = row
	}
	this.Data["group"] = list
}

func (this *CountyController) Table() {
	county := new(models.County)
	result, count := county.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["name"] = v.Name
		row["file"] = v.Icon
		region := new(models.Region)
		region.Id = v.ParentId
		if err := region.Query(); err == nil {
			row["parent"] = region.Name
		}
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
	this.parent(0)
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
	this.parent(town.ParentId)
	row["id"] = town.Id
	row["name"] = town.Name
	row["file"] = town.Icon
	this.setFileSize(row, town.Icon)
	this.Data["row"] = row
	this.display(this.getBgAreaAction("town/edit"))
}

func (this *TownController) AjaxSave() {
	town := new(models.Town)
	town.Name = this.getString("name", "名称不能为空!", 1)
	town.Icon = this.getString("file", "File不能为空!", defaultMinSize)
	town.Id = this.getInt("id", 0)
	if town.Id == 0 {
		town.ParentId = this.getInt("group_id", 0)
		town.CreateId = this.userId
		town.CreateTime = time.Now().Unix()
		if _, err := town.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	town.ParentId = this.getInt("group_id", 0)
	town.UpdateId = this.userId
	town.UpdateTime = time.Now().Unix()
	if _, err := town.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *TownController) parent(id int) {
	county := new(models.County)
	result, count := county.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["id"] = v.Id
		row["name"] = v.Name
		if v.Id == id {
			row["selected"] = true
		} else {
			row["selected"] = false
		}
		list[k] = row
	}
	this.Data["group"] = list
}

func (this *TownController) Table() {
	town := new(models.Town)
	result, count := town.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["name"] = v.Name
		row["file"] = v.Icon
		county := new(models.County)
		county.Id = v.ParentId
		if err := county.Query(); err == nil {
			row["parent"] = county.Name
		}
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
	this.parent(0)
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
	this.parent(country.ParentId)
	row["id"] = country.Id
	row["name"] = country.Name
	row["file"] = country.Icon
	this.setFileSize(row, country.Icon)
	this.Data["row"] = row
	this.display(this.getBgAreaAction("country/edit"))
}

func (this *CountryController) AjaxSave() {
	country := new(models.Country)
	country.Name = this.getString("name", "名称不能为空!", 1)
	country.Icon = this.getString("file", "File不能为空!", defaultMinSize)
	country.Id = this.getInt("id", 0)
	if country.Id == 0 {
		country.ParentId = this.getInt("group_id", 0)
		country.CreateId = this.userId
		country.CreateTime = time.Now().Unix()
		if _, err := country.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	country.ParentId = this.getInt("group_id", 0)
	country.UpdateId = this.userId
	country.UpdateTime = time.Now().Unix()
	if _, err := country.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *CountryController) parent(id int) {
	town := new(models.Town)
	result, count := town.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["id"] = v.Id
		row["name"] = v.Name
		if v.Id == id {
			row["selected"] = true
		} else {
			row["selected"] = false
		}
		list[k] = row
	}
	this.Data["group"] = list
}

func (this *CountryController) Table() {
	country := new(models.Country)
	result, count := country.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["name"] = v.Name
		row["file"] = v.Icon
		town := new(models.Town)
		town.Id = v.ParentId
		if err := town.Query(); err == nil {
			row["parent"] = town.Name
		}
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
	this.parent(0)
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
	this.parent(village.ParentId)
	row["id"] = village.Id
	row["name"] = village.Name
	row["file"] = village.Icon
	this.setFileSize(row, village.Icon)
	this.Data["row"] = row
	this.display(this.getBgAreaAction("village/edit"))
}

func (this *VillageController) AjaxSave() {
	village := new(models.Village)
	village.Name = this.getString("name", "名称不能为空!", 1)
	village.Icon = this.getString("file", "File不能为空!", defaultMinSize)
	village.Id = this.getInt("id", 0)
	if village.Id == 0 {
		village.ParentId = this.getInt("group_id", 0)
		village.CreateId = this.userId
		village.CreateTime = time.Now().Unix()
		if _, err := village.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	village.ParentId = this.getInt("group_id", 0)
	village.UpdateId = this.userId
	village.UpdateTime = time.Now().Unix()
	if _, err := village.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *VillageController) parent(id int) {
	country := new(models.Country)
	result, count := country.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["id"] = v.Id
		row["name"] = v.Name
		if v.Id == id {
			row["selected"] = true
		} else {
			row["selected"] = false
		}
		list[k] = row
	}
	this.Data["group"] = list
}

func (this *VillageController) Table() {
	village := new(models.Village)
	result, count := village.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["name"] = v.Name
		row["file"] = v.Icon
		country := new(models.Country)
		country.Id = v.ParentId
		if err := country.Query(); err == nil {
			row["parent"] = country.Name
		}
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
	this.parent(0)
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
	this.parent(group.ParentId)
	row["id"] = group.Id
	row["name"] = group.Name
	row["file"] = group.Icon
	this.setFileSize(row, group.Icon)
	this.Data["row"] = row
	this.display(this.getBgAreaAction("group/edit"))
}

func (this *GroupController) AjaxSave() {
	group := new(models.Group)
	group.Name = this.getString("name", "名称不能为空!", 1)
	group.Icon = this.getString("file", "File不能为空!", defaultMinSize)
	group.Id = this.getInt("id", 0)
	if group.Id == 0 {
		group.ParentId = this.getInt("group_id", 0)
		group.CreateId = this.userId
		group.CreateTime = time.Now().Unix()
		if _, err := group.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	group.ParentId = this.getInt("group_id", 0)
	group.UpdateId = this.userId
	group.UpdateTime = time.Now().Unix()
	if _, err := group.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *GroupController) parent(id int) {
	village := new(models.Village)
	result, count := village.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["id"] = v.Id
		row["name"] = v.Name
		if v.Id == id {
			row["selected"] = true
		} else {
			row["selected"] = false
		}
		list[k] = row
	}
	this.Data["group"] = list
}

func (this *GroupController) Table() {
	group := new(models.Group)
	result, count := group.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["name"] = v.Name
		row["file"] = v.Icon
		village := new(models.Village)
		village.Id = v.ParentId
		if err := village.Query(); err == nil {
			row["parent"] = village.Name
		}
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
	this.parent(0)
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
	this.parent(team.ParentId)
	row["id"] = team.Id
	row["name"] = team.Name
	row["file"] = team.Icon
	this.setFileSize(row, team.Icon)
	this.Data["row"] = row
	this.display(this.getBgAreaAction("team/edit"))
}

func (this *TeamController) AjaxSave() {
	team := new(models.Team)
	team.Id = this.getInt("id", 0)
	team.Name = this.getString("name", "名称不能为空!", 1)
	team.Icon = this.getString("file", "File不能为空!", defaultMinSize)
	if team.Id == 0 {
		team.ParentId = this.getInt("group_id", 0)
		team.CreateId = this.userId
		team.CreateTime = time.Now().Unix()
		if _, err := team.Add(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}
	team.ParentId = this.getInt("group_id", 0)
	team.UpdateId = this.userId
	team.UpdateTime = time.Now().Unix()
	if _, err := team.Update(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *TeamController) parent(id int) {
	group := new(models.Group)
	result, count := group.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["id"] = v.Id
		row["name"] = v.Name
		if v.Id == id {
			row["selected"] = true
		} else {
			row["selected"] = false
		}
		list[k] = row
	}
	this.Data["group"] = list
}

func (this *TeamController) Table() {
	team := new(models.Team)
	result, count := team.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		row["name"] = v.Name
		row["file"] = v.Icon
		group := new(models.Group)
		group.Id = v.ParentId
		if err := group.Query(); err == nil {
			row["parent"] = group.Name
		}
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
