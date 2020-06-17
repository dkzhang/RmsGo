package resourceApplication

import (
	"github.com/dkzhang/RmsGo/webapi/model/resource"
	"time"
)

type ApplicationWA struct {
	ProjectID     int    `json:"project_id"`
	ApplicationID int    `json:"application_id"`
	Action        int    `json:"action"`
	BasicContent  string `json:"basic_content"`
	ExtraContent  string `json:"extra_content"`
}

/////////////////////////////////////////////////////////////////

type AppNewProRes struct {
	ProjectName string `json:"project_name"`
	resource.Resource
	StartDate time.Time `json:"start_date"`
	DaysOfUse int       `json:"days_of_use"`
	EndDate   time.Time `json:"end_date"`
}

type AppResChange struct {
	resource.Resource
	DaysExtended int       `json:"days_extended"`
	EndDate      time.Time `json:"end_date"`
}

type AppResComReturn struct {
	CpuNodesNum int
	CpuNodesMap string
	GpuNodesNum int
	GpuNodesMap string
}

//type AppResStoReturn struct {
//	// nothing
//}
