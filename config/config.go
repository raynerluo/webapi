package config

import (
	"github.com/go-ini/ini"
)

var (
	Config *ini.File
)

//初始化配置文件
func init() {
	var err error
	Config, err = ini.Load("./config/config.ini")
	if err != nil {
		panic(err)
	}
}
