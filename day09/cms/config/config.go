package config

import (
	"os"
	"encoding/json"
)

//服务端配置
type AppConfig struct {
	AppName    string `json:"app_name"`
	Port       string `json:"port"`
	StaticPath string `json:"static_path"`
	Mode       string `json:"mode"`
}

var ServConfig AppConfig

//初始化服务器配置
func InitConfig() *AppConfig {
	file, err := os.Open("/Users/liyaochu/mygo/src/studygo/day09/cms/config.json")
	if err != nil {
		panic(err.Error())
	}
	//json 解析器解析到AppConfig的结构体中
	decoder := json.NewDecoder(file)
	conf := AppConfig{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err.Error())
	}
	return &conf
}
