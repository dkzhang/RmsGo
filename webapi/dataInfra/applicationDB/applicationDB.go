package applicationDB

import (
	"github.com/dkzhang/RmsGo/webapi/model/application"
)

type ApplicationDB interface {
	QueryApplicationByID(applicationID int) (application.Application, error)
	QueryApplicationByOwner(userID int) ([]application.Application, error)
	QueryApplicationByDepartmentCode(dc string) ([]application.Application, error)
	QueryApplicationByFilter(appFilter func(application.Application) bool) ([]application.Application, error)

	InsertApplication(applicationInfo application.Application) (appID int, err error)
	UpdateApplication(applicationInfo application.Application) (err error)

	InsertAppOps(record application.AppOpsRecord) (recordID int, err error)
	QueryAppOpsByAppId(applicationID int) (records []application.AppOpsRecord, err error)

	Close()
}
