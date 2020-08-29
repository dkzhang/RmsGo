package gfApplication

import (
	"github.com/dkzhang/RmsGo/webapi/model/resource"
	"time"
)

type AppNewProRes struct {
	ProjectName string `json:"project_name"`
	resource.Resource
	StartDate      time.Time `json:"start_date"`
	TotalDaysApply int       `json:"total_days_apply"`
	EndDate        time.Time `json:"end_date"`
}

type AppResChange struct {
	resource.Resource
	EndDate time.Time `json:"end_date"`
}

type AppResComReturn struct {
	CpuNodesReturnNum int    `json:"cpu_nodes_return_num"`
	CpuNodesReturnMap string `json:"cpu_nodes_return_map"`
	GpuNodesReturnNum int    `json:"gpu_nodes_return_num"`
	GpuNodesReturnMap string `json:"cpu_nodes_return_map"`
}

//type AppResStoReturn struct {
//	// nothing
//}

type AppCtrlProjectInfo struct {
	ProjectCode string `json:"project_code"`
}

//type AppCtrlAccountInfo struct {
//	AccountInfo string    `json:"account_info"`
//}
