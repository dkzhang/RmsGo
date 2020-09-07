package projectResDB

import (
	"github.com/dkzhang/RmsGo/ResourceSM/model/projectRes"
	"github.com/jmoiron/sqlx"
)

type DBInfo struct {
	TheDB     *sqlx.DB
	TableName string
}

type ProjectResDB interface {
	QueryByID(projectID int) (pr projectRes.ResInfo, err error)
	QueryAll() ([]projectRes.ResInfo, error)
	Insert(pr projectRes.ResInfo) (err error)
	Update(pr projectRes.ResInfo) (err error)
	Delete(projectID int) (err error)
	Close()
}
