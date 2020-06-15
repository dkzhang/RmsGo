package resourceApplication

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Set up a JSON-based Application struct in a unified format
// to facilitate the generation of historical record arrays
type Application struct {
	gorm.Model
	ProjectID     int
	Type          int
	Status        int
	BasicContent  string
	ExtraContent  string
	TheOpsRecords []AppOpsRecord
}

type AppOpsRecord struct {
	gorm.Model
	OpsUserID          int
	OpsUserChineseName int
	Action             int
	ActionStr          string
	BasicInfo          string
	ExtraInfo          string
}

const (
	AppStatusProjectChief  = 1
	AppStatusApprover      = 2
	AppStatusController    = 7
	AppStatusToBeProcessed = 8
	AppStatusArchived      = 9
)

const (
	AppTypeNew           = 1
	AppTypeChange        = 2
	AppTypeReturnCompute = 3
	AppTypeReturnStorage = 4
	AppTypeMetering      = 5
)
