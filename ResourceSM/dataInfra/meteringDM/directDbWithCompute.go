package meteringDM

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/compute"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/meteringDB"
	"github.com/dkzhang/RmsGo/ResourceSM/dataInfra/resAllocDM"
	"github.com/dkzhang/RmsGo/ResourceSM/model/metering"
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

func (dm DirectDbWithCompute) ComputeSettlement(projectID int) (ms metering.Statement, err error) {
	ms = metering.Statement{
		//MeteringID:           0,
		MeteringType:         metering.TypeSettlement,
		MeteringTypeInfo:     "",
		CpuAmountInDays:      0,
		GpuAmountInDays:      0,
		StorageAmountInDays:  0,
		CpuAmountInHours:     0,
		GpuAmountInHours:     0,
		StorageAmountInHours: 0,
		CpuNodeMetering:      make([]metering.MeteringItem, 0),
		GpuNodeMetering:      make([]metering.MeteringItem, 0),
		StorageMetering:      make([]metering.MeteringItem, 0),
		CpuNodeMeteringStr:   "",
		GpuNodeMeteringStr:   "",
		StorageMeteringStr:   "",
		CreatedAt:            time.Now(),
	}

	// TODO
	return ms, fmt.Errorf("not accomplishment")
}

func queryAndCompute(radm resAllocDM.ResAllocReadOnlyDM, projectID int) (amountInHours int, mis []metering.MeteringItem, err error) {
	records, err := radm.QueryByProjectID(projectID)
	if err != nil {
		return -1, nil, fmt.Errorf("resAllocDM.ResAllocReadOnlyDM QueryByProjectID (projectID=%d) error: %v", projectID, err)
	}

	if records == nil {

	}

	if len(records) == 1 {
		//error
	}

	if records[len(records)-1].NumAfter != 0 {

	}

	amountInHours, mis = compute.MeteringCompute(records, records[0].CreatedAt, records[len(records)-1].CreatedAt)

	return amountInHours, mis, nil
}
