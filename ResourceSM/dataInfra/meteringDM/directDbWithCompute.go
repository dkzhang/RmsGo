package meteringDM

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/meteringDB"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resAllocDM"
	"github.com/dkzhang/RmsGo/ResourceSM/model/metering"
	"github.com/dkzhang/RmsGo/ResourceSM/resMetering/meteringComputation"
	"github.com/dkzhang/RmsGo/myUtils/timeParse"
	"time"
)

type DirectDbWithCompute struct {
	theDB meteringDB.MeteringDB

	cpuAllocDM     resAllocDM.ResAllocReadOnlyDM
	gpuAllocDM     resAllocDM.ResAllocReadOnlyDM
	storageAllocDM resAllocDM.ResAllocReadOnlyDM
}

func NewResAllocDirectDB(mdb meteringDB.MeteringDB,
	cdm resAllocDM.ResAllocReadOnlyDM, gdm resAllocDM.ResAllocReadOnlyDM, sdm resAllocDM.ResAllocReadOnlyDM) DirectDbWithCompute {
	return DirectDbWithCompute{
		theDB:          mdb,
		cpuAllocDM:     cdm,
		gpuAllocDM:     gdm,
		storageAllocDM: sdm,
	}
}

func (dm DirectDbWithCompute) Query(projectID int, mType int, typeInfo string) (ms metering.Statement, err error) {
	return dm.theDB.Query(projectID, mType, typeInfo)
}

func (dm DirectDbWithCompute) QueryAll(projectID int, mType int) (mss []metering.Statement, err error) {
	return dm.theDB.QueryAll(projectID, mType)
}

///////////////////////////////////////////////////////////////////////////////////////////////////

func (dm DirectDbWithCompute) QueryWithCreate(projectID int, mType int, typeInfo string) (ms metering.Statement, err error) {
	ms, err = dm.theDB.Query(projectID, mType, typeInfo)

	if err != nil {
		if err != sql.ErrNoRows {
			// db error
			return metering.Statement{}, fmt.Errorf("query IsExist in MeteringDB error")
		}
	} else {
		//metering.Statement exist
		return ms, nil
	}

	//metering.Statement dose not exist
	switch mType {
	case metering.TypeMonthly:
		return metering.Statement{}, fmt.Errorf("metering Monthly feature is not implemented yet")
	case metering.TypeQuarterly:
		return metering.Statement{}, fmt.Errorf("metering Quarterly feature is not implemented yet")
	case metering.TypeAnnual:
		return metering.Statement{}, fmt.Errorf("metering Annual feature is not implemented yet")
	case metering.TypeAnyPeriod:
		return metering.Statement{}, fmt.Errorf("metering AnyPeriod feature is not implemented yet")
	case metering.TypeSettlement:
		ms, err = dm.ComputeSettlement(projectID)
		if err != nil {
			return metering.Statement{}, fmt.Errorf("ComputeSettlement error: %v", err)
		}

		_, _ = dm.theDB.Insert(ms)
		//ignore db insert metering error

		return ms, nil
	default:
		return metering.Statement{}, fmt.Errorf("unsupported metering type: %d", mType)
	}
}

func (dm DirectDbWithCompute) ComputeSettlement(projectID int) (ms metering.Statement, err error) {
	ms = metering.Statement{
		//MeteringID:           0,
		ProjectID:            projectID,
		MeteringType:         metering.TypeSettlement,
		MeteringTypeInfo:     "",
		CpuAmountInDays:      0,
		GpuAmountInDays:      0,
		StorageAmountInDays:  0,
		CpuAmountInHours:     0,
		GpuAmountInHours:     0,
		StorageAmountInHours: 0,
		CpuNodeMeteringJson:  "",
		GpuNodeMeteringJson:  "",
		StorageMeteringJson:  "",
		CreatedAt:            time.Now(),
	}

	cpuNodeMetering := make([]metering.MeteringItem, 0)
	gpuNodeMetering := make([]metering.MeteringItem, 0)
	storageMetering := make([]metering.MeteringItem, 0)

	///////////////////////////////////////////////////////////////////////////////////////////////
	//Compute settlement from ResAllocDM
	ms.CpuAmountInHours, cpuNodeMetering, err = computeSettlementFromResAllocDM(dm.cpuAllocDM, projectID)
	if err != nil {
		return metering.Statement{}, fmt.Errorf("computeSettlement From CPU ResAllocDM error: %v", err)
	}

	ms.GpuAmountInHours, gpuNodeMetering, err = computeSettlementFromResAllocDM(dm.gpuAllocDM, projectID)
	if err != nil {
		return metering.Statement{}, fmt.Errorf("computeSettlement From GPU ResAllocDM error: %v", err)
	}

	ms.StorageAmountInHours, storageMetering, err = computeSettlementFromResAllocDM(dm.storageAllocDM, projectID)
	if err != nil {
		return metering.Statement{}, fmt.Errorf("computeSettlement From Storage ResAllocDM error: %v", err)
	}

	///////////////////////////////////////////////////////////////////////////////////////////////
	// json Metering details
	var jsonB []byte
	jsonB, err = json.Marshal(cpuNodeMetering)
	if err != nil {
		return metering.Statement{}, fmt.Errorf("json.Marshal ms.CpuNodeMetering error: %v", err)
	}
	ms.CpuNodeMeteringJson = string(jsonB)

	jsonB, err = json.Marshal(gpuNodeMetering)
	if err != nil {
		return metering.Statement{}, fmt.Errorf("json.Marshal ms.CpuNodeMetering error: %v", err)
	}
	ms.GpuNodeMeteringJson = string(jsonB)

	jsonB, err = json.Marshal(storageMetering)
	if err != nil {
		return metering.Statement{}, fmt.Errorf("json.Marshal ms.CpuNodeMetering error: %v", err)
	}
	ms.StorageMeteringJson = string(jsonB)

	///////////////////////////////////////////////////////////////////////////////////////////////
	// meteringComputation amountInDays
	ms.CpuAmountInDays = timeParse.HoursToDays(ms.CpuAmountInHours)
	ms.GpuAmountInDays = timeParse.HoursToDays(ms.GpuAmountInHours)
	ms.StorageAmountInDays = timeParse.HoursToDays(ms.StorageAmountInHours)

	return ms, nil
}

func computeSettlementFromResAllocDM(radm resAllocDM.ResAllocReadOnlyDM, projectID int) (amountInHours int, mis []metering.MeteringItem, err error) {
	records, err := radm.QueryByProjectID(projectID)
	if err != nil {
		return -1, nil, fmt.Errorf("resAllocDM.ResAllocReadOnlyDM QueryByProjectID (projectID=%d) error: %v", projectID, err)
	}

	if records == nil || len(records) == 0 {
		// empty records is normal
		return 0, []metering.MeteringItem{}, nil
	}

	if records[len(records)-1].NumAfter != 0 {
		// the last record in settlement meteringComputation must have NumAfter==0
	}

	amountInHours, mis = meteringComputation.MeteringCompute(records, records[0].CreatedAt, records[len(records)-1].CreatedAt)

	return amountInHours, mis, nil
}
