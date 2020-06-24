package project

import (
	"time"
)

type DynamicInfo struct {
	ProjectID            int
	BasicStatus          int
	ComputingAllocStatus int
	StorageAllocStatus   int

	AppInProgressNum        int
	AppAccomplishedNum      int
	MeteringInProgressNum   int
	MeteringAccomplishedNum int
	ResAllocNum             int // the number of resource allocation records

	CpuNodesExpected    int `json:"cpu_nodes_expected"`
	GpuNodesExpected    int `json:"gpu_nodes_expected"`
	StorageSizeExpected int `json:"storage_size_expected"`
	CpuNodesAcquired    int `json:"cpu_nodes_acquired"`
	GpuNodesAcquired    int `json:"gpu_nodes_acquired"`
	StorageSizeAcquired int `json:"storage_size_acquired"`

	StartBillingAt time.Time
	DaysApply      int //Number of days to apply for resources
	EndReminderAt  time.Time
	UpdatedAt      time.Time
}

const (
	BasicStatusApplying    = 1
	BasicStatusEstablished = 2
	BasicStatusArchived    = -9
)

const (
	ResNotYetAssigned  = 1
	ResFullAllocation  = 2
	ResUnderAllocation = 3
	ResOverAllocation  = 4
	ResAllReturned     = 5
)
