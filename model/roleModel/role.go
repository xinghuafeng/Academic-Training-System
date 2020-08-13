package roleModel

import "time"

/*
	role 角色信息
*/
type Role struct {
	Uuid       string                     // 主键
	Name       string                     // 角色名称
	Desc       string                     // 角色描述
	AppId      string                     // 应用Id
	TenantId   string                     // 租户Id
	CreateTime time.Time `xorm:"created"` // 创建时间
	CreateUser string                     // 创建者
	UpdateTime time.Time `xorm:"updated"` // 修改时间
	UpdateUser string                     // 修改者
}
