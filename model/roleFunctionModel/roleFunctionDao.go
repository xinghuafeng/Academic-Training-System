package roleFunctionModel

import (
	"database/sql"
	"sauth/db"
)

/*
	删除
*/
func Delete(roleId string) (sql.Result, error) {
	sql := `DELETE FROM role_function
			WHERE role_id = ?`
	res, err := db.Engine.Exec(sql, roleId)
	return res, err
}
