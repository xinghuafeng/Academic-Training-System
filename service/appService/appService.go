package appService

import (
	"Academic-Training-System/db"
	"Academic-Training-System/model/appModel"
	"Academic-Training-System/util"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func All() ([]appModel.Application, error) {
	model := make([]appModel.Application, 0)
	err := db.Engine.Find(&model)
	return model, err
}

func Find(ctx *gin.Context) ([]appModel.Application, error) {
	model := make([]appModel.Application, 0)
	var err error
	param := ctx.Query("param")
	if "" == param {
		err = db.Engine.Find(&model)
	} else {
		param = `%` + param + `%`
		err = db.Engine.Where("id LIKE ? OR name LIKE ?", param).Find(&model)
	}
	return model, err
}

func Save(ctx *gin.Context) (int64, error) {
	var model appModel.Application
	var res int64
	var err error
	err = ctx.Bind(&model)
	if nil != err {
		return 0, err
	}
	if "" != model.Uuid {
		res, err = db.Engine.Update(&model, &appModel.Application{Uuid: model.Uuid})
	} else {
		model.Uuid = util.GetUuid()
		res, err = db.Engine.Insert(&model)
	}
	return res, err
}

func Delete(ctx *gin.Context) (sql.Result, error) {
	var model appModel.Application
	err := ctx.Bind(&model)
	if nil != err {
		return nil, err
	}
	res, err := appModel.Delete(model.Uuid)
	return res, err
}
