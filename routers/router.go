package routers

import (
	"zd112/controllers"
	"github.com/astaxie/beego"
)

func init() {

    beego.Router("/",&controllers.IndexController{})
    beego.Router("/login",&controllers.UserController{},"*:Login")
    beego.Router("/reg",&controllers.UserController{},"*:Reg")
    beego.Router("/admin",&controllers.UserController{},"*:Admin")
    beego.Router("/loginOut",&controllers.UserController{},"*:LoginOut")
    beego.Router("/backstage",&controllers.BackstageController{},"*:Index")
    beego.Router("/backstage/home/start",&controllers.BackstageController{},"*:Start")
    beego.Router("/backstage/user/edit",&controllers.UserController{},"*:Edit")

    beego.Router("/backstage/admin/list",&controllers.AdminController{},"*:List")
    beego.Router("/backstage/admin/add",&controllers.AdminController{},"*:Add")
    beego.Router("/backstage/admin/edit",&controllers.AdminController{},"*:Edit")
    beego.Router("/backstage/admin/table",&controllers.AdminController{},"*:Table")

    beego.AutoRouter(&controllers.AdminController{})

    beego.Router("/backstage/role/list",&controllers.RoleController{},"*:List")
    beego.Router("/backstage/role/add",&controllers.RoleController{},"*:Add")
    beego.Router("/backstage/role/edit",&controllers.RoleController{},"*:Edit")
    beego.Router("/backstage/role/table",&controllers.RoleController{},"*:Table")
    beego.AutoRouter(&controllers.RoleController{})

}
