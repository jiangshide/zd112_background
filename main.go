package main

import (
	_ "zd112/routers"
	"github.com/astaxie/beego"
	"zd112/models"
)

func main() {
	models.Init()
	beego.Run()
}
