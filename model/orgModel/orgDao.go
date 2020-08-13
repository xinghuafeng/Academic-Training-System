package orgModel

import (
	"sauth/db"
	"database/sql"
)

func Delete(uuid string) (sql.Result, error) {
	sql := `DELETE FROM org
			WHERE uuid = ?`
	res, err := db.Engine.Exec(sql, uuid)
	return res, err
}
