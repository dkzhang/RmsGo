package application

import (
	"time"
)

// Set up a JSON-based Application struct in a unified format
// to facilitate the generation of historical record arrays
type Application struct {
	ApplicationID int
	ProjectID     int
	Type          int
	Status        int
	BasicContent  string
	ExtraContent  string
	CreatedAt     time.Time
}

type AppOpsRecord struct {
	RecordID           int       `db:"record_id" json:"record_id"`
	ProjectID          int       `db:"project_id" json:"project_id"`
	ApplicationID      int       `db:"application_id" json:"application_id"`
	OpsUserID          int       `db:"ops_user_id" json:"ops_user_id"`
	OpsUserChineseName int       `db:"ops_user_cn_name" json:"ops_user_cn_name"`
	Action             int       `db:"action" json:"action"`
	ActionStr          string    `db:"action_str" json:"action_str"`
	BasicInfo          string    `db:"project_id" json:"project_id"` //TODO
	ExtraInfo          string    `db:"project_id" json:"project_id"`
	CreatedAt          time.Time `db:"project_id" json:"project_id"`
}

const (
	AppStatusProjectChief = 1
	AppStatusApprover     = 2
	AppStatusController   = 7
	AppStatusArchived     = 8
)

const (
	AppTypeNew           = 1
	AppTypeChange        = 2
	AppTypeReturnCompute = 3
	AppTypeReturnStorage = 4
	AppTypeMetering      = 5
)
