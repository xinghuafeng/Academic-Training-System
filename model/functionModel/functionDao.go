package functionModel

import "Academic-Training-System/db"

func FindTree(PKey string) ([]map[string][]byte, error) {
	sql := `SELECT T.Uuid , T.Name , T.Type ,T.Key , IFNULL(T.Url, '') AS Url , IFNULL(T.Class, '') AS Class , IFNULL(T.Method, '') AS Method , IFNULL(T.Params, '') AS Params , T.p_key AS PKey , T.Lvl , IFNULL(T.Sort, '') AS Sort,
				(
					CASE WHEN (SELECT COUNT(*) FROM function WHERE p_key = T.Key) > 0
					THEN 'true'
					ELSE 'false'
					END
				) AS hasChild
			FROM function T
			WHERE T.p_key = ?`
	res, err := db.Engine.Query(sql, PKey)
	return res, err
}
