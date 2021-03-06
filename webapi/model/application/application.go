package application

import (
	"fmt"
	"time"
)

// Set up a JSON-based Application struct in a unified format
// to facilitate the generation of historical record arrays
type Application struct {
	ApplicationID            int       `db:"application_id" json:"application_id"`
	ProjectID                int       `db:"project_id" json:"project_id"`
	Type                     int       `db:"application_type" json:"application_type"`
	Status                   int       `db:"status" json:"status"`
	ApplicantUserID          int       `db:"app_user_id" json:"app_user_id"`
	ApplicantUserChineseName string    `db:"app_user_cn_name" json:"app_user_cn_name"`
	DepartmentCode           string    `db:"department_code" json:"department_code"`
	BasicContent             string    `db:"basic_content" json:"basic_content"`
	ExtraContent             string    `db:"extra_content" json:"extra_content"`
	CreatedAt                time.Time `db:"created_at" json:"created_at"`
	UpdatedAt                time.Time `db:"updated_at" json:"updated_at"`
}

type AppOpsRecord struct {
	RecordID           int       `db:"record_id" json:"record_id"`
	ProjectID          int       `db:"project_id" json:"project_id"`
	ApplicationID      int       `db:"application_id" json:"application_id"`
	OpsUserID          int       `db:"ops_user_id" json:"ops_user_id"`
	OpsUserChineseName string    `db:"ops_user_cn_name" json:"ops_user_cn_name"`
	Action             int       `db:"action" json:"action"`
	ActionStr          string    `db:"action_str" json:"action_str"`
	BasicInfo          string    `db:"basic_info" json:"basic_info"`
	ExtraInfo          string    `db:"extra_info" json:"extra_info"`
	CreatedAt          time.Time `db:"created_at" json:"-"`
	CreatedAtStr       string    `db:"-" json:"created_at"`
}

var schemaApp = `
		CREATE TABLE %s (
    		application_id SERIAL PRIMARY KEY,
			project_id int,
			application_type int,
			status int,	
			app_user_id int,
			app_user_cn_name varchar(32),
			department_code varchar(32),
			basic_content varchar(16384),
			extra_content varchar(16384),			
			created_at TIMESTAMP WITH TIME ZONE,
			updated_at TIMESTAMP WITH TIME ZONE
		);`

var schemaAppOps = `
		CREATE TABLE %s (
			record_id SERIAL PRIMARY KEY,
			project_id int,
    		application_id int,
			ops_user_id int,
			ops_user_cn_name varchar(32),			
			action int,
			action_str varchar(32),				
			basic_info varchar(16384),
			extra_info varchar(16384),			
			created_at TIMESTAMP WITH TIME ZONE
		);`

// 16K = 1024 * 16 = 16384

var TableApp = "application"
var TableHistoryApp = "history_application"

var TableAppOps = "app_ops_record"
var TableHistoryAppOps = "history_app_ops_record"

func GetSchemaApp() string {
	return fmt.Sprintf(schemaApp, TableApp)
}
func GetSchemaHistoryApp() string {
	return fmt.Sprintf(schemaApp, TableHistoryApp)
}

func GetSchemaAppOps() string {
	return fmt.Sprintf(schemaAppOps, TableAppOps)
}
func GetSchemaHistoryAppOps() string {
	return fmt.Sprintf(schemaAppOps, TableHistoryAppOps)
}
