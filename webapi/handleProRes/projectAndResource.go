package handleProRes

import (
	"github.com/gin-gonic/gin"
	"time"
)

type NewProResWA struct {
	ProjectID     int
	ApplicationID int
	ProjectName   string
	Resource      ResourceApplication
	Action        int
	ExtraInfo     string
}

func ApplyNewProRes(c *gin.Context) {

}

// Set up a JSON-based Application struct in a unified format
// to facilitate the generation of historical record arrays
type Application struct {
	ProjectID     int
	ApplicationID int
	Type          int
	Status        int
	BasicContent  string
	ExtraContent  string
	TheOpsRecords []AppOpsRecord
}

type AppOpsRecord struct {
	OpsUserID          int
	OpsUserChineseName int
	Action             int
	OpsDatetime        time.Time
	BasicInfo          string
	ExtraInfo          string
}

const (
	AppStatusProjectChief = 1
	AppStatusApprover     = 2
	AppStatusController   = 7
	AppStatusArchived     = 9
)

const (
	AppTypeNew           = 1
	AppTypeExpansion     = 2
	AppTypeReturnCompute = 3
	AppTypeReturnStorage = 4
	AppTypeMetering      = 5
)

type ProjectInfo struct {
	ID   int
	Name string

	DepartmentCode   string
	Department       string
	ChiefID          int
	ChiefChineseName string

	CreateDateTime  time.Time
	ArchiveDateTime time.Time
}

type ResourceApplication struct {
	CpuNodes    int
	GpuNodes    int
	StorageSize int
	StartDate   time.Time
	EndDate     time.Time
}

type ResourceExpansionApplication struct {
	CpuNodes    int
	GpuNodes    int
	StorageSize int
	EndDate     time.Time
	Type        int
}

const (
	ResExTypeCPU = 1 << iota
	ResExTypeGPU
	ResExTypeStorage
	ResExTypeDate
)
