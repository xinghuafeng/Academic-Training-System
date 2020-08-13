package roleModel

import (
	"database/sql"
	"sauth/db"
)

func Delete(uuid string) (sql.Result, error) {
	sql := `DELETE FROM role
			WHERE uuid = ?`
	res, err := db.Engine.Exec(sql, uuid)
	return res, err
}
