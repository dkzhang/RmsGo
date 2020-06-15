package handleProRes

import (
	"github.com/gin-gonic/gin"
	"time"
)

type AppNewProResWA struct {
	ProjectID     int
	ApplicationID int
	ProjectName   string
	Resource      ResourceApplication
	Action        int
	ExtraInfo     string
}

func ApplyNewProRes(c *gin.Context) {

}

type AppExResWA struct {
}

type AppReComWA struct {
}

type AppReStoWA struct {
}

type AppMeterWA struct {
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
	ActionStr          string
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
	AppTypeChange        = 2
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

	Status int

	ResourceExpected Resource
	ResourceAcquired Resource

	CreateDateTime  time.Time
	ArchiveDateTime time.Time
}

type Resource struct {
	CpuNodes    int
	GpuNodes    int
	StorageSize int
}

type ResourceDetails struct {
	CpuNodes      int
	CpuNodeBefore string
	CpuNodeAfter  string
	GpuNodes      int
	GpuNodeBefore string
	GpuNodeAfter  string
	StorageSize   int
	StorageBefore int
	StorageAfter  int
}

type ResourceAllocationRecord struct {
	ProjectID       int
	AllocationID    int
	AllocResDetails ResourceDetails
	UserID          int
	UserChineseName string
	Datetime        time.Time
}

type ResourceApplication struct {
	CpuNodes    int
	GpuNodes    int
	StorageSize int
	StartDate   time.Time
	EndDate     time.Time
}
