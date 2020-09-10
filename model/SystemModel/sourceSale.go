package SystemModel

/*
	baseinfo 销售来源
*/
type XiaoSalesourceinfo struct {
	Id         string `xorm:"pk autoincr" form:"Id" json:"Id"`
	Baseid     string `form:"Baseid" json:"Baseid"`
	Sourcename string `form:"Sourcename" json:"Sourcename"`
	Adddata    string `form:"Adddata" json:"Adddata"`
	Remark     int    `form:"Remark" json:"Remark"`
}
