package controllers

import (
	"fmt"
	"quickstart/common"
	"quickstart/model"
	"testing"
)

func TestGetUSer(t *testing.T) {
	common.InitConfig()
	common.InitMysql()
	defer common.Db.Close()
	common.InitRedis()
	defer common.RedisClient.Close()
	user := model.GetUser("15112345678")
	fmt.Println(user)
}
