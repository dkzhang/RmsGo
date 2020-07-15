package applicationDB

import (
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"github.com/jmoiron/sqlx"
	"time"
)

type ApplicationPg struct {
	ApplicationDbInfo
}

func NewApplicationPg(sqlxdb *sqlx.DB, appname, opsname string) ApplicationPg {
	return ApplicationPg{
		ApplicationDbInfo{
			TheDB:        sqlxdb,
			AppTableName: appname,
			OpsTableName: opsname,
		},
	}
}

func (apg ApplicationPg) Close() {
	apg.TheDB.Close()
}

func (apg ApplicationPg) QueryApplicationByID(applicationID int) (app application.Application, err error) {
	queryByID := fmt.Sprintf(`SELECT * FROM %s WHERE application_id=$1`, apg.AppTableName)
	err = apg.TheDB.Get(&app, queryByID, applicationID)
	if err != nil {
		return application.Application{},
			fmt.Errorf("QueryApplicationByID in TheDB error: %v", err)
	}
	return app, nil
}

func (apg ApplicationPg) QueryApplicationByOwner(userID int, appType int, appStatus int) (apps []application.Application, err error) {
	var queryByOwner string
	if appType == application.AppTypeALL {
		if appStatus == application.AppStatusALL {
			queryByOwner = fmt.Sprintf(`SELECT * FROM %s WHERE app_user_id=$1`, apg.AppTableName)
			err = apg.TheDB.Select(&apps, queryByOwner, userID)
			if err != nil {
				return nil, fmt.Errorf("QueryGeneralFormDraftByOwner (appType=%d AND appStatus=%d) from TheDB error: %v", err, appType, appStatus)
			}
			return apps, nil
		} else {
			queryByOwner = fmt.Sprintf(`SELECT * FROM %s WHERE app_user_id=$1 AND status=$2`, apg.AppTableName)
			err = apg.TheDB.Select(&apps, queryByOwner, userID, appStatus)
			if err != nil {
				return nil, fmt.Errorf("QueryGeneralFormDraftByOwner (appType=%d AND appStatus=%d) from TheDB error: %v", err, appType, appStatus)
			}
			return apps, nil
		}
	} else {
		if appStatus == application.AppStatusALL {
			queryByOwner = fmt.Sprintf(`SELECT * FROM %s WHERE app_user_id=$1 AND application_type=$2`, apg.AppTableName)
			err = apg.TheDB.Select(&apps, queryByOwner, userID, appType)
			if err != nil {
				return nil, fmt.Errorf("QueryGeneralFormDraftByOwner (appType=%d AND appStatus=%d) from TheDB error: %v", err, appType, appStatus)
			}
			return apps, nil
		} else {
			queryByOwner = fmt.Sprintf(`SELECT * FROM %s WHERE app_user_id=$1 AND application_type=$2 AND status=$3`, apg.AppTableName)
			err = apg.TheDB.Select(&apps, queryByOwner, userID, appType, appStatus)
			if err != nil {
				return nil, fmt.Errorf("QueryGeneralFormDraftByOwner (appType=%d AND appStatus=%d) from TheDB error: %v", err, appType, appStatus)
			}
			return apps, nil
		}
	}
}

func (apg ApplicationPg) QueryApplicationByDepartmentCode(dc string, appType int, appStatus int) (apps []application.Application, err error) {
	var queryByDC string
	if appType == application.AppTypeALL {
		if appStatus == application.AppStatusALL {
			queryByDC = fmt.Sprintf(`SELECT * FROM %s WHERE department_code=$1`, apg.AppTableName)
			err = apg.TheDB.Select(&apps, queryByDC, dc)
			if err != nil {
				return nil, fmt.Errorf("QueryApplicationByDepartmentCode (appType=%d AND appStatus=%d) from TheDB error: %v", err, appType, appStatus)
			}
			return apps, nil
		} else {
			queryByDC = fmt.Sprintf(`SELECT * FROM %s WHERE department_code=$1 AND status=$2`, apg.AppTableName)
			err = apg.TheDB.Select(&apps, queryByDC, dc, appStatus)
			if err != nil {
				return nil, fmt.Errorf("QueryApplicationByDepartmentCode (appType=%d AND appStatus=%d) from TheDB error: %v", err, appType, appStatus)
			}
			return apps, nil
		}
	} else {
		if appStatus == application.AppStatusALL {
			queryByDC = fmt.Sprintf(`SELECT * FROM %s WHERE department_code=$1 AND application_type=$2`, apg.AppTableName)
			err = apg.TheDB.Select(&apps, queryByDC, dc, appType)
			if err != nil {
				return nil, fmt.Errorf("QueryApplicationByDepartmentCode (appType=%d AND appStatus=%d) from TheDB error: %v", err, appType, appStatus)
			}
			return apps, nil
		} else {
			queryByDC = fmt.Sprintf(`SELECT * FROM %s WHERE department_code=$1 AND application_type=$2 AND status=$3`, apg.AppTableName)
			err = apg.TheDB.Select(&apps, queryByDC, dc, appType, appStatus)
			if err != nil {
				return nil, fmt.Errorf("QueryApplicationByDepartmentCode (appType=%d AND appStatus=%d) from TheDB error: %v", err, appType, appStatus)
			}
			return apps, nil
		}
	}
}

