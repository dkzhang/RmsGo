package applicationDB

import (
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"github.com/jmoiron/sqlx"
)

type DBInfo struct {
	TheDB        *sqlx.DB
	AppTableName string
	OpsTableName string
}

type ApplicationDB interface {
	ApplicationHistoryDB

	Insert(applicationInfo application.Application) (appID int, err error)
	Update(applicationInfo application.Application) (err error)

	InsertAppOps(record application.AppOpsRecord) (recordID int, err error)

	ArchiveToHistory(historyADI DBInfo, projectID int) (err error)

	Close()
}

type ApplicationHistoryDB interface {
	QueryByID(applicationID int) (application.Application, error)
	QueryByOwner(userID int, appType int, appStatus int) ([]application.Application, error)
	QueryByDepartmentCode(dc string, appType int, appStatus int) ([]application.Application, error)
	QueryAll(appType int, appStatus int) ([]application.Application, error)

	QueryAppOpsByAppId(applicationID int) (records []application.AppOpsRecord, err error)

	Close()
}
