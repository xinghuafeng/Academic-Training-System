package SystemModel

/*
	baseinfo 校區信息
*/
type Baseinfo struct {
	Id            string `xorm:"pk autoincr" form:"Id" json:"Id"`
	Baseid        string `form:"Baseid" json:"Baseid"`
	Basename      string `form:"Basename" json:"Basename"`
	Code          string `form:"Code" json:"Code"`
	Personcharge  string `form:"Personcharge" json:"Personcharge"`
	Tel           string `form:"Tel" json:"Tel"`
	Address       string `form:"Address" json:"Address"`
	Basetype      string `form:"Basetype" json:"Basetype"`
	Iplimit       int    `form:"Iplimit" json:"Iplimit"`
	Fulltimecheck int    `form:"Fulltimecheck" json:"Fulltimecheck"`
	Parttimesign  int    `form:"Parttimesign" json:"Parttimesign"`
	Isparent      int    `form:"Isparent" json:"Isparent"`
}
