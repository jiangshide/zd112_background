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
}
