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
	QueryApplicationByID(applicationID int) (application.Application, error)
	QueryApplicationByOwner(userID int, appType int, appStatus int) ([]application.Application, error)
	QueryApplicationByDepartmentCode(dc string, appType int, appStatus int) ([]application.Application, error)
	QueryApplicationAll(appType int, appStatus int) ([]application.Application, error)
	QueryApplicationByFilter(appFilter func(application.Application) bool) ([]application.Application, error)

	InsertApplication(applicationInfo application.Application) (appID int, err error)
	UpdateApplication(applicationInfo application.Application) (err error)

	InsertAppOps(record application.AppOpsRecord) (recordID int, err error)
	QueryAppOpsByAppId(applicationID int) (records []application.AppOpsRecord, err error)

	ArchiveToHistory(historyADI DBInfo, projectID int) (err error)

	Close()
}

type ApplicationHistoryDB interface {
	QueryApplicationByID(applicationID int) (application.Application, error)
	QueryApplicationByOwner(userID int, appType int, appStatus int) ([]application.Application, error)
	QueryApplicationByDepartmentCode(dc string, appType int, appStatus int) ([]application.Application, error)
	QueryApplicationAll(appType int, appStatus int) ([]application.Application, error)
	QueryApplicationByFilter(appFilter func(application.Application) bool) ([]application.Application, error)

	QueryAppOpsByAppId(applicationID int) (records []application.AppOpsRecord, err error)

	Close()
}
