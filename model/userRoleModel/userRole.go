package userRoleModel

import (
	"Academic-Training-System/model/roleModel"
	"time"
)

// 用户-角色
type UserRole struct {
	Uuid       string    // 主键
	UserId     string    // 用户 ID
	RoleId     string    // 角色 ID
	AppId      string    // 应用 ID
	CreateTime time.Time `xorm:"created"` // 创建时间
	CreateUser string    // 创建者
	UpdateTime time.Time `xorm:"updated"` // 修改时间
	UpdateUser string    // 修改者
}

type PostData struct {
	RoleData     []roleModel.Role `json:"roleData"`
	UserName     string           `json:"userName"`
	UserTenantId string           `json:"userTenantId"`
}
