package userRoleService

import (
	"github.com/gin-gonic/gin"
	"sauth/model/userRoleModel"
	"sauth/db"
	"sauth/util"
)

func Find(ctx *gin.Context) ([]map[string]interface{}, error) {
	userId := ctx.Query("userId")
	res, err := userRoleModel.FindUserRole(userId)
	data := util.DataTrans(res)
	return data, err
}

// 保存用户-角色
func Save(ctx *gin.Context) (int64, error) {
	var postData userRoleModel.PostData
	err := ctx.BindJSON(&postData)
	if nil != err {
		return -1, err
	}

	roleData := postData.RoleData
	userId := postData.UserTenantId + `-` + postData.UserName

	session := db.Engine.NewSession()
	defer session.Close()
	err = session.Begin()
	_, err = session.Exec("DELETE from user_role WHERE user_id = ?", userId)
	if nil != err {
		session.Rollback()
		return -1, err
	}

	userRole := make([]userRoleModel.UserRole, len(roleData))
	for i, role := range roleData {
		userRole[i].Uuid = util.GetUuid()
		userRole[i].UserId = userId
		userRole[i].RoleId = role.Uuid
		userRole[i].AppId = role.AppId
	}

	res, err := session.Insert(&userRole)
	if nil != err {
		session.Rollback()
		return -1, err
	}

	err = session.Commit()
	if nil != err {
		return -1, err
	}

	return res, err
}
