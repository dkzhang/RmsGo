package applicationDB

import (
	"github.com/dkzhang/RmsGo/webapi/model/application"
)

type ApplicationDB interface {
	QueryApplicationByID(applicationID int) (application.Application, error)
	QueryApplicationByOwner(userID int) []application.Application
	QueryApplicationByDepartmentCode(dc string) []application.Application
	QueryApplicationByFilter(userFilter func(application.Application) bool) []application.Application

	InsertApplication(applicationInfo application.Application) (appID int, err error)
	UpdateApplication(applicationInfo application.Application) (err error)
	InsertAppOps(appOps application.AppOpsRecord) (err error)
}
