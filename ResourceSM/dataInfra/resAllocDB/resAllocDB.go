package resAllocDB

import (
	"github.com/dkzhang/RmsGo/ResourceSM/model/resAlloc"
	"github.com/jmoiron/sqlx"
)

type DBInfo struct {
	TheDB     *sqlx.DB
	TableName string
}

type ResAllocDB interface {
	QueryByID(recordID int) (rar resAlloc.Record, err error)
	QueryByProjectID(projectID int) ([]resAlloc.Record, error)
	QueryAll() ([]resAlloc.Record, error)
	Insert(node resAlloc.Record) (err error)
	Close()
}
