package gfApplication

import (
	"github.com/dkzhang/RmsGo/webapi/model/resource"
	"time"
)

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
	CpuNodesReturnNum int
	CpuNodesReturnMap string
	GpuNodesReturnNum int
	GpuNodesReturnMap string
}

//type AppResStoReturn struct {
//	// nothing
//}
