package main

import (
	"apiproject/controllers"
	_ "apiproject/routers"             //路由
	"github.com/astaxie/beego"         //框架
	"github.com/astaxie/beego/orm"     //orm
	_ "github.com/go-sql-driver/mysql" //mssql
)

func init()  {
	mysql_dataSource := beego.AppConfig.String("mysql.dataSource")
	//注册数据库驱动
	orm.RegisterDriver("mysql",orm.DRMySQL)
	//注册默认数据库
	orm.RegisterDataBase("default","mysql", mysql_dataSource)
	orm.SetMaxIdleConns("default",1000)
	orm.SetMaxOpenConns("default",2000)
	//orm.DefaultTimeLoc = time.Local //时区设置
	//自动化建表
	//orm.RunSyncdb("default", false, true)
	controllers.RedisTest()

}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
