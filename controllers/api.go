package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

func HttpSign() {
	var u ApiController
	//fmt.Println(string(u.Ctx.Input.RequestBody))
	fmt.Println(u.Ctx.Request)
}