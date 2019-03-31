package main

import (
	_ "apiproject/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init()  {
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	orm.SetMaxIdleConns("default",1000)
	orm.SetMaxOpenConns("default",2000)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
