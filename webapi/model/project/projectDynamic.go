package project

import (
	"time"
)

type DynamicInfo struct {
	ProjectID            int `db:"project_id" json:"project_id"`
	BasicStatus          int `db:"basic_status" json:"basic_status"`
	ComputingAllocStatus int `db:"computing_alloc_status" json:"computing_alloc_status"`
	StorageAllocStatus   int `db:"storage_alloc_status" json:"storage_alloc_status"`

	StartDate time.Time `db:"start_date" json:"start_date"`
	EndDate   time.Time `db:"end_date" json:"end_date"`
	//StartBillingAt time.Time `db:"start_billing_at" json:"start_billing_at"`
	TotalDaysApply int       `db:"total_days_apply" json:"total_days_apply"`
	EndReminderAt  time.Time `db:"end_reminder_at" json:"end_reminder_at"`

	AppInProgressNum        int `db:"app_in_progress_num" json:"app_in_progress_num"`
	AppAccomplishedNum      int `db:"app_accomplished_num" json:"app_accomplished_num"`
	MeteringInProgressNum   int `db:"metering_in_progress_num" json:"metering_in_progress_num"`
	MeteringAccomplishedNum int `db:"metering_accomplished_num" json:"metering_accomplished_num"`
	ResAllocNum             int `db:"res_alloc_num" json:"res_alloc_num"`

	CpuNodesExpected    int `db:"cpu_nodes_expected" json:"cpu_nodes_expected"`
	GpuNodesExpected    int `db:"gpu_nodes_expected" json:"gpu_nodes_expected"`
	StorageSizeExpected int `db:"storage_size_expected" json:"storage_size_expected"`
	CpuNodesAcquired    int `db:"cpu_nodes_acquired" json:"cpu_nodes_acquired"`
	GpuNodesAcquired    int `db:"gpu_nodes_acquired" json:"gpu_nodes_acquired"`
	StorageSizeAcquired int `db:"storage_size_acquired" json:"storage_size_acquired"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

var SchemaDynamicInfo = `
		CREATE TABLE user_info (
    		project_id int UNIQUE,
			basic_status int,
			computing_alloc_status int,
			storage_alloc_status int,
			start_date TIMESTAMP WITH TIME ZONE,
			days_of_use int,
			end_date TIMESTAMP WITH TIME ZONE,
			app_in_progress_num int,
			app_accomplished_num int,
			metering_in_progress_num int,
			metering_accomplished_num int,
			res_alloc_num int,
			cpu_nodes_expected int,
			gpu_nodes_expected int,
			storage_size_expected int,
			cpu_nodes_acquired int,
			gpu_nodes_acquired int,
			storage_size_acquired int,			
			total_days_apply int,
			end_reminder_at TIMESTAMP WITH TIME ZONE,
			created_at TIMESTAMP WITH TIME ZONE,
			updated_at TIMESTAMP WITH TIME ZONE
		);
		`

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
