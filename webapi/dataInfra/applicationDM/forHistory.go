package applicationDM

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/applicationDB"
	"github.com/dkzhang/RmsGo/webapi/model/application"
)

type ForHistory struct {
	theDB applicationDB.ApplicationDB
}

func NewForHistory(adb applicationDB.ApplicationDB, theLogMap logMap.LogMap) (nFHis ForHistory, err error) {
	nFHis.theDB = adb
	return nFHis, nil
}

func (fhis ForHistory) Close() {
	fhis.theDB.Close()
}

func (fhis ForHistory) QueryByID(appID int) (application.Application, error) {
	app, err := fhis.theDB.QueryByID(appID)
	if err != nil {
		return application.Application{}, fmt.Errorf("QueryByID in db error: %v", err)
	}
	return app, nil
}

func (fhis ForHistory) QueryByOwner(userID int, appType int, appStatus int) (apps []application.Application, err error) {
	apps, err = fhis.theDB.QueryByOwner(userID, appType, appStatus)
	if err != nil {
		return nil, fmt.Errorf("QueryByOwner in db error: %v", err)
	}

	return apps, nil
}
func (fhis ForHistory) QueryByDepartmentCode(dc string, appType int, appStatus int) (apps []application.Application, err error) {
	apps, err = fhis.theDB.QueryByDepartmentCode(dc, appType, appStatus)
	if err != nil {
		return nil, fmt.Errorf("QueryByDepartmentCode in db error: %v", err)
	}

	return apps, nil
}
func (fhis ForHistory) QueryAll(appType int, appStatus int) (apps []application.Application, err error) {
	apps, err = fhis.theDB.QueryAll(appType, appStatus)
	if err != nil {
		return nil, fmt.Errorf("QueryAll in db error: %v", err)
	}

	return apps, nil
}
func (fhis ForHistory) QueryByFilter(appFilter func(application.Application) bool) (apps []application.Application, err error) {
	appsALL, err := fhis.theDB.QueryAll(application.AppTypeALL, application.AppStatusALL)
	if err != nil {
		return nil, fmt.Errorf("QueryAll in db error: %v", err)
	}

	for _, app := range appsALL {
		if appFilter(app) == true {
			apps = append(apps, app)
		}
	}

	return apps, nil
}

func (fhis ForHistory) QueryAppOpsByAppId(appID int) (records []application.AppOpsRecord, err error) {
	records, err = fhis.theDB.QueryAppOpsByAppId(appID)
	if err != nil {
		return nil, fmt.Errorf("QueryAppOpsByAppId from db error: %v", err)
	}
	return records, nil
}
