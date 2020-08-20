package applicationDB

import (
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"github.com/jmoiron/sqlx"
	"time"
)

type ApplicationPg struct {
	DBInfo
}

func NewApplicationPg(sqlxdb *sqlx.DB, appname, opsname string) ApplicationPg {
	return ApplicationPg{
		DBInfo{
			TheDB:        sqlxdb,
			AppTableName: appname,
			OpsTableName: opsname,
		},
	}
}

func (apg ApplicationPg) Close() {
	apg.TheDB.Close()
}

func (apg ApplicationPg) QueryByID(applicationID int) (app application.Application, err error) {
	queryByID := fmt.Sprintf(`SELECT * FROM %s WHERE application_id=$1`, apg.AppTableName)
	err = apg.TheDB.Get(&app, queryByID, applicationID)
	if err != nil {
		return application.Application{},
			fmt.Errorf("QueryByID in TheDB error: %v", err)
	}
	return app, nil
}

func (apg ApplicationPg) QueryByOwner(userID int, appType int, appStatus int) (apps []application.Application, err error) {
	queryByOwner := fmt.Sprintf(`SELECT * FROM %s WHERE app_user_id=$1 AND application_type & $2 != 0 AND status & $3 != 0`, apg.AppTableName)
	err = apg.TheDB.Select(&apps, queryByOwner, userID, appType, appStatus)
	if err != nil {
		return nil, fmt.Errorf("QueryByDepartmentCode (userID=%d AND appType=%d AND appStatus=%d) from TheDB error: %v", userID, appType, appStatus, err)
	}
	return apps, nil
}

func (apg ApplicationPg) QueryByDepartmentCode(dc string, appType int, appStatus int) (apps []application.Application, err error) {
	queryByDC := fmt.Sprintf(`SELECT * FROM %s WHERE department_code=$1 AND application_type & $2 != 0 AND status & $3 != 0`, apg.AppTableName)
	err = apg.TheDB.Select(&apps, queryByDC, dc, appType, appStatus)
	if err != nil {
		return nil, fmt.Errorf("QueryByDepartmentCode (dc=%s AND appType=%d AND appStatus=%d) from TheDB error: %v", dc, appType, appStatus, err)
	}
	return apps, nil
}

func (apg ApplicationPg) QueryAll(appType int, appStatus int) (apps []application.Application, err error) {
	queryALL := fmt.Sprintf(`SELECT * FROM %s WHERE application_type & $1 != 0 AND status & $2 != 0`, apg.AppTableName)
	err = apg.TheDB.Select(&apps, queryALL, appType, appStatus)
	if err != nil {
		return nil, fmt.Errorf("QueryAll (appType=%d AND appStatus=%d) from TheDB error: %v", err, appType, appStatus)
	}
	return apps, nil
}

func (apg ApplicationPg) Insert(app application.Application) (appID int, err error) {
	execInsert := fmt.Sprintf(`INSERT INTO %s (project_id, application_type, status, app_user_id, app_user_cn_name, department_code, basic_content, extra_content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING application_id`, apg.AppTableName)
	err = apg.TheDB.Get(&appID, execInsert,
		app.ProjectID, app.Type, app.Status,
		app.ApplicationID, app.ApplicantUserChineseName, app.DepartmentCode,
		app.BasicContent, app.ExtraContent,
		time.Now(), time.Now())
	if err != nil {
		return -1, fmt.Errorf("TheDB.Get Insert in TheDB error: %v", err)
	}
	return appID, nil
}

func (apg ApplicationPg) Update(app application.Application) (err error) {
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

func (apg ApplicationPg) ArchiveToHistory(historyADI DBInfo, projectID int) (err error) {
	//TODO
	return fmt.Errorf("the method has not been implemented")
}
