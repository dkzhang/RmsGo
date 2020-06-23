package applicationDB

import "github.com/jmoiron/sqlx"

type ApplicationPg struct {
	db *sqlx.DB
}

func NewApplicationPg(sqlxdb *sqlx.DB) ApplicationPg {
	return ApplicationPg{db: sqlxdb}
}
