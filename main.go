package main

import (
	"quickstart/common"
	_ "quickstart/routers"
	"github.com/astaxie/beego"
)

func main() {
	common.InitConfig()
	common.InitMysql()
	common.InitRedis()
	beego.Run()
}

