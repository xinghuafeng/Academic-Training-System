package SystemModel

import (
	"Academic-Training-System/db"
	"database/sql"
)

func DeleteSourceSales(uuid string) (sql.Result, error) {
	sql := `DELETE FROM xiao_salesourceinfo
			WHERE ID = ?`
	res, err := db.Engine.Exec(sql, uuid)
	return res, err
}
