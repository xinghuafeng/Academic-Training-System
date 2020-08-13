package orgModel

import "time"

/*
	机构
*/
type Org struct {
	Uuid       string                     // 主键
	Id         string                     // 唯一 ID = OrgId + TenantId
	Name       string                     // 机构名称
	Pid        string                     // 上级机构编号
	OrgId      string                     // 机构 ID
	TenantId   string                     // 租户 ID
	Comment    string                     // 说明
	CreateTime time.Time `xorm:"created"` // 创建时间
	CreateUser string                     // 创建者
	UpdateTime time.Time `xorm:"updated"` // 修改时间
	UpdateUser string                     // 修改者
}
