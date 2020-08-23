package systemService

import (
	"Academic-Training-System/db"
	"Academic-Training-System/model/SystemModel"
	"github.com/gin-gonic/gin"
)

func All(ctx *gin.Context) ([]SystemModel.Baseinfo, error) {
	model := make([]SystemModel.Baseinfo, 0)
	var err error
	param := ctx.Query("param")
	if "" == param {
		err = db.Engine.Find(&model)
	} else {
		param = `%` + param + `%`
		err = db.Engine.Where("basename = ?", param).Find(&model)
	}
	return model, err

}
