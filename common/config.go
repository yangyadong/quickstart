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
	// 打开文件
	file, _ := os.Open("D:/workspace/quickstart/conf/config.json")
	// 关闭文件
	defer file.Close()
	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	err := decoder.Decode(&ConfigInfo)
	if err != nil {
		panic(err)
	}
}