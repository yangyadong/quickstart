package main

import (
	"quickstart/common"
	_ "quickstart/routers"
	"github.com/astaxie/beego"
)

func main() {
	common.InitRedis()
	common.InitMysql()
	beego.Run()
}

