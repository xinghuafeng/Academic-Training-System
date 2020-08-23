package roleFunctionModel

import (
	"Academic-Training-System/model/roleModel"
	"time"
)

/*
	角色-功能
*/
type RoleFunction struct {
	Uuid       string    // 主键
	RoleId     string    // 角色ID
	RoleName   string    // 角色名称
	TenantId   string    // 租户ID
	FunctionId string    // 功能ID
	CreateUser string    // 创建者
	CreateTime time.Time `xorm:"created"` // 创建时间
	UpdateUser string    // 修改者
	UpdateTime time.Time `xorm:"updated"` // 修改时间
}

/*
	web 端发送的数据结构
*/
type PostData struct {
	FunctionData []string       `json:"functionData""`
	RoleData     roleModel.Role `json:"roleData"`
}
