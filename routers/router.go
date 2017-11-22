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
}
