package services

import (
	"apiproject/models"
	"github.com/astaxie/beego/orm"
)

func AddProfile(o orm.Ormer,profile models.Profile) (int64,error) {
	return o.Insert(&profile)
}
