package meteringDB

import (
	"github.com/dkzhang/RmsGo/ResourceSM/model/metering"
	"github.com/jmoiron/sqlx"
)

type DBInfo struct {
	TheDB     *sqlx.DB
	TableName string
}

type MeteringDB interface {
	MeteringReadOnlyDB

	Insert(ms metering.Statement) (mid int, err error)
}

type MeteringReadOnlyDB interface {
	Query(projectID int, mType int, typeInfo string) (ms metering.Statement, err error)
	IsExist(projectID int, mType int, typeInfo string) (isExist bool, err error)
	QueryAll(projectID int, mType int) (mss []metering.Statement, err error)
	Close()
}
