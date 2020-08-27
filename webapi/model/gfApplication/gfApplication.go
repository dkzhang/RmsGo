package gfApplication

import (
	"encoding/json"
	"fmt"
	"github.com/dkzhang/RmsGo/webapi/model/resource"
	"time"
)

type AppNewProRes struct {
	ProjectName string `json:"project_name"`
	resource.Resource
	StartDateStr   string    `json:"start_date"`
	StartDate      time.Time `json:"-"`
	TotalDaysApply int       `json:"total_days_apply"`
	EndDateStr     string    `json:"end_date"`
	EndDate        time.Time `json:"-"`
}

func JsonUnmarshalAppNewProRes(jsonStr string) (appNewProRes AppNewProRes, err error) {
	err = json.Unmarshal([]byte(jsonStr), &appNewProRes)
	if err != nil {
		return AppNewProRes{}, fmt.Errorf("json.Unmarshal error: %v", err)
	}

	return appNewProRes, nil
}

type AppResChange struct {
	resource.Resource
	DaysExtended int       `json:"days_extended"`
	EndDateStr   string    `json:"end_date"`
	EndDate      time.Time `json:"-"`
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
