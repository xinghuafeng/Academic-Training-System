package userModel

import (
	"Academic-Training-System/db"
	"Academic-Training-System/util"
	"database/sql"
)

func Insert(model *User) (sql.Result, error) {
	sql := `INSERT INTO user
			(uuid, user_name, password, nick_name, sex, position, tel, mail, address, remark, create_user, create_time, update_user, update_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	res, err := db.Engine.Exec(sql,
		util.GetUuid(), model.UserName, model.Password, model.NickName, model.Sex, model.Position, model.Tel, model.Mail, model.Address, model.Remark, model.CreateUser, util.CurrentTime(), model.UpdateUser, util.CurrentTime())
	return res, err
}

func FindCountByUserNameAndPasswordAndTenantId(userName string, password string, tenantId string) ([]map[string][]byte, error) {
	sql := `SELECT COUNT(*) AS count FROM user
			where user_name = ? and password = ? and tenant_id = ?`
	res, err := db.Engine.Query(sql,
		userName, password, tenantId)
	return res, err
}

func FindCountByUserName(userName string) ([]map[string][]byte, error) {
	sql := `SELECT COUNT(*) AS count FROM user
			where user_name = ?`
	res, err := db.Engine.Query(sql, userName)
	return res, err
}

func Delete(uuid string) (sql.Result, error) {
	sql := `DELETE FROM user
			WHERE uuid = ?`
	res, err := db.Engine.Exec(sql, uuid)
	return res, err
}
