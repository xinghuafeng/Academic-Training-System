package roleService

import (
	"github.com/gin-gonic/gin"
	"sauth/model/roleModel"
	"sauth/db"
	"sauth/util"
	"database/sql"
)

func All() ([]roleModel.Role, error) {
	model := make([]roleModel.Role, 0)
	err := db.Engine.Where("app_id = ?", "AM").Find(&model)
	return model, err
}

func Find(ctx *gin.Context) ([]roleModel.Role, error) {
	model := make([]roleModel.Role, 0)
	param := ctx.Query("param")
	err := db.Engine.Where("app_id = ?", param).Find(&model)
	return model, err
}

func Save(ctx *gin.Context) (int64, error) {
	var model roleModel.Role
	var res int64
	var err error
	err = ctx.Bind(&model)
	if nil != err {
		return 0, err
	}
	if "" != model.Uuid {
		res, err = db.Engine.Update(&model, &roleModel.Role{Uuid: model.Uuid})
	} else {
		model.Uuid = util.GetUuid()
		res, err = db.Engine.Insert(model)
	}
	return res, err
}

func Delete(ctx *gin.Context) (sql.Result, error) {
	var model roleModel.Role
	err := ctx.Bind(&model)
	if nil != err {
		return nil, err
	}
	res, err := roleModel.Delete(model.Uuid)
	return res, err
}
