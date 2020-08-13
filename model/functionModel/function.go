package functionModel

import "time"

type Function struct {
	Uuid       string                     // 主键
	Name       string                     // 名称
	Type       string                     // 类型
	Key        string                     // ID
	Url        string                     // url
	Class      string                     // class选择器
	Method     string                     // Method
	Params     string                     // Params
	PKey       string                     // parent ID
	Lvl        int                        // 层级
	Sort       int                        // 排序
	CreateUser string                     // 创建者
	CreateTime time.Time `xorm:"created"` // 创建时间
	UpdateUser string                     // 修改者
	UpdateTime time.Time `xorm:"updated"` // 修改时间
}
