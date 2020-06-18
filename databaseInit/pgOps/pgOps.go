package pgOps

import (
	"database/sql"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/userDB"
	"github.com/dkzhang/RmsGo/webapi/model/appTemp"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/jmoiron/sqlx"
)

var tableList = map[string]string{
	"user_info":             user.SchemaUser,
	"application_temporary": appTemp.SchemaAppTemp,
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

func CreateAllTable(db *sqlx.DB) {
	for name, scheme := range tableList {
		dropTable(db, name)
		createTable(db, scheme)
	}
}

func SeedAllTable(db *sqlx.DB) {
	seedUserTable(db)
}

func ImportFromFile(db *sqlx.DB) {

}

func seedUserTable(db *sqlx.DB) {
	theUserDB := userDB.NewUserInPostgre(db)
	theUserDB.InsertUser(user.UserInfo{
		UserName:       "ctrl-zhj001",
		ChineseName:    "张俊001",
		Department:     "调度小组",
		DepartmentCode: "ctrl",
		Section:        "测试",
		Mobile:         "15383026353",
		Role:           user.RoleController,
		Status:         user.StatusNormal,
		Remarks:        "测试用账号",
	})

	theUserDB.InsertUser(user.UserInfo{
		UserName:       "ctrl-zxq",
		ChineseName:    "翟修齐",
		Department:     "调度小组",
		DepartmentCode: "ctrl",
		Section:        "测试",
		Mobile:         "18699622740",
		Role:           user.RoleController,
		Status:         user.StatusNormal,
		Remarks:        "测试用账号for翟",
	})
}
