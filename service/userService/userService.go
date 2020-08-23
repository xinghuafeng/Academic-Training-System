package userService

import (
	"Academic-Training-System/db"
	"Academic-Training-System/model/userModel"
	"Academic-Training-System/util"
	"database/sql"
	"github.com/gin-gonic/gin"
)

// 保存用户
func Save(ctx *gin.Context) (int64, error) {
	var model userModel.User
	var res int64
	var err error

	// 数据绑定
	err = ctx.Bind(&model)
	if nil != err {
		return -1, err
	}

	// 从 cookie 中获取执行此次操作的用户名
	userNameInCookie, _ := ctx.Cookie(util.AuthUserCookieId)

	// 用户ID = '租户ID' - '登录名'
	model.Id = model.TenantId + `-` + model.UserName

	// update or insert
	if "" != model.Uuid {

		// 修改者 = userNameInCookie(执行此次操作的用户名)
		model.UpdateUser = userNameInCookie

		// Update
		res, err = db.Engine.Update(&model, &userModel.User{Uuid: model.Uuid})

	} else {
		model.Uuid = util.GetUuid()

		// 加密
		// 初始密码 = 用户名
		model.Password = util.Md5Encrypt(model.UserName)

		// 创建者(修改者) = userNameInCookie(执行此次操作的用户名)
		model.CreateUser = userNameInCookie
		model.UpdateUser = userNameInCookie

		// Insert
		res, err = db.Engine.Insert(&model)
	}
	return res, err
}

// 删除用户
func Delete(ctx *gin.Context) (sql.Result, error) {
	var model userModel.User

	// 数据绑定
	err := ctx.Bind(&model)
	if nil != err {
		return nil, err
	}

	// Delete
	res, err := userModel.Delete(model.Uuid)

	return res, err
}

func All(ctx *gin.Context) ([]userModel.User, error) {
	var err error
	orgId := ctx.Query("orgId")
	model := make([]userModel.User, 0)
	if "TOPALL" == orgId || "" == orgId {
		err = db.Engine.Find(&model)
	} else {
		db.Engine.Where("org_id = ?", orgId).Find(&model)
	}
	return model, err
}
