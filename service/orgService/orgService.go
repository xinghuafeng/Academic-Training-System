package orgService

import (
	"github.com/gin-gonic/gin"
	"sauth/model/orgModel"
	"sauth/db"
	"sauth/util"
	"database/sql"
)

func Find(ctx *gin.Context) ([]orgModel.Org, error) {
	model := make([]orgModel.Org, 0)
	err := db.Engine.Find(&model)
	return model, err
}

// 保存机构
func Save(ctx *gin.Context) (int64, error) {
	var model orgModel.Org
	var res int64
	var err error
	err = ctx.Bind(&model)
	if nil != err {
		return -1, err
	}
	userName, _ := ctx.Cookie(util.AuthUserCookieId)
	if "" != model.Uuid { // update
		model.UpdateUser = userName
		res, err = db.Engine.Update(&model, &orgModel.Org{Uuid: model.Uuid})
	} else { // insert
		model.Uuid = util.GetUuid()
		model.Id = model.TenantId + `-` + model.OrgId // Id = TenantId + OrgId
		model.CreateUser = userName
		model.UpdateUser = userName

		res, err = db.Engine.Insert(&model)
	}
	return res, err
}

// 删除机构
func Delete(ctx *gin.Context) (sql.Result, error) {
	var model orgModel.Org
	err := ctx.Bind(&model)
	if nil != err {
		return nil, err
	}
	res, err := orgModel.Delete(model.Uuid)
	return res, err
}
