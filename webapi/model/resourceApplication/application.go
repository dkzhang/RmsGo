package resourceApplication

import (
	"time"
)

type ApplicationWA struct {
	ProjectID     int    `json:"project_id"`
	ApplicationID int    `json:"application_id"`
	Action        int    `json:"action"`
	BasicContent  string `json:"basic_content"`
	ExtraContent  string `json:"extra_content"`
}

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
	TheOpsRecords []AppOpsRecord
}

type AppOpsRecord struct {
	RecordID           int
	OpsUserID          int
	OpsUserChineseName int
	Action             int
	ActionStr          string
	BasicInfo          string
	ExtraInfo          string
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
