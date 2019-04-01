package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "apiproject/routers"
	"github.com/astaxie/beego"
	"time"
)

func init()  {
	//注册数据库驱动
	orm.RegisterDriver("mysql",orm.DRMySQL)
	//注册默认数据库
	orm.RegisterDataBase("default","mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	orm.SetMaxIdleConns("default",1000)
	orm.SetMaxOpenConns("default",2000)
	orm.DefaultTimeLoc = time.Local
	//注册model
	//orm.RegisterModel(new(models.User),new(models.Profile),new(models.Post),new(models.Tag))
	//初始化ormer
	//O = orm.NewOrm()
	//自动化建表
	//orm.RunSyncdb("default", false, true)
	//初始化数据
	//datainit()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//o := orm.NewOrm()
	//o.Using("default") // 默认使用 default，你可以指定为其他数据库
	beego.Run()
}


func datainit() {
	 //rel  : 自动生成外键为 表名_id
	 //sql1 := "insert into user (name,profile_id) values ('ming',1),('hua',2),('qiang',3);"
	 //sql2 := "insert into profile (age) values (16),(14),(15);"
	 //sql3 := "insert into tag (name) values ('offical'),('beta'),('dev');"
	 //sql4 := "insert into post (title,user_id) values ('paper1',1),('paper2',1),('paper3',2),('paper4',3),('paper5',3);"
	 //// m2m 生成的 表名：子表_主表s  主键自增
	 //sql5 := "insert into post_tags (tag_id, post_id) values (1,1),(1,3),(2,2),(3,3),(2,4),(3,4),(3,5); "
	 //
	 ////使用Raw（）.Exec（）执行sql
	 //O.Raw(sql1).Exec()
	 //O.Raw(sql2).Exec()
	 //O.Raw(sql3).Exec()
	 //O.Raw(sql4).Exec()
	 //O.Raw(sql5).Exec()
}