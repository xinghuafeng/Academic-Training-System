package systemService

import (
	"Academic-Training-System/db"
	"Academic-Training-System/model/SystemModel"
	"database/sql"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetSourceSales(ctx *gin.Context) ([]SystemModel.XiaoSalesourceinfo, error) {
	model := make([]SystemModel.XiaoSalesourceinfo, 0)
	var err error
	param := ctx.Query("param")
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNo, _ := strconv.Atoi(ctx.Query("pageNo"))
	if "" == param {
		err = db.Engine.Limit(pageSize, (pageNo-1)*pageSize).Find(&model)
	} else {
		err = db.Engine.Where("sourceName = ?", param).Limit(pageSize, (pageNo-1)*pageSize).Find(&model)
	}

	return model, err

}

// 删除销售来源
func DeleteSourceSales(ctx *gin.Context) (sql.Result, error) {
	var model SystemModel.XiaoSalesourceinfo
	model.Id = ctx.Query("id")
	// 数据绑定
	// Delete
	res, err := SystemModel.DeleteSourceSales(model.Id)
	return res, err
}

// 保存校区或者修改销售来源
func UpdateOrInsertSourceSales(ctx *gin.Context) (int64, error) {
	var model SystemModel.XiaoSalesourceinfo
	var res int64
	var err error
	err = ctx.Bind(&model)
	if nil != err {
		return -1, err
	}

	if "" != model.Id { // update

		res, err = db.Engine.Update(&model, &SystemModel.XiaoSalesourceinfo{Id: model.Id})

	} else { // insert
		res, err = db.Engine.Insert(&model)
	}
	return res, err
}