func (apg ApplicationPg) QueryApplicationAll(appType int, appStatus int) (apps []application.Application, err error) {
	var queryALL string
	if appType == application.AppTypeALL {
		if appStatus == application.AppStatusALL {
			queryALL = fmt.Sprintf(`SELECT * FROM %s`, apg.AppTableName)
			err = apg.TheDB.Select(&apps, queryALL)
			if err != nil {
				return nil, fmt.Errorf("QueryApplicationAll (appType=%d AND appStatus=%d) from TheDB error: %v", err, appType, appStatus)
			}
			return apps, nil
		} else {
			queryALL = fmt.Sprintf(`SELECT * FROM %s WHERE status=$2`, apg.AppTableName)
			err = apg.TheDB.Select(&apps, queryALL, appStatus)
			if err != nil {
				return nil, fmt.Errorf("QueryApplicationAll (appType=%d AND appStatus=%d) from TheDB error: %v", err, appType, appStatus)
			}
			return apps, nil
		}
	} else {
		if appStatus == application.AppStatusALL {
			queryALL = fmt.Sprintf(`SELECT * FROM %s WHERE application_type=$2`, apg.AppTableName)
			err = apg.TheDB.Select(&apps, queryALL, appType)
			if err != nil {
				return nil, fmt.Errorf("QueryApplicationAll (appType=%d AND appStatus=%d) from TheDB error: %v", err, appType, appStatus)
			}
			return apps, nil
		} else {
			queryALL = fmt.Sprintf(`SELECT * FROM %s WHERE application_type=$2 AND status=$3`, apg.AppTableName)
			err = apg.TheDB.Select(&apps, queryALL, appType, appStatus)
			if err != nil {
				return nil, fmt.Errorf("QueryApplicationAll (appType=%d AND appStatus=%d) from TheDB error: %v", err, appType, appStatus)
			}
			return apps, nil
		}
	}
}

func (apg ApplicationPg) QueryApplicationByFilter(appFilter func(application.Application) bool) (apps []application.Application, err error) {
	queryAll := fmt.Sprintf(`SELECT * FROM %s`, apg.AppTableName)
	var appAll []application.Application
	err = apg.TheDB.Select(&appAll, queryAll)
	if err != nil {
		return nil, fmt.Errorf("query all application from TheDB error: %v", err)
	}

	// apply filter
	for _, app := range appAll {
		if appFilter(app) == true {
			apps = append(apps, app)
		}
	}
	return apps, nil
}

func (apg ApplicationPg) InsertApplication(app application.Application) (appID int, err error) {
	execInsert := fmt.Sprintf(`INSERT INTO %s (project_id, application_type, status, app_user_id, app_user_cn_name, department_code, basic_content, extra_content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING application_id`, apg.AppTableName)
	err = apg.TheDB.Get(&appID, execInsert,
		app.ProjectID, app.Type, app.Status,
		app.ApplicationID, app.ApplicantUserChineseName, app.DepartmentCode,
		app.BasicContent, app.ExtraContent,
		time.Now(), time.Now())
	if err != nil {
		return -1, fmt.Errorf("TheDB.Get InsertApplication in TheDB error: %v", err)
	}
	return appID, nil
}

func (apg ApplicationPg) UpdateApplication(app application.Application) (err error) {
	execUpdate := fmt.Sprintf(`UPDATE %s SET status=:status, basic_content=:basic_content, extra_content=:extra_content, updated_at=:updated_at WHERE application_id=:application_id`, apg.AppTableName)

	_, err = apg.TheDB.NamedExec(execUpdate, app)
	if err != nil {
		return fmt.Errorf("TheDB.NamedExec UPDATE Application error: %v", err)
	}
	return nil
}

func (apg ApplicationPg) InsertAppOps(record application.AppOpsRecord) (recordID int, err error) {
	execInsert := fmt.Sprintf(`INSERT INTO %s (project_id, application_id, ops_user_id, ops_user_cn_name, action, action_str, basic_info, extra_info, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING record_id`, apg.OpsTableName)
	err = apg.TheDB.Get(&recordID, execInsert,
		record.ProjectID, record.ApplicationID, record.OpsUserID, record.OpsUserChineseName,
		record.Action, record.ActionStr, record.BasicInfo, record.ExtraInfo, time.Now())
	if err != nil {
		return -1, fmt.Errorf("TheDB.Get InsertAppOps in TheDB error: %v", err)
	}
	return recordID, nil
}

func (apg ApplicationPg) QueryAppOpsByAppId(applicationID int) (records []application.AppOpsRecord, err error) {
	queryByAppId := fmt.Sprintf(`SELECT * FROM %s WHERE application_id=$1`, apg.OpsTableName)
	err = apg.TheDB.Select(&records, queryByAppId, applicationID)
	if err != nil {
		return nil, fmt.Errorf("query application operations By application id from TheDB error: %v", err)
	}
	return records, nil
}

func (apg ApplicationPg) ArchiveToHistory(historyADI ApplicationDbInfo, projectID int) (err error) {
	//TODO
	return fmt.Errorf("the method has not been implemented")
}
