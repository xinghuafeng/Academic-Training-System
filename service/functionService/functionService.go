package functionService

import (
	"sauth/model/functionModel"
	"sauth/db"
	"github.com/gin-gonic/gin"
	"sauth/util"
)

func All() ([]functionModel.Function, error) {
	model := make([]functionModel.Function, 0)
	err := db.Engine.Where("lvl = ?", 1).Find(&model)
	return model, err
}

func Find(ctx *gin.Context) ([]map[string]interface{}, error) {
	param := ctx.Query("param")
	if "" == param {
		param = "#"
	}
	res, err := functionModel.FindTree(param)
	if nil != err {
		return nil, err
	}
	data := util.DataTrans(res)
	return data, err
}
