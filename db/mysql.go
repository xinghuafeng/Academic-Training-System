package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/astaxie/beego/logs"
	"sauth/util"
)

var Engine *xorm.Engine

func init() {
	var err error
	config := util.ParseConfigFile()
	Engine, err = xorm.NewEngine(config.DriverName, config.DataSourceName)
	if nil != err {
		logs.Error("mysql引擎初始化失败", err)
	} else {
		logs.Info("mysql 初始化成功......")
	}
	Engine.ShowSQL(true)
}
