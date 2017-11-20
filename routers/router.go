package routers

import (
	"zd112/controllers"
	"github.com/astaxie/beego"
)

func init() {

    beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.LoginController{})
}
