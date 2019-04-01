package main

import (
	_ "apiproject/routers"             //路由
	"github.com/astaxie/beego"         //框架
	"github.com/astaxie/beego/orm"     //orm
	_ "github.com/go-sql-driver/mysql" //mssql
)

func init()  {
	//注册数据库驱动
	orm.RegisterDriver("mysql",orm.DRMySQL)
	//注册默认数据库
	orm.RegisterDataBase("default","mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	orm.SetMaxIdleConns("default",1000)
	orm.SetMaxOpenConns("default",2000)
	//orm.DefaultTimeLoc = time.Local //时区设置
	//自动化建表
	//orm.RunSyncdb("default", false, true)

}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
}
