package SystemModel

import (
	"Academic-Training-System/db"
	"database/sql"
)

func DeleteBaseinfo(uuid string) (sql.Result, error) {
	sql := `DELETE FROM baseinfo
			WHERE ID = ?`
	res, err := db.Engine.Exec(sql, uuid)
	return res, err
}
