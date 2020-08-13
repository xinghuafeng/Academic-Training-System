package appModel

import (
	"time"
)

/*
	app 应用信息
*/
type Application struct {
	Uuid       string                     // 主键
	Id         string                     // 应用编号
	Name       string                     // 应用名称
	Url        string                     // url
	Origin     string                     // origin
	Admin      string                     // 管理员
	Comment    string                     // 说明
	CreateUser string                     // 创建者
	CreateTime time.Time `xorm:"created"` // 创建时间
	UpdateUser string                     // 修改者
	UpdateTime time.Time `xorm:"updated"` // 修改时间
}
