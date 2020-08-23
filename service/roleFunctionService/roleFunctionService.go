package roleFunctionService

import (
	"Academic-Training-System/db"
	"Academic-Training-System/model/roleFunctionModel"
	"Academic-Training-System/util"
	"github.com/gin-gonic/gin"
)

// 保存角色-功能
func Save(ctx *gin.Context) (int64, error) {
	var postData roleFunctionModel.PostData // web端发送的json对应的数据结构 = 角色 + 功能
	err := ctx.BindJSON(&postData)          // json 与数据结构绑定
	if nil != err {
		return -1, err
	}

	functionData := postData.FunctionData // 功能
	roleData := postData.RoleData         // 角色

	// 开启事务
	session := db.Engine.NewSession()                                                   /* init */
	defer session.Close()                                                               /* close */
	err = session.Begin()                                                               /* begin */
	_, err = session.Exec("DELETE from role_function WHERE role_id = ?", roleData.Uuid) // 删除当前角色的所有关系
	if err != nil {
		session.Rollback() /* rollback */
		return -1, err
	}

	roleFunctions := make([]roleFunctionModel.RoleFunction, len(functionData)) // 进行批量新增的数据结构
	for i, functionKey := range functionData {
		roleFunctions[i].Uuid = util.GetUuid()
		roleFunctions[i].RoleId = roleData.Uuid
		roleFunctions[i].RoleName = roleData.Name
		roleFunctions[i].FunctionId = functionKey
	}
	res, err := session.Insert(&roleFunctions) // 批量新增
	if nil != err {
		session.Rollback() /* rollback */
		return -1, err
	}

	err = session.Commit() /* commit */
	if nil != err {
		return -1, err
	}

	return res, err
}

func Find(ctx *gin.Context) ([]roleFunctionModel.RoleFunction, error) {
	roleId := ctx.Query("param")
	model := make([]roleFunctionModel.RoleFunction, 0)
	err := db.Engine.Where("role_id = ?", roleId).Find(&model)
	return model, err
}
