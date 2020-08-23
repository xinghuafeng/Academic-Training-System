package userRoleModel

import "Academic-Training-System/db"

func FindUserRole(userId string) ([]map[string][]byte, error) {
	sql := `SELECT
				t.user_id,
				t.role_id,
				t.app_id,
			    t1.name AS appName,
				t2.name AS roleName
			from user_role t
			left join application t1 on t1.id = t.app_id
			left join role t2 on t2.uuid = t.role_id
			where t.user_id = ?`
	res, err := db.Engine.Query(sql, userId)
	return res, err
}
