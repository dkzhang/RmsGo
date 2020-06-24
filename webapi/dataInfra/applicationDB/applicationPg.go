package applicationDB

import (
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/model/application"
	"github.com/jmoiron/sqlx"
	"time"
)

type ApplicationPg struct {
	db        *sqlx.DB
	tableName string
}

func NewApplicationPg(sqlxdb *sqlx.DB, tname string) ApplicationPg {
	return ApplicationPg{
		db:        sqlxdb,
		tableName: tname}
}

func (apg ApplicationPg) QueryApplicationByID(applicationID int) (app application.Application, err error) {
	queryByID := fmt.Sprintf(`SELECT * FROM %s WHERE application_id=$1`, apg.tableName)
	err = apg.db.Get(&app, queryByID, applicationID)
	if err != nil {
		return application.Application{},
			fmt.Errorf("QueryApplicationByID in db error: %v", err)
	}
	return app, nil
}

func (apg ApplicationPg) QueryApplicationByOwner(userID int) (apps []application.Application, err error) {
	queryByOwner := fmt.Sprintf(`SELECT * FROM %s WHERE app_user_id=$1`, apg.tableName)
	err = apg.db.Select(&apps, queryByOwner, userID)
	if err != nil {
		return nil, fmt.Errorf("QueryGeneralFormDraftByOwner from db error: %v", err)
	}
	return apps, nil
}

func (apg ApplicationPg) QueryApplicationByDepartmentCode(dc string) (apps []application.Application, err error) {
	queryByOwner := fmt.Sprintf(`SELECT * FROM %s WHERE department_code=$1 AND status=$2 `, apg.tableName)
	err = apg.db.Select(&apps, queryByOwner, dc)
	if err != nil {
		return nil, fmt.Errorf("QueryApplicationByDepartmentCode from db error: %v", err)
	}
	return apps, nil
}

func (apg ApplicationPg) QueryApplicationByFilter(appFilter func(application.Application) bool) (apps []application.Application, err error) {
	queryAll := fmt.Sprintf(`SELECT * FROM %s`, apg.tableName)
	var appAll []application.Application
	err = apg.db.Select(&appAll, queryAll)
	if err != nil {
		return nil, fmt.Errorf("query all application from db error: %v", err)
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
	execInsert := fmt.Sprintf(`INSERT INTO %s (project_id, application_type, status, app_user_id, app_user_cn_name, department_code, basic_content, extra_content, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING application_id`, apg.tableName)
	err = apg.db.Get(&appID, execInsert,
		app.ProjectID, app.Type, app.Status,
		app.ApplicationID, app.ApplicantUserChineseName, app.DepartmentCode,
		app.BasicContent, app.ExtraContent, time.Now())
	if err != nil {
		return -1, fmt.Errorf("db.Get InsertApplication in db error: %v", err)
	}
	return appID, nil
}

func (apg ApplicationPg) UpdateApplication(app application.Application) (err error) {
	execUpdate := fmt.Sprintf(`UPDATE %s SET status=:status, basic_content=:basic_content, extra_content=:extra_content WHERE application_id=:application_id`, apg.tableName)

	_, err = apg.db.NamedExec(execUpdate, app)
	if err != nil {
		return fmt.Errorf("db.NamedExec UPDATE Application error: %v", err)
	}
	return nil
}

func (apg ApplicationPg) InsertAppOps(record application.AppOpsRecord) (recordID int, err error) {
	execInsert := fmt.Sprintf(`INSERT INTO %s (project_id, application_id, ops_user_id, ops_user_cn_name, action, action_str, basic_info, extra_info, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING application_id`, apg.tableName)
	err = apg.db.Get(&recordID, execInsert,
		record.ProjectID, record.ApplicationID, record.OpsUserID, record.OpsUserChineseName,
		record.Action, record.ActionStr, record.BasicInfo, record.ExtraInfo, time.Now())
	if err != nil {
		return -1, fmt.Errorf("db.Get InsertAppOps in db error: %v", err)
	}
	return recordID, nil
}

func (apg ApplicationPg) QueryAppOpsByAppId(applicationID int) (records []application.AppOpsRecord, err error) {
	queryByAppId := fmt.Sprintf(`SELECT * FROM %s WHERE application_id=$1`, apg.tableName)
	err = apg.db.Select(&records, queryByAppId, applicationID)
	if err != nil {
		return nil, fmt.Errorf("query application operations By application id from db error: %v", err)
	}
	return records, nil
}
