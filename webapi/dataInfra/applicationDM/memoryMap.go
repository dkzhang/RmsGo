package applicationDM

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/applicationDB"
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"github.com/sirupsen/logrus"
	"time"
)

type MemoryMap struct {
	theAppMap map[int]*application.Application
	theDB     applicationDB.ApplicationDB
}

func NewMemoryMap(adb applicationDB.ApplicationDB, theLogMap logMap.LogMap) (nmm MemoryMap, err error) {
	nmm.theDB = adb
	nmm.theAppMap = make(map[int]*application.Application)

	apps, err := nmm.theDB.QueryAll(application.AppTypeALL, application.AppStatusALL)
	if err != nil {
		return MemoryMap{},
			fmt.Errorf("generate new MemoryMap failed since ApplicationDB.GetAllArray error: %v", err)
	}
	theLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"len(apps)": len(apps),
	}).Info("NewMemoryMap ApplicationDB.GetAllArray success.")

	for i := range apps {
		nmm.theAppMap[apps[i].ApplicationID] = &apps[i]
	}

	theLogMap.Log(logMap.NORMAL).Info("NewMemoryMap load data to map success.")

	return nmm, nil
}

func (adm MemoryMap) Close() {
	adm.theDB.Close()
}

func (adm MemoryMap) QueryByID(appID int) (application.Application, error) {
	if app, ok := adm.theAppMap[appID]; ok {
		return *app, nil
	} else {
		return application.Application{}, fmt.Errorf("application (id=%d) does not exist", appID)
	}
}

func (adm MemoryMap) QueryByOwner(userID int, appType int, appStatus int) (apps []application.Application, err error) {
	for _, app := range adm.theAppMap {
		if app.ApplicantUserID == userID {
			if appType&app.Type != 0 && appStatus&app.Status != 0 {
				apps = append(apps, *app)
			}
		}
	}

	return apps, nil
}
func (adm MemoryMap) QueryByDepartmentCode(dc string, appType int, appStatus int) (apps []application.Application, err error) {
	for _, app := range adm.theAppMap {
		if app.DepartmentCode == dc {
			if appType&app.Type != 0 && appStatus&app.Status != 0 {
				apps = append(apps, *app)
			}
		}
	}

	return apps, nil
}
func (adm MemoryMap) QueryAll(appType int, appStatus int) (apps []application.Application, err error) {
	for _, app := range adm.theAppMap {
		if appType&app.Type != 0 && appStatus&app.Status != 0 {
			apps = append(apps, *app)
		}
	}

	return apps, nil
}
func (adm MemoryMap) QueryByFilter(appFilter func(application.Application) bool) (apps []application.Application, err error) {
	for _, app := range adm.theAppMap {
		if appFilter(*app) == true {
			apps = append(apps, *app)
		}
	}

	return apps, nil
}

func (adm MemoryMap) QueryAppOpsByAppId(appID int) (records []application.AppOpsRecord, err error) {
	records, err = adm.theDB.QueryAppOpsByAppId(appID)
	if err != nil {
		return nil, fmt.Errorf("QueryAppOpsByAppId from db error: %v", err)
	}
	return records, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////

func (adm MemoryMap) Insert(appInfo application.Application) (appID int, err error) {
	appInfo.CreatedAt = time.Now()
	appInfo.UpdatedAt = time.Now()
	appID, err = adm.theDB.Insert(appInfo)
	if err != nil {
		return -1, fmt.Errorf("Insert in db error: %v", err)
	}

	appInfo.ApplicationID = appID
	adm.theAppMap[appID] = &appInfo

	return appID, nil
}

func (adm MemoryMap) Update(appInfo application.Application) (err error) {
	if _, ok := adm.theAppMap[appInfo.ApplicationID]; !ok {
		return fmt.Errorf("application id %d does not exist", appInfo.ApplicationID)
	}

	appInfo.UpdatedAt = time.Now()
	err = adm.theDB.Update(appInfo)
	if err != nil {
		return fmt.Errorf("Update in db error: %v", err)
	}
	adm.theAppMap[appInfo.ApplicationID] = &appInfo

	return nil
}

func (adm MemoryMap) InsertAppOps(record application.AppOpsRecord) (recordID int, err error) {
	record.CreatedAt = time.Now()
	recordID, err = adm.theDB.InsertAppOps(record)
	if err != nil {
		return -1, fmt.Errorf("InsertAppOps in db error: %v", err)
	}

	return recordID, nil
}

func (adm MemoryMap) ArchiveToHistory(historyADI applicationDB.DBInfo, projectID int) (err error) {
	// TODO
	return fmt.Errorf("unimplemented ")
}
