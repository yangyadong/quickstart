package common

import "testing"

func TestInitConfig(t *testing.T) {
	InitConfig()
	InitMysql()
	InitRedis()
}
