package lib

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

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