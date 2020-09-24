package webapi

import (
	pgSM "github.com/dkzhang/RmsGo/ResourceSM/databaseInit/pgOpsSqlx"
	"github.com/dkzhang/RmsGo/databaseInit/pgOpsSqlx"
	"github.com/jmoiron/sqlx"
)

func ResetALL(db *sqlx.DB) {
	pgOpsSqlx.CreateAllTable(db)
	pgOpsSqlx.SeedAllTable(db)
	pgSM.CreateAllTable(db)
	pgSM.SeedAllTable(db)
}
