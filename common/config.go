package common

import (
	"encoding/json"
	"os"
)

type Redis struct {
	Addr string `mapstructure:"addr" json:"addr" ini:"addr"`
	Pwd string `mapstructure:"pwd" json:"pwd" ini:"pwd"`
}

type Mysql struct {
	Dialect string `mapstructure:"dialect" json:"dialect" ini:"dialect"`
	Addr string `mapstructure:"addr" json:"addr" ini:"addr"`
	User string `mapstructure:"user" json:"user" ini:"user"`
	Pwd string `mapstructure:"pwd" json:"pwd" ini:"pwd"`
	Name string `mapstructure:"name" json:"name" ini:"name"`
}

type Config struct {
	Redis Redis `json:"redis" ini:"redis"`
	Mysql Mysql `json:"mysql" ini:"mysql"`
}

var (
	ConfigInfo = new(Config)
)

func InitConfig() {
	file, _ := os.Open("D:/workspace/quickstart/conf/config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&ConfigInfo)
	if err != nil {
		panic(err)
	}
}