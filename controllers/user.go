package controllers

import (
	"apiproject/models"
	"apiproject/services"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)


// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users,err := services.GetAllUsers(orm.NewOrm())
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = users
	}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	fmt.Println(uid,"Get=====",u.Ctx.Request)
	if uid != "" {
		uid,err := strconv.Atoi(uid) //类型转换
		user, err := services.GetUserById(orm.NewOrm(),uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	userName := u.Input().Get("username")
	age := u.GetString("age")
	u.Data["json"] = "insert failed!"
	if age != "" && userName != "" {
		db := orm.NewOrm()
		err := db.Begin()   //开启事务
		//插入子表
		age,err:=strconv.Atoi(age)
		profile := models.Profile{Age:age}
		id,err := services.AddProfile(db,profile)
		if err != nil {
			db.Rollback()
		}
		//插入父表
		user := models.User{Name:userName,Profile:&models.Profile{Id:int(id)}}
		//_,err1 := db.Insert(&user) //全局insert
		_,err1 := services.AddUser(db,user)
		if err1 != nil {
			db.Rollback()
		} else {
			u.Data["json"] = "insert success!"
		}
		err = db.Commit()
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	username := u.GetString("username")
	if uid != ""  && username!= ""{
		uid,err := strconv.Atoi(uid)
		user := models.User{Id:uid,Name:username}
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := services.UpdateUser(orm.NewOrm(),user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	fmt.Println(username,"Login=====",u.Ctx.Request)
	u.Data["json"] = "login failed!"
	if username != ""{
		user,err := services.Login(orm.NewOrm(),username)
		if err == nil {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

