package project

import (
	"github.com/dkzhang/RmsGo/webapi/model/metering"
	"github.com/dkzhang/RmsGo/webapi/model/resource"
	"github.com/dkzhang/RmsGo/webapi/model/resourceApplication"
	"time"
)

type ProjectInfo struct {
	ProjectID   int
	ProjectName string

	DepartmentCode   string
	Department       string
	ChiefID          int
	ChiefChineseName string
	ExtraInfo        string

	Status int

	ResourceExpected resource.Resource
	ResourceAcquired resource.Resource

	TheApplications    []resourceApplication.Application
	TheResAllocRecords []resource.ResAllocRecord
	TheMeteringList    []metering.MeteringList

	ArchivedAt time.Time
	CreatedAt  time.Time
}

type ProjectDynamicInfo struct {
	Status int

	ResourceExpected resource.Resource
	ResourceAcquired resource.Resource

	//CpuNodesExpected    int `json:"cpu_nodes_expected"`
	//GpuNodesExpected    int `json:"gpu_nodes_expected"`
	//StorageSizeExpected int `json:"storage_size_expected"`
	//CpuNodesAcquired    int `json:"cpu_nodes_acquired"`
	//GpuNodesAcquired    int `json:"gpu_nodes_acquired"`
	//StorageSizeAcquired int `json:"storage_size_acquired"`

	StartBillingAt time.Time
	DaysApply      int //Number of days to apply for resources
	EndReminderAt  time.Time
	ArchivedAt     time.Time

	AppNumSubmitted  int
	AppNumUnfinished int
	ResAllocNum      int // the number of resource allocation records
}
