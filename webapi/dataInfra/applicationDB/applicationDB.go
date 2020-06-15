package applicationDB

import (
	"github.com/dkzhang/RmsGo/webapi/model/resourceApplication"
)

type ApplicationDB interface {
	QueryApplicationByID(applicationID int) (resourceApplication.Application, error)
	QueryApplicationByOwner(userID int) []resourceApplication.Application
	QueryApplicationByDepartmentCode(dc string) []resourceApplication.Application
	QueryApplicationByFilter(userFilter func(resourceApplication.Application) bool) []resourceApplication.Application

	InsertApplication(applicationInfo resourceApplication.Application) (err error)
}
