package applicationDM

import (
	"github.com/dkzhang/RmsGo/webapi/dataInfra/applicationDB"
	"github.com/dkzhang/RmsGo/webapi/model/application"
)

type ApplicationDM interface {
	ApplicationHistoryDM

	Insert(applicationInfo application.Application) (appID int, err error)
	Update(applicationInfo application.Application) (err error)

	InsertAppOps(record application.AppOpsRecord) (recordID int, err error)

	ArchiveToHistory(historyADI applicationDB.DBInfo, projectID int) (err error)
}

type ApplicationHistoryDM interface {
	QueryByID(applicationID int) (application.Application, error)
	QueryByOwner(userID int, appType int, appStatus int) ([]application.Application, error)
	QueryByDepartmentCode(dc string, appType int, appStatus int) ([]application.Application, error)
	QueryAll(appType int, appStatus int) ([]application.Application, error)
	QueryByFilter(appFilter func(application.Application) bool) ([]application.Application, error)

	QueryAppOpsByAppId(applicationID int) (records []application.AppOpsRecord, err error)

	Close()
}
