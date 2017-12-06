package routers

import (
	"zd112/controllers"
	"github.com/astaxie/beego"
	"time"
	"zd112/utils"
)

func init() {

	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login", &controllers.UserController{}, "*:Login")
	beego.Router("/reg", &controllers.UserController{}, "*:Reg")
	beego.Router("/admin", &controllers.UserController{}, "*:Admin")
	beego.Router("/loginOut", &controllers.UserController{}, "*:LoginOut")
	beego.Router("/backstage", &controllers.BackstageController{}, "*:Index")
	beego.Router("/backstage/home/start", &controllers.BackstageController{}, "*:Start")
	beego.Router("/backstage/user/edit", &controllers.UserController{}, "*:Edit")

	beego.Router("/backstage/admin/list", &controllers.AdminController{}, "*:List")
	beego.Router("/backstage/admin/add", &controllers.AdminController{}, "*:Add")
	beego.Router("/backstage/admin/edit", &controllers.AdminController{}, "*:Edit")
	beego.Router("/backstage/admin/table", &controllers.AdminController{}, "*:Table")

	beego.AutoRouter(&controllers.AdminController{})

	beego.Router("/backstage/role/list", &controllers.RoleController{}, "*:List")
	beego.Router("/backstage/role/add", &controllers.RoleController{}, "*:Add")
	beego.Router("/backstage/role/edit", &controllers.RoleController{}, "*:Edit")
	beego.Router("/backstage/role/table", &controllers.RoleController{}, "*:Table")
	beego.AutoRouter(&controllers.RoleController{})

	beego.Router("/backstage/auth/list", &controllers.AuthController{}, "*:List")
	//beego.Router("/backstage/auth/add",&controllers.AuthController{},"*:Add")
	//beego.Router("/backstage/auth/edit",&controllers.AuthController{},"*:Edit")
	//beego.Router("/backstage/auth/table",&controllers.AuthController{},"*:Table")
	beego.AutoRouter(&controllers.AuthController{})

	beego.Router("/backstage/nation/list", &controllers.NationController{}, "*:List")
	beego.Router("/backstage/nation/add", &controllers.NationController{}, "*:Add")
	beego.Router("/backstage/nation/edit", &controllers.NationController{}, "*:Edit")
	beego.Router("/backstage/nation/table", &controllers.NationController{}, "*:Table")
	beego.Router("/backstage/nation/upload", &controllers.NationController{}, "*:Upload")
	beego.Router("/backstage/nation/ajaxdel", &controllers.NationController{}, "*:AjaxDel")
	beego.AutoRouter(&controllers.NationController{})

	beego.Router("/backstage/area/continent", &controllers.ContinentController{}, "*:List")
	beego.Router("/backstage/area/continent/add", &controllers.ContinentController{}, "*:Add")
	beego.Router("/backstage/area/continent/edit", &controllers.ContinentController{}, "*:Edit")
	beego.Router("/backstage/area/continent/table", &controllers.ContinentController{}, "*:Table")
	beego.Router("/backstage/area/continent/ajaxsave", &controllers.ContinentController{}, "*:AjaxSave")
	beego.Router("/backstage/area/continent/ajaxdel", &controllers.ContinentController{}, "*:AjaxDel")

	beego.Router("/backstage/area/state", &controllers.StateController{}, "*:List")
	beego.Router("/backstage/area/state/add", &controllers.StateController{}, "*:Add")
	beego.Router("/backstage/area/state/table", &controllers.StateController{}, "*:Table")
	beego.Router("/backstage/area/state/edit", &controllers.StateController{}, "*:Edit")
	beego.Router("/backstage/area/state/ajaxsave", &controllers.StateController{}, "*:AjaxSave")
	beego.Router("/backstage/area/state/ajaxdel", &controllers.StateController{}, "*:AjaxDel")

	beego.Router("/backstage/area/province", &controllers.ProvinceController{}, "*:List")
	beego.Router("/backstage/area/province/add", &controllers.ProvinceController{}, "*:Add")
	beego.Router("/backstage/area/province/edit", &controllers.ProvinceController{}, "*:Edit")
	beego.Router("/backstage/area/province/table", &controllers.ProvinceController{}, "*:Table")
	beego.Router("/backstage/area/province/ajaxsave", &controllers.ProvinceController{}, "*:AjaxSave")
	beego.Router("/backstage/area/province/ajaxdel", &controllers.ProvinceController{}, "*:AjaxDel")

	beego.Router("/backstage/area/city", &controllers.CityController{}, "*:List")
	beego.Router("/backstage/area/city/add", &controllers.CityController{}, "*:Add")
	beego.Router("/backstage/area/city/edit", &controllers.CityController{}, "*:Edit")
	beego.Router("/backstage/area/city/table", &controllers.CityController{}, "*:Table")
	beego.Router("/backstage/area/city/ajaxsave", &controllers.CityController{}, "*:AjaxSave")
	beego.Router("/backstage/area/city/ajaxdel", &controllers.CityController{}, "*:AjaxDel")

	beego.Router("/backstage/area/region", &controllers.RegionController{}, "*:List")
	beego.Router("/backstage/area/region/add", &controllers.RegionController{}, "*:Add")
	beego.Router("/backstage/area/region/edit", &controllers.RegionController{}, "*:Edit")
	beego.Router("/backstage/area/region/table", &controllers.RegionController{}, "*:Table")
	beego.Router("/backstage/area/region/ajaxsave", &controllers.RegionController{}, "*:AjaxSave")
	beego.Router("/backstage/area/region/ajaxdel", &controllers.RegionController{}, "*:AjaxDel")

	beego.Router("/backstage/area/county", &controllers.CountyController{}, "*:List")
	beego.Router("/backstage/area/county/add", &controllers.CountyController{}, "*:Add")
	beego.Router("/backstage/area/county/edit", &controllers.CountyController{}, "*:Edit")
	beego.Router("/backstage/area/county/table", &controllers.CountyController{}, "*:Table")
	beego.Router("/backstage/area/county/ajaxsave", &controllers.CountyController{}, "*:AjaxSave")
	beego.Router("/backstage/area/county/ajaxdel", &controllers.CountyController{}, "*:AjaxDel")

	beego.Router("/backstage/area/town", &controllers.TownController{}, "*:List")
	beego.Router("/backstage/area/town/add", &controllers.TownController{}, "*:Add")
	beego.Router("/backstage/area/town/edit", &controllers.TownController{}, "*:Edit")
	beego.Router("/backstage/area/town/table", &controllers.TownController{}, "*:Table")
	beego.Router("/backstage/area/town/ajaxsave", &controllers.TownController{}, "*:AjaxSave")
	beego.Router("/backstage/area/town/ajaxdel", &controllers.TownController{}, "*:AjaxDel")

	beego.Router("/backstage/area/country", &controllers.CountryController{}, "*:List")
	beego.Router("/backstage/area/country/add", &controllers.CountryController{}, "*:Add")
	beego.Router("/backstage/area/country/edit", &controllers.CountryController{}, "*:Edit")
	beego.Router("/backstage/area/country/table", &controllers.CountryController{}, "*:Table")
	beego.Router("/backstage/area/country/ajaxsave", &controllers.CountryController{}, "*:AjaxSave")
	beego.Router("/backstage/area/country/ajaxdel", &controllers.CountryController{}, "*:AjaxDel")

	beego.Router("/backstage/area/village", &controllers.VillageController{}, "*:List")
	beego.Router("/backstage/area/village/add", &controllers.VillageController{}, "*:Add")
	beego.Router("/backstage/area/village/edit", &controllers.VillageController{}, "*:Edit")
	beego.Router("/backstage/area/village/table", &controllers.VillageController{}, "*:Table")
	beego.Router("/backstage/area/village/ajaxsave", &controllers.VillageController{}, "*:AjaxSave")
	beego.Router("/backstage/area/village/ajaxdel", &controllers.VillageController{}, "*:AjaxDel")

	beego.Router("/backstage/area/group", &controllers.GroupController{}, "*:List")
	beego.Router("/backstage/area/group/add", &controllers.GroupController{}, "*:Add")
	beego.Router("/backstage/area/group/edit", &controllers.GroupController{}, "*:Edit")
	beego.Router("/backstage/area/group/table", &controllers.GroupController{}, "*:Table")
	beego.Router("/backstage/area/group/ajaxsave", &controllers.GroupController{}, "*:AjaxSave")
	beego.Router("/backstage/area/group/ajaxdel", &controllers.GroupController{}, "*:AjaxDel")

	beego.Router("/backstage/area/team", &controllers.TeamController{}, "*:List")
	beego.Router("/backstage/area/team/add", &controllers.TeamController{}, "*:Add")
	beego.Router("/backstage/area/team/edit", &controllers.TeamController{}, "*:Edit")
	beego.Router("/backstage/area/team/table", &controllers.TeamController{}, "*:Table")
	beego.Router("/backstage/area/team/ajaxsave", &controllers.TeamController{}, "*:AjaxSave")
	beego.Router("/backstage/area/team/ajaxdel", &controllers.TeamController{}, "*:AjaxDel")

	beego.Router("/backstage/web/banner", &controllers.BannerController{}, "*:List")
	beego.Router("/backstage/web/banner/add", &controllers.BannerController{}, "*:Add")
	beego.Router("/backstage/web/banner/edit", &controllers.BannerController{}, "*:Edit")
	beego.Router("/backstage/web/banner/table", &controllers.BannerController{}, "*:Table")
	beego.Router("/backstage/web/banner/ajaxsave", &controllers.BannerController{}, "*:AjaxSave")
	beego.Router("/backstage/web/banner/ajaxdel", &controllers.BannerController{}, "*:AjaxDel")

	beego.Router("/upload", &controllers.BaseController{}, "*:Upload")
	beego.ErrorController(&controllers.ErrorController{})
	//taskTime()
}

func taskTime() {
	ticker := time.NewTicker(time.Minute * 1)
	go func() {
		for _ = range ticker.C {
			res, err := utils.ExecCommand("/usr/bin/git --git-dir=" + utils.GetCurrentDir("") + "/.git checkout master")
			beego.Info("res:", string(res), " | err:", err)
		}
	}()
}
