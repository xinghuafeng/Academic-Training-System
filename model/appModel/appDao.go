package appModel

import (
	"Academic-Training-System/db"
	"database/sql"
)

/*
	删除应用
*/
func Delete(uuid string) (sql.Result, error) {
	sql := `DELETE FROM application
			WHERE uuid = ?`
	res, err := db.Engine.Exec(sql, uuid)
	return res, err
}

/*
	select2 中的选项
*/
func FindOptions() ([]map[string][]byte, error) {
	sql := `SELECT id, name as text FROM application`
	res, err := db.Engine.Query(sql)
	return res, err
}
