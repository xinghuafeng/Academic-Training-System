package systemService

import (
	"Academic-Training-System/db"
	"Academic-Training-System/model/SystemModel"
	"github.com/gin-gonic/gin"
	"strconv"
)

func All(ctx *gin.Context) ([]SystemModel.Baseinfo, error) {
	model := make([]SystemModel.Baseinfo, 0)
	var err error
	param := ctx.Query("param")
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNo, _ := strconv.Atoi(ctx.Query("pageNo"))
	if "" == param {
		err = db.Engine.Limit(pageSize, (pageNo-1)*pageSize).Find(&model)
	} else {
		err = db.Engine.Where("basename = ?", param).Limit(pageSize, (pageNo-1)*pageSize).Find(&model)
	}

	return model, err

}
