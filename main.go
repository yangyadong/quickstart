package main

import (
	"quickstart/common"
	_ "quickstart/routers"
	"github.com/astaxie/beego"
)

func main() {
	common.InitConfig()
	common.InitMysql()
	defer common.Db.Close()
	common.InitRedis()
	defer common.RedisClient.Close()
	beego.Run()
}

