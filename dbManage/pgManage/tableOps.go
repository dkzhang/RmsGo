package pgManage

import (
	"database/sql"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/jmoiron/sqlx"
)

var tableList = map[string]string{
	"user_info": user.SchemaUser,
}

func createTable(db *sqlx.DB, schema string) (result sql.Result, err error) {
	result, err = db.Exec(schema)
	return result, err
}

func dropTable(db *sqlx.DB, tableName string) (result sql.Result, err error) {
	exec := `DROP Table ` + tableName
	result, err = db.Exec(exec)
	return result, err
}

func createAllTable(db *sqlx.DB) {
	for name, scheme := range tableList {
		dropTable(db, name)
		createTable(db, scheme)
	}
}
