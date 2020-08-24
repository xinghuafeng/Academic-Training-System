package systemService

import (
	"Academic-Training-System/db"
	"Academic-Training-System/model/SystemModel"
	"database/sql"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetBaseinfo(ctx *gin.Context) ([]SystemModel.Baseinfo, error) {
	model := make([]SystemModel.Baseinfo, 0)
	var err error
	param := ctx.Query("param")
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNo, _ := strconv.Atoi(ctx.Query("pageNo"))
	if "" == param {
		err = db.Engine.Limit(pageSize, (pageNo-
			1)*pageSize).Find(&model)
	} else {
		err = db.Engine.Where("basename = ?", param).Limit(pageSize, (pageNo-1)*pageSize).Find(&model)
	}

	return model, err

}

// 删除校区
func DeleteBaseinfo(ctx *gin.Context) (sql.Result, error) {
	var model SystemModel.Baseinfo
	model.Id = ctx.Query("id")
	// 数据绑定
	// Delete
	res, err := SystemModel.DeleteBaseinfo(model.Id)
	return res, err
}

// 保存校区或者修改校区
func InsertOrUpdateBaseInfo(ctx *gin.Context) (int64, error) {
	var model SystemModel.Baseinfo
	var res int64
	var err error
	err = ctx.Bind(&model)
	if nil != err {
		return -1, err
	}

	if "" != model.Id { // update

		res, err = db.Engine.Update(&model, &SystemModel.Baseinfo{Id: model.Id})

	} else { // insert
		res, err = db.Engine.Insert(&model)
	}
	return res, err
}
