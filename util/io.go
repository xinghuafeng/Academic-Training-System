package util

import (
	"github.com/astaxie/beego/logs"
	//"github.com/gin-gonic/gin/json"
	"encoding/json"
	"os"
)

const (
	ConfigFileName         = "config/config.json"
	OpenConfigFileErrMsg   = "打开配置文件失败"
	DecodeConfigFileErrMsg = "解析配置文件失败"
)

// 配置信息
type Config struct {
	DriverName           string
	DataSourceName       string
	RedisConnectAddress  string
	RedisConnectPassword string
}

// 解析配置文件
func ParseConfigFile() *Config {
	file, err := os.Open(ConfigFileName)
	if nil != err {
		logs.Error(OpenConfigFileErrMsg, err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var config Config
	err = decoder.Decode(&config)
	if nil != err {
		logs.Error(DecodeConfigFileErrMsg, err)
	}
	return &config
}
