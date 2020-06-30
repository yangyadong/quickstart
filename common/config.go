package common

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/config"
	"os"
)

type System struct {
	Mode string `mapstructure:"mode" json:"mode" ini:"mode"`
}

type Log struct {
	Prefix  string `mapstructure:"prefix" json:"prefix" ini:"prefix"`
	LogFile bool   `mapstructure:"log-file" json:"log-file" ini:"log-file" yaml:"log-file" toml:"log-file"`
	Stdout  string `mapstructure:"stdout" json:"stdout" ini:"stdout"`
	File    string `mapstructure:"file" json:"file" ini:"file"`
}

type Config struct {
	System System `json:"system" ini:"system"`
	Log Log `json:"log" ini:"log"`
}

var (
	CONFIG = new(config.Config)
)

func InitConfig() {
	// 打开文件
	file, _ := os.Open("conf/config.json")
	// 关闭文件
	defer file.Close()
	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	err := decoder.Decode(&CONFIG)
	if err != nil {
		panic(err)
	}
	fmt.Println(CONFIG)
}