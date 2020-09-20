package gfApplication

import (
	"encoding/json"
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/timeParse"
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

// 将时间从0点统一调整至特定时刻
const TimeAdjust = time.Hour * 10

func JsonUnmarshalAppNewProRes(jsonStr string) (appn AppNewProRes, err error) {
	err = json.Unmarshal([]byte(jsonStr), &appn)
	if err != nil {
		return AppNewProRes{}, fmt.Errorf("json.Unmarshal error: %v", err)
	}

	appn.StartDate, err = timeParse.ParseDateShangHai(appn.StartDateStr)
	if err != nil {
		return AppNewProRes{}, fmt.Errorf("timeParse StartDate error: %v", err)
	}
	appn.EndDate, err = timeParse.ParseDateShangHai(appn.EndDateStr)
	if err != nil {
		return AppNewProRes{}, fmt.Errorf("timeParse EndDate error: %v", err)
	}

	if appn.TotalDaysApply != int(appn.EndDate.Sub(appn.StartDate).Hours()/24) {
		return AppNewProRes{},
			fmt.Errorf("TotalDaysApply compution error, start=%v, end=%v, duration=%d days",
				appn.StartDate, appn.EndDate, appn.TotalDaysApply)
	}

	// 将时间从0点统一调整至特定时刻
	appn.StartDate = appn.StartDate.Add(TimeAdjust)
	appn.EndDate = appn.EndDate.Add(TimeAdjust)

	return appn, nil
}

type AppResChange struct {
	resource.Resource
	DaysExtended int       `json:"days_extended"`
	EndDateStr   string    `json:"end_date"`
	EndDate      time.Time `json:"-"`
}

func JsonUnmarshalAppResChange(jsonStr string) (appc AppResChange, err error) {
	err = json.Unmarshal([]byte(jsonStr), &appc)
	if err != nil {
		return AppResChange{}, fmt.Errorf("json.Unmarshal error: %v", err)
	}

	appc.EndDate, err = timeParse.ParseDateShangHai(appc.EndDateStr)
	if err != nil {
		return AppResChange{}, fmt.Errorf("timeParse EndDate error: %v", err)
	}

	// 将时间从0点统一调整至特定时刻
	appc.EndDate = appc.EndDate.Add(TimeAdjust)

	return appc, nil
}

type AppResComReturn struct {
	CGpuType         int     `json:"cgpu_type"` // cpu=1, gpu=2
	NodesReturnArray []int64 `json:"nodes_return_array"`
}

func JsonUnmarshalAppResComReturn(jsonStr string) (arcr AppResComReturn, err error) {
	err = json.Unmarshal([]byte(jsonStr), &arcr)
	if err != nil {
		return AppResComReturn{}, fmt.Errorf("json.Unmarshal error: %v", err)
	}
	return arcr, nil
}

//type AppResStoReturn struct {
//	// nothing
//}

type CtrlApprovalInfo struct {
	ProjectCode string `json:"project_code"`
	Comment     string `json:"comment"`
}

type ApprApprovalInfo struct {
	Comment string `json:"comment"`
}
