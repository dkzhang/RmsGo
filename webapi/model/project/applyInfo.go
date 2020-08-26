package project

import "time"

type ApplyInfoEx struct {
	ProjectID int `db:"project_id" json:"project_id"`
	ApplyInfo
}

type ApplyInfo struct {
	StartDate      time.Time `db:"start_date" json:"start_date"`
	TotalDaysApply int       `db:"total_days_apply" json:"total_days_apply"`
	EndReminderAt  time.Time `db:"end_reminder_at" json:"end_reminder_at"`

	CpuNodesExpected    int `db:"cpu_nodes_expected" json:"cpu_nodes_expected"`
	GpuNodesExpected    int `db:"gpu_nodes_expected" json:"gpu_nodes_expected"`
	StorageSizeExpected int `db:"storage_size_expected" json:"storage_size_expected"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
