package db

import (
	"github.com/go-redis/redis"
	"github.com/astaxie/beego/logs"
	"sauth/util"
)

var RedisClient *redis.Client

const (
	RedisInitErrMsg    = "redis初始化失败"
	RedisInitFinishMsg = "redis初始化成功......\t response of ping:"
)

func init() {
	config := util.ParseConfigFile() // 解析配置文件
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisConnectAddress,
		Password: config.RedisConnectPassword,
	})
	res, err := client.Ping().Result() // ping
	if nil != err {
		logs.Error(RedisInitErrMsg, err)
	} else {
		logs.Info(RedisInitFinishMsg, res)
	}
	RedisClient = client
}
