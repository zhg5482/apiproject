package services

import (
	"apiproject/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

func Login(o orm.Ormer,username, password string)(models.User,error){
	var user models.User
	var sql = "select * from user where username = ? and password = "
	err := o.Raw(sql,username,password).QueryRow(&user)
	fmt.Printf("ERR: %v\n", err)
	return user,err
}

func GetAllUsers(o orm.Ormer)([]models.User,error) {
	var users [] models.User
	var sql = "select * from user limit 10"
	num,err := o.Raw(sql).QueryRows(&users)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	return users,err
}

func AddUser(o orm.Ormer,user models.User) (int64,error) {
	return o.Insert(&user)
}

func GetUserById(o orm.Ormer,id int) (models.User,error) {
	var user models.User
	var sql = "select * from user where id = ?"
	err := o.Raw(sql,id).QueryRow(&user)
	fmt.Printf("GetUser ERR: %v\n", err)
	return user,err
}

func GetUserByQuery(o orm.Ormer,sql string,paras [] interface{})([] models.User,error)  {
	var users [] models.User
	num,err := o.Raw(sql,paras).QueryRows(&users)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	return users,err
}

func UpdateUser(o orm.Ormer,user models.User) (int64,error) {
	return o.Update(&user,"user_name")
}