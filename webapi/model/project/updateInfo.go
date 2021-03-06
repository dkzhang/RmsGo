package project

import "time"

type BasicInfo struct {
	ProjectID   int       `db:"project_id" json:"project_id"`
	ProjectName string    `db:"project_name" json:"project_name"`
	ExtraInfo   string    `db:"extra_info" json:"extra_info"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type CodeInfo struct {
	ProjectID   int       `db:"project_id" json:"project_id"`
	ProjectCode string    `db:"project_code" json:"project_code"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type StatusInfo struct {
	ProjectID   int       `db:"project_id" json:"project_id"`
	BasicStatus int       `db:"basic_status" json:"basic_status"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type ApplyInfo struct {
	ProjectID           int       `db:"project_id" json:"project_id"`
	StartDate           time.Time `db:"start_date" json:"start_date"`
	TotalDaysApply      int       `db:"total_days_apply" json:"total_days_apply"`
	EndReminderAt       time.Time `db:"end_reminder_at" json:"end_reminder_at"`
	CpuNodesExpected    int       `db:"cpu_nodes_expected" json:"cpu_nodes_expected"`
	GpuNodesExpected    int       `db:"gpu_nodes_expected" json:"gpu_nodes_expected"`
	StorageSizeExpected int       `db:"storage_size_expected" json:"storage_size_expected"`
	UpdatedAt           time.Time `db:"updated_at" json:"updated_at"`
}

type AllocNum struct {
	ProjectID           int       `db:"project_id" json:"project_id"`
	CpuNodesAcquired    int       `db:"cpu_nodes_acquired" json:"cpu_nodes_acquired"`
	GpuNodesAcquired    int       `db:"gpu_nodes_acquired" json:"gpu_nodes_acquired"`
	StorageSizeAcquired int       `db:"storage_size_acquired" json:"storage_size_acquired"`
	UpdatedAt           time.Time `db:"updated_at" json:"updated_at"`
}

type AllocInfo struct {
	ProjectID        int       `db:"project_id" json:"project_id"`
	AccountAllocInfo string    `db:"account_alloc_info" json:"account_alloc_info"`
	StorageAllocInfo string    `db:"storage_alloc_info" json:"storage_alloc_info"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
}
