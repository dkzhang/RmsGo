package postgreOps

import (
	"database/sql"
	"fmt"
	"github.com/dkzhang/RmsGo/datebaseCommon/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectToDatabase(pg security.PgSecurity) (db *sqlx.DB, err error) {
	dataSourceStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		pg.Host, pg.Port, pg.User, pg.Password, pg.DbName, pg.Sslmode)
	db, err = sqlx.Open("postgres", dataSourceStr)
	return db, err
}

func CreateTable(db *sqlx.DB, schema string) (result sql.Result, err error) {
	result, err = db.Exec(schema)
	return result, err
}

func DropTable(db *sqlx.DB, tableName string) (result sql.Result, err error) {
	exec := `DROP Table ` + tableName
	result, err = db.Exec(exec)
	return result, err
}
