package controllers

import (
	"apiproject/lib"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

//切换数据库
func ChangeDb()  {
	//o := orm.NewOrm()
	//o.Using("default") // 默认使用 default，你可以指定为其他数据库
}

//redis 测试
func RedisTest()  {
	res := lib.RedisSet("egew","rtr",10*time.Second)
	fmt.Println(res,"===========")

	//fmt.Println("========",value)
	beego.Run()
}

//初始化数据
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
