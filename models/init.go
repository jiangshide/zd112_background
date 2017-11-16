package models

import (
	"github.com/astaxie/beego"
	"net/url"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	dbHost := beego.AppConfig.String("db.host")
	dbPort := beego.AppConfig.String("db.port")
	dbUser := beego.AppConfig.String("db.user")
	dbPsw := beego.AppConfig.String("db.psw")
	dbName := beego.AppConfig.String("db.name")
	timeZone := beego.AppConfig.String("db.timezone")
	if dbPort == "" {
		dbPort = "3306"
	}
	dsn := dbUser + ":" + dbPsw + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	if timeZone != "" {
		dsn += "&loc=" + url.QueryEscape(timeZone)
	}
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel()
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}