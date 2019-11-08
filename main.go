package main

import (
	_ "apiproject/routers" //路由
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego" //框架
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"     //orm
	_ "github.com/go-sql-driver/mysql" //mysql
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
	//controllers.RedisTest()
	initLogger()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.SetStaticPath("/swagger", "swagger")
		//beego.BConfig.WebConfig.DirectoryIndex = true
		//beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

// logs
func initLogger()(err error) {
	config := make(map[string]interface{})
	config["filename"] = beego.AppConfig.String("log_path")

	// map 转 json
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("initLogger failed, marshal err:", err)
		return
	}
	// log 的配置
	beego.SetLogger(logs.AdapterFile, string(configStr))
	// log打印文件名和行数
	beego.SetLogFuncCall(true)
	fmt.Println(string(configStr))
	return
}